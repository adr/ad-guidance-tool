package decision

import (
	"errors"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestComment_StoresActualText is the architectural anchor of PR 1b: the §A.1
// data-loss bug, where the legacy implementation stored a placeholder count in
// Comment.Comment and lost the real text, must not regress. The Decision passed
// to Save must carry the author + the actual comment text (and a non-numeric
// Text field).
func TestComment_StoresActualText(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	d := &Decision{ID: "0001"}
	commentText := "This is a real comment that must be preserved"

	mockRepo.On("LoadBody", "model", "0001").Return("body", nil)
	mockRepo.On("Save", "model", mock.MatchedBy(func(saved *Decision) bool {
		if len(saved.Comments) != 1 {
			return false
		}
		c := saved.Comments[0]
		if c.Author != "Jane" || c.Text != commentText {
			return false
		}
		// Defend against the §A.1 regression: Text must not be a numeric placeholder.
		if _, err := strconv.Atoi(strings.TrimSpace(c.Text)); err == nil {
			return false
		}
		return true
	}), "body").Return(nil)

	err := service.Comment("model", d, "Jane", commentText)

	assert.NoError(t, err)
	assert.Len(t, d.Comments, 1)
	assert.Equal(t, commentText, d.Comments[0].Text)
	mockRepo.AssertExpectations(t)
}

func TestAddNew_InvalidTitle(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	result, err := service.AddNew("model", "12345 !!!")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "title must contain at least one letter")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestAddNew_DelegatesToRepo(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	created := &Decision{ID: "0001", Title: "Create something", Status: "proposed"}
	mockRepo.On("Create", "model", "", mock.MatchedBy(func(d *Decision) bool {
		return d.Title == "Create something" && d.Status == "proposed" && d.Date != ""
	})).Return(created, nil)

	result, err := service.AddNew("model", "Create something")

	assert.NoError(t, err)
	assert.Equal(t, created, result)
	mockRepo.AssertExpectations(t)
}

func TestLink_RefusesSupersedesTag(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	source := &Decision{ID: "0001"}
	target := &Decision{ID: "0002"}

	err := service.Link("model", source, target, "supersedes", "superseded-by")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "use 'adg supersede'")
	mockRepo.AssertNotCalled(t, "Save")
}

func TestLink_CustomTagsWithReverse(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	source := &Decision{ID: "0001"}
	target := &Decision{ID: "0002"}

	mockRepo.On("LoadBody", "model", "0001").Return("source body", nil)
	mockRepo.On("Save", "model", source, "source body").Return(nil)
	mockRepo.On("LoadBody", "model", "0002").Return("target body", nil)
	mockRepo.On("Save", "model", target, "target body").Return(nil)

	err := service.Link("model", source, target, "relates", "linked-back")

	assert.NoError(t, err)
	assert.Equal(t, []string{"0002"}, source.Links["relates"])
	assert.Equal(t, []string{"0001"}, target.Links["linked-back"])
	mockRepo.AssertExpectations(t)
}

func TestTag_AddsNewTag(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	d := &Decision{ID: "0001", Tags: []string{"existing"}}

	mockRepo.On("LoadBody", "model", "0001").Return("body", nil)
	mockRepo.On("Save", "model", d, "body").Return(nil)

	err := service.Tag("model", d, "new-tag")

	assert.NoError(t, err)
	assert.Contains(t, d.Tags, "new-tag")
	mockRepo.AssertExpectations(t)
}

func TestTag_DuplicateTagFails(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	d := &Decision{ID: "0001", Tags: []string{"duplicate"}}

	err := service.Tag("model", d, "duplicate")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")
	mockRepo.AssertNotCalled(t, "Save")
}

func TestGetBody_DelegatesToRepo(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	mockRepo.On("LoadBody", "model", "0001").Return("body content", nil)

	body, err := service.GetBody("model", "0001")

	assert.NoError(t, err)
	assert.Equal(t, "body content", body)
	mockRepo.AssertExpectations(t)
}

func TestGetBody_PropagatesError(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	mockRepo.On("LoadBody", "model", "0001").Return("", errors.New("not found"))

	body, err := service.GetBody("model", "0001")

	assert.Error(t, err)
	assert.Empty(t, body)
}

func TestFilterDecisions_ByID(t *testing.T) {
	service := &DecisionServiceImplementation{}
	decisions := []Decision{
		{ID: "0001"},
		{ID: "0002"},
		{ID: "0003"},
	}

	filters := map[string][]string{"id": {"0001,0003"}}

	filtered, err := service.FilterDecisions(decisions, filters)
	assert.NoError(t, err)
	assert.Len(t, filtered, 2)
	assert.Equal(t, "0001", filtered[0].ID)
	assert.Equal(t, "0003", filtered[1].ID)
}

func TestFilterDecisions_ByTitle(t *testing.T) {
	service := &DecisionServiceImplementation{}
	decisions := []Decision{
		{ID: "0001", Title: "Use Kafka"},
		{ID: "0002", Title: "Migrate to gRPC"},
	}

	filtered, err := service.FilterDecisions(decisions, map[string][]string{"title": {"Kafka"}})
	assert.NoError(t, err)
	assert.Len(t, filtered, 1)
	assert.Equal(t, "0001", filtered[0].ID)
}

func TestFilterDecisions_ByTag(t *testing.T) {
	service := &DecisionServiceImplementation{}
	decisions := []Decision{
		{ID: "1", Tags: []string{"infra", "backend"}},
		{ID: "2", Tags: []string{"frontend"}},
		{ID: "3", Tags: []string{"infra"}},
	}

	filtered, err := service.FilterDecisions(decisions, map[string][]string{"tag": {"infra"}})
	assert.NoError(t, err)
	assert.Len(t, filtered, 2)
}

func TestFilterDecisions_ByStatus(t *testing.T) {
	service := &DecisionServiceImplementation{}
	decisions := []Decision{
		{ID: "1", Status: "proposed"},
		{ID: "2", Status: "accepted"},
		{ID: "3", Status: "proposed"},
	}

	filtered, err := service.FilterDecisions(decisions, map[string][]string{"status": {"proposed"}})
	assert.NoError(t, err)
	assert.Len(t, filtered, 2)
}

func TestFilterDecisions_InvalidTitleRegex(t *testing.T) {
	service := &DecisionServiceImplementation{}
	_, err := service.FilterDecisions([]Decision{{ID: "1"}}, map[string][]string{"title": {"*["}})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid title regex")
}

func TestFilterDecisions_InvalidIDRange(t *testing.T) {
	service := &DecisionServiceImplementation{}
	_, err := service.FilterDecisions([]Decision{{ID: "1"}}, map[string][]string{"id": {"0010-0005"}})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid ID range")
}

func TestCopy_DelegatesToRepo(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	mockRepo.On("Copy", "source", "target", "0042").Return(nil)

	err := service.Copy("source", "target", "0042")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
