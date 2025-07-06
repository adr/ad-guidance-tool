package decision

import (
	"adg/internal/domain"
	"errors"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddNew_ValidTitle(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/path"
	title := "Create something"

	expectedDecision := &Decision{
		Title:    title,
		Status:   "open",
		Tags:     []string{},
		Links:    Links{Precedes: []string{}, Succeeds: []string{}},
		Comments: []Comment{},
	}
	expectedContent := &DecisionContent{}

	mockRepo.On("Create", modelPath, "", mock.Anything, expectedContent).
		Return(expectedDecision, nil)

	result, err := service.AddNew(modelPath, title)

	assert.NoError(t, err)
	assert.Equal(t, expectedDecision.Title, result.Title)
	assert.Equal(t, expectedDecision.Status, result.Status)
	mockRepo.AssertExpectations(t)
}

func TestAddNew_InvalidTitle(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/path"
	title := "12345 !!!"

	result, err := service.AddNew(modelPath, title)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "title must contain at least one letter")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestAddExisting_Success(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	sourcePath := "source/model"
	targetPath := "target/model"
	increment := 10

	decision := &Decision{
		ID:    "0012",
		Title: "Test",
		Links: Links{
			Precedes: []string{"0010"},
			Succeeds: []string{"0008"},
			Custom: map[string][]string{
				"relates": {"0005", "x"}, // x should be preserved
			},
		},
	}
	content := &DecisionContent{}

	// stub out path lookup
	mockRepo.On("FindDecisionFile", sourcePath, decision.ID).
		Return(filepath.Join(sourcePath, "0012", "index.md"), nil)

	expectedDecision := &Decision{
		ID:    "0012",
		Title: "Test",
		Links: Links{
			Precedes: []string{"0020"},
			Succeeds: []string{"0018"},
			Custom: map[string][]string{
				"relates": {"0015", "x"},
			},
		},
	}

	mockRepo.On("Create", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(expectedDecision, nil)

	result, err := service.AddExisting(sourcePath, targetPath, decision, content, increment)

	assert.NoError(t, err)
	assert.Equal(t, "0020", result.Links.Precedes[0])
	assert.Equal(t, "0018", result.Links.Succeeds[0])
	assert.Equal(t, []string{"0015", "x"}, result.Links.Custom["relates"])

	mockRepo.AssertExpectations(t)
}

func TestAddExisting_SubFolderPathError(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	sourcePath := "source"
	targetPath := "target"
	decision := &Decision{ID: "0012"}
	content := &DecisionContent{}

	mockRepo.On("FindDecisionFile", sourcePath, decision.ID).
		Return("", errors.New("file not found"))

	result, err := service.AddExisting(sourcePath, targetPath, decision, content, 5)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "file not found")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestGetAllDecisions_PrimarySuccess(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "model/path"
	expected := []Decision{{ID: "001", Title: "T"}}

	mockRepo.On("LoadAllByIndex", modelPath).Return(expected, nil)

	result, err := service.GetAllDecisions(modelPath)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertCalled(t, "LoadAllByIndex", modelPath)
	mockRepo.AssertNotCalled(t, "LoadAllByData", modelPath)
}

func TestGetAllDecisions_FallbackSuccess(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "model/path"
	expected := []Decision{{ID: "002", Title: "Fallback"}}

	mockRepo.On("LoadAllByIndex", modelPath).Return(nil, errors.New("index error"))
	mockRepo.On("LoadAllByData", modelPath).Return(expected, nil)

	result, err := service.GetAllDecisions(modelPath)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllDecisions_BothFail(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "model/path"

	mockRepo.On("LoadAllByIndex", modelPath).Return(nil, errors.New("index error"))
	mockRepo.On("LoadAllByData", modelPath).Return(nil, errors.New("data error"))

	result, err := service.GetAllDecisions(modelPath)

	assert.Nil(t, result)
	assert.EqualError(t, err, "data error")
	mockRepo.AssertExpectations(t)
}

func TestGetDecisionByID(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "model"
	id := "001"
	expected := &Decision{ID: id, Title: "Decision"}

	mockRepo.On("LoadById", modelPath, id).Return(expected, nil)

	result, err := service.GetDecisionByID(modelPath, id)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetDecisionByTitle(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "model"
	title := "Important"
	expected := &Decision{ID: "123", Title: title}

	mockRepo.On("LoadByTitle", modelPath, title).Return(expected, nil)

	result, err := service.GetDecisionByTitle(modelPath, title)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetDecisionContent(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "model"
	id := "001"
	expected := &DecisionContent{ID: id, Question: "What?"}

	mockRepo.On("LoadDecisionContent", modelPath, id).Return(expected, nil)

	result, err := service.GetDecisionContent(modelPath, id)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestLink_PrecedesSavesBoth(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	source := &Decision{ID: "001"}
	target := &Decision{ID: "002"}

	mockRepo.On("LoadById", "model/path", "002").Return(target, nil) // used for cycle check
	mockRepo.On("Save", "model/path", source).Return(nil)
	mockRepo.On("Save", "model/path", target).Return(nil)

	err := service.Link("model/path", source, target, "precedes", "succeeds")

	assert.NoError(t, err)
	assert.Equal(t, []string{"002"}, source.Links.Precedes)
	assert.Equal(t, []string{"001"}, target.Links.Succeeds)
	mockRepo.AssertExpectations(t)
}

func TestLink_CustomTagsWithReverse(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	source := &Decision{ID: "D1", Links: Links{Custom: map[string][]string{}}}
	target := &Decision{ID: "D2", Links: Links{}} // target.Links.Custom nil on purpose

	mockRepo.On("Save", "model/path", source).Return(nil)
	mockRepo.On("Save", "model/path", target).Return(nil)

	err := service.Link("model/path", source, target, "relates", "linked-back")

	assert.NoError(t, err)
	assert.Equal(t, []string{"D2"}, source.Links.Custom["relates"])
	assert.Equal(t, []string{"D1"}, target.Links.Custom["linked-back"])
	mockRepo.AssertExpectations(t)
}

func TestLink_CustomTagsWithoutReverse(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	source := &Decision{ID: "X", Links: Links{Custom: map[string][]string{}}}
	target := &Decision{ID: "Y"} // no reverseTag

	mockRepo.On("Save", "model/path", source).Return(nil)
	mockRepo.On("Save", "model/path", target).Return(nil)

	err := service.Link("model/path", source, target, "custom-tag", "")

	assert.NoError(t, err)
	assert.Equal(t, []string{"Y"}, source.Links.Custom["custom-tag"])
	assert.Nil(t, target.Links.Custom[""])
	mockRepo.AssertExpectations(t)
}

func TestLink_CycleDetected(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	source := &Decision{ID: "A"}
	target := &Decision{
		ID:    "B",
		Links: Links{Precedes: []string{"A"}}, // Simulate B → A
	}

	mockRepo.On("LoadById", "model/path", "B").Return(target, nil)

	err := service.Link("model/path", source, target, "precedes", "succeeds")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "would create a cycle")
}

func TestLink_SaveSourceFails(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	source := &Decision{ID: "001", Links: Links{Custom: map[string][]string{}}}
	target := &Decision{ID: "002", Links: Links{Custom: map[string][]string{}}}

	mockRepo.On("Save", "model/path", source).Return(errors.New("save failed"))

	err := service.Link("model/path", source, target, "relates", "back")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to save source")
}

func TestLink_SaveTargetFails(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	source := &Decision{ID: "001", Links: Links{Custom: map[string][]string{}}}
	target := &Decision{ID: "002", Links: Links{Custom: map[string][]string{}}}

	mockRepo.On("Save", "model/path", source).Return(nil)
	mockRepo.On("Save", "model/path", target).Return(errors.New("target save failed"))

	err := service.Link("model/path", source, target, "relates", "back")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to save target")
}

func TestTag_AddsNewTag(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	decision := &Decision{
		ID:   "1234",
		Tags: []string{"existing"},
	}

	mockRepo.On("Save", "model/path", decision).Return(nil)

	err := service.Tag("model/path", decision, "new-tag")

	assert.NoError(t, err)
	assert.Contains(t, decision.Tags, "new-tag")
	mockRepo.AssertExpectations(t)
}

func TestTag_DuplicateTagFails(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	decision := &Decision{
		ID:   "1234",
		Tags: []string{"duplicate"},
	}

	err := service.Tag("model/path", decision, "duplicate")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")
	mockRepo.AssertNotCalled(t, "Save")
}

func TestTag_SaveFails(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	decision := &Decision{
		ID:   "9999",
		Tags: []string{},
	}

	mockRepo.On("Save", "model/path", decision).Return(errors.New("db failure"))

	err := service.Tag("model/path", decision, "crash")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to save")
	mockRepo.AssertExpectations(t)
}

func TestFilterDecisions_ByID(t *testing.T) {
	service := &DecisionServiceImplementation{}
	decisions := []Decision{
		{ID: "0001"},
		{ID: "0002"},
		{ID: "0003"},
	}

	filters := map[string][]string{
		"id": {"0001,0003"},
	}

	filtered, err := service.FilterDecisions(decisions, filters)
	assert.NoError(t, err)
	assert.Len(t, filtered, 2)
	assert.Equal(t, "0001", filtered[0].ID)
	assert.Equal(t, "0003", filtered[1].ID)
}

func TestFilterDecisions_ByTitle(t *testing.T) {
	service := &DecisionServiceImplementation{}
	decisions := []Decision{
		{ID: "001", Title: "Use Kafka"},
		{ID: "002", Title: "Migrate to gRPC"},
		{ID: "003", Title: "Deprecate SOAP"},
	}

	filters := map[string][]string{
		"title": {"Kafka"},
	}

	filtered, err := service.FilterDecisions(decisions, filters)
	assert.NoError(t, err)
	assert.Len(t, filtered, 1)
	assert.Equal(t, "001", filtered[0].ID)
}

func TestFilterDecisions_ByTag(t *testing.T) {
	service := &DecisionServiceImplementation{}
	decisions := []Decision{
		{ID: "1", Tags: []string{"infra", "backend"}},
		{ID: "2", Tags: []string{"frontend"}},
		{ID: "3", Tags: []string{"infra"}},
	}

	filters := map[string][]string{
		"tag": {"infra"},
	}

	filtered, err := service.FilterDecisions(decisions, filters)
	assert.NoError(t, err)
	assert.Len(t, filtered, 2)
	assert.Equal(t, "1", filtered[0].ID)
	assert.Equal(t, "3", filtered[1].ID)
}

func TestFilterDecisions_ByStatus(t *testing.T) {
	service := &DecisionServiceImplementation{}
	decisions := []Decision{
		{ID: "1", Status: "open"},
		{ID: "2", Status: "decided"},
		{ID: "3", Status: "open"},
	}

	filters := map[string][]string{
		"status": {"open"},
	}

	filtered, err := service.FilterDecisions(decisions, filters)
	assert.NoError(t, err)
	assert.Len(t, filtered, 2)
	assert.Equal(t, "1", filtered[0].ID)
	assert.Equal(t, "3", filtered[1].ID)
}

func TestFilterDecisions_InvalidTitleRegex(t *testing.T) {
	service := &DecisionServiceImplementation{}
	decisions := []Decision{{ID: "1", Title: "valid"}}

	filters := map[string][]string{
		"title": {"*["}, // invalid regex
	}

	filtered, err := service.FilterDecisions(decisions, filters)
	assert.Nil(t, filtered)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid title regex")
}

func TestFilterDecisions_InvalidIDRange(t *testing.T) {
	service := &DecisionServiceImplementation{}
	decisions := []Decision{{ID: "1"}}

	filters := map[string][]string{
		"id": {"0010-0005"}, // invalid range
	}

	filtered, err := service.FilterDecisions(decisions, filters)
	assert.Nil(t, filtered)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid ID range")
}

func TestFilterDecisions_MultipleMatches(t *testing.T) {
	service := &DecisionServiceImplementation{}
	decisions := []Decision{
		{ID: "0001", Title: "Kafka", Tags: []string{"data"}, Status: "open"},
		{ID: "0002", Title: "Redis", Tags: []string{"infra"}, Status: "closed"},
	}

	filters := map[string][]string{
		"id":     {"0001"},
		"tag":    {"infra"},
		"status": {"closed"},
	}

	filtered, err := service.FilterDecisions(decisions, filters)
	assert.NoError(t, err)
	assert.Len(t, filtered, 2)
}

func TestDecide_ExistingOption(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "path"
	decision := &Decision{ID: "001", Status: "open"}
	option := "Option A"
	rationale := "it’s the best fit"

	mockRepo.On("OptionExists", modelPath, decision.ID, option).Return(true, nil)
	mockRepo.On("ResolveOptionNumber", modelPath, decision.ID, option).Return(1, nil)
	mockRepo.On("AppendOutcomeSection", modelPath, decision.ID, "We decided for [Option 1](#option-1) because: it’s the best fit").Return(nil)
	mockRepo.On("Save", modelPath, decision).Return(nil)

	err := service.Decide(modelPath, decision, option, rationale, false)
	assert.NoError(t, err)
	assert.Equal(t, "decided", decision.Status)
	mockRepo.AssertExpectations(t)
}

func TestDecide_NonExistentOption_WithoutEnforce(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "path"
	decision := &Decision{ID: "001"}
	option := "non-existent"

	mockRepo.On("OptionExists", modelPath, decision.ID, option).Return(false, nil)

	err := service.Decide(modelPath, decision, option, "", false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "option does not exist")
	mockRepo.AssertExpectations(t)
}

func TestDecide_NumericOptionWithForce(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "path"
	decision := &Decision{ID: "001"}
	option := "2"

	mockRepo.On("OptionExists", modelPath, decision.ID, option).Return(false, nil)

	err := service.Decide(modelPath, decision, option, "", true)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot auto-create numeric option")
	mockRepo.AssertExpectations(t)
}

func TestDecide_AppendOutcomeFails(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "path"
	decision := &Decision{ID: "001"}
	option := "Option A"

	mockRepo.On("OptionExists", modelPath, decision.ID, option).Return(true, nil)
	mockRepo.On("ResolveOptionNumber", modelPath, decision.ID, option).Return(1, nil)
	mockRepo.On("AppendOutcomeSection", modelPath, decision.ID, mock.Anything).Return(errors.New("write error"))

	err := service.Decide(modelPath, decision, option, "", false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "write error")
	mockRepo.AssertExpectations(t)
}

func TestRevise_Success(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	original := &Decision{
		ID:    "0001",
		Title: "Make a choice",
		Tags:  []string{"important"},
	}

	content := &DecisionContent{
		Outcome:  "Some outcome",
		Comments: "Some comments",
	}

	expectedRevised := &Decision{
		Title:  "Make a choice (Revised)",
		Status: "open",
		Tags:   []string{"important"},
	}

	expectedContent := &DecisionContent{
		Outcome:  "",
		Comments: "",
	}

	mockRepo.On("LoadDecisionContent", modelPath, original.ID).Return(content, nil)
	mockRepo.On("FindDecisionFile", modelPath, original.ID).
		Return(filepath.Join(modelPath, original.ID, "index.md"), nil)
	mockRepo.On("Create", modelPath, original.ID, mock.Anything, expectedContent).
		Return(expectedRevised, nil)

	revised, err := service.Revise(modelPath, original)

	assert.NoError(t, err)
	assert.Equal(t, "Make a choice (Revised)", revised.Title)
	assert.Equal(t, "open", revised.Status)
	assert.Equal(t, []string{"important"}, revised.Tags)

	mockRepo.AssertExpectations(t)
}

func TestRevise_LoadContentError(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	original := &Decision{ID: "0001"}

	mockRepo.On("LoadDecisionContent", modelPath, original.ID).
		Return(nil, errors.New("failed to load"))

	revised, err := service.Revise(modelPath, original)

	assert.Nil(t, revised)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load original content")
}

func TestRevise_GetSubFolderPathError(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	original := &Decision{ID: "0001", Title: "Test"}

	content := &DecisionContent{}

	mockRepo.On("LoadDecisionContent", modelPath, original.ID).Return(content, nil)
	mockRepo.On("FindDecisionFile", modelPath, original.ID).
		Return("", errors.New("not found"))

	revised, err := service.Revise(modelPath, original)

	assert.Nil(t, revised)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get relative path")
}

func TestCopy_Success(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "source/model"
	targetPath := "target/model"
	decisionID := "0042"

	mockRepo.On("Copy", modelPath, targetPath, decisionID).Return(nil)

	err := service.Copy(modelPath, targetPath, decisionID)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Copy", modelPath, targetPath, decisionID)
}

func TestComment_Success(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	author := "Jane"
	commentText := "Looks good"

	decision := &Decision{
		ID:       "0001",
		Comments: []Comment{},
	}

	mockRepo.On("Save", modelPath, decision).Return(nil)
	mockRepo.On("AppendCommentSection", modelPath, decision.ID, commentText, 1, author, mock.AnythingOfType("string")).
		Return(nil)

	err := service.Comment(modelPath, decision, author, commentText)

	assert.NoError(t, err)
	assert.Len(t, decision.Comments, 1)
	assert.Equal(t, author, decision.Comments[0].Author)
	assert.Equal(t, "1", decision.Comments[0].Comment)

	mockRepo.AssertExpectations(t)
}

func TestComment_SaveFails(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	author := "Jane"
	commentText := "Some comment"

	decision := &Decision{
		ID:       "0001",
		Comments: []Comment{},
	}

	mockRepo.On("Save", modelPath, decision).Return(errors.New("save failed"))

	err := service.Comment(modelPath, decision, author, commentText)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "save failed")
	mockRepo.AssertNotCalled(t, "AppendCommentSection")
}

func TestComment_FailsToAppendSection(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	decision := &Decision{
		ID:       "0001",
		Comments: []Comment{},
	}
	author := "TestUser"
	commentText := "This is a comment"

	mockRepo.On("Save", modelPath, decision).Return(nil)

	mockRepo.On("AppendCommentSection", modelPath, decision.ID, commentText, 1, author, mock.AnythingOfType("string")).
		Return(errors.New("append error"))

	err := service.Comment(modelPath, decision, author, commentText)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to append comment section")
	mockRepo.AssertExpectations(t)
}

func Test_appendToSection(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	decisionID := "0001"
	section := "question"
	existing := "What is the best option?"
	newContent := "Why is it the best?"

	expected := []string{"What is the best option?", "Why is it the best?"}

	mockRepo.On("UpdateSection", modelPath, decisionID, section, expected).Return(nil)

	err := service.(*DecisionServiceImplementation).appendToSection(modelPath, decisionID, section, existing, newContent)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_appendQuestion_Success(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	decisionID := "0001"
	existing := "Existing Q"
	newQ := "Follow-up?"

	content := &DecisionContent{Question: existing}
	combined := []string{"Existing Q", "Follow-up?"}

	mockRepo.On("LoadDecisionContent", modelPath, decisionID).Return(content, nil)
	mockRepo.On("UpdateSection", modelPath, decisionID, domain.AnchorSectionQuestion, combined).Return(nil)

	err := service.(*DecisionServiceImplementation).appendQuestion(modelPath, decisionID, newQ)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
func Test_appendCriteria_Success(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	decisionID := "0002"
	existing := "Must be scalable"
	newC := "And secure"

	content := &DecisionContent{Criteria: existing}
	combined := []string{"Must be scalable", "And secure"}

	mockRepo.On("LoadDecisionContent", modelPath, decisionID).Return(content, nil)
	mockRepo.On("UpdateSection", modelPath, decisionID, domain.AnchorSectionCriteria, combined).Return(nil)

	err := service.(*DecisionServiceImplementation).appendCriteria(modelPath, decisionID, newC)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
func Test_countExistingOptions(t *testing.T) {
	lines := []string{
		`1. <a name="option-1"></a> Option A`,
		`2. <a name="option-2"></a> Option B`,
		`Some unrelated line`,
	}
	count := countExistingOptions(lines)
	assert.Equal(t, 2, count)
}
func Test_validateOptionDoesNotExist_Success(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	decisionID := "0003"
	option := "New Option"

	mockRepo.On("OptionExists", modelPath, decisionID, option).Return(false, nil)

	err := service.(*DecisionServiceImplementation).validateOptionDoesNotExist(modelPath, decisionID, option)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_validateOptionDoesNotExist_Exists(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	decisionID := "0003"
	option := "Duplicate Option"

	mockRepo.On("OptionExists", modelPath, decisionID, option).Return(true, nil)

	err := service.(*DecisionServiceImplementation).validateOptionDoesNotExist(modelPath, decisionID, option)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")
}
func Test_formatOptionLine(t *testing.T) {
	option := "Faster delivery"
	result := formatOptionLine(3, option)
	expected := fmt.Sprintf("3. %s %s", domain.AnchorForOption(3), option)
	assert.Equal(t, expected, result)
}

func Test_appendOptions_Success(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	decisionID := "0001"
	newOptions := []string{"Option C", "Option D"}

	content := &DecisionContent{
		Options: `1. <a name="option-1"></a> Option A
2. <a name="option-2"></a> Option B`,
	}

	lines := []string{
		"1. <a name=\"option-1\"></a> Option A",
		"2. <a name=\"option-2\"></a> Option B",
		"3. <a name=\"option-3\"></a> Option C",
		"4. <a name=\"option-4\"></a> Option D",
	}

	mockRepo.On("LoadDecisionContent", modelPath, decisionID).Return(content, nil)
	mockRepo.On("OptionExists", modelPath, decisionID, "Option C").Return(false, nil)
	mockRepo.On("OptionExists", modelPath, decisionID, "Option D").Return(false, nil)
	mockRepo.On("UpdateSection", modelPath, decisionID, domain.AnchorSectionOptions, lines).Return(nil)

	err := service.(*DecisionServiceImplementation).appendOptions(modelPath, decisionID, newOptions)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
func Test_appendOptions_OptionAlreadyExists(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	decisionID := "0002"
	newOptions := []string{"Option X"}

	content := &DecisionContent{Options: ""}

	mockRepo.On("LoadDecisionContent", modelPath, decisionID).Return(content, nil)
	mockRepo.On("OptionExists", modelPath, decisionID, "Option X").Return(true, nil)

	err := service.(*DecisionServiceImplementation).appendOptions(modelPath, decisionID, newOptions)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")

	mockRepo.AssertExpectations(t)
}
func Test_appendOptions_LoadContentError(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	decisionID := "0003"
	newOptions := []string{"Option Y"}

	mockRepo.On("LoadDecisionContent", modelPath, decisionID).
		Return(nil, fmt.Errorf("load failed"))

	err := service.(*DecisionServiceImplementation).appendOptions(modelPath, decisionID, newOptions)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "load failed")
}
func Test_appendOptions_OptionExistsError(t *testing.T) {
	mockRepo := new(MockDecisionRepository)
	service := NewDecisionService(mockRepo)

	modelPath := "test/model"
	decisionID := "0004"
	newOptions := []string{"Option Z"}

	content := &DecisionContent{Options: ""}

	mockRepo.On("LoadDecisionContent", modelPath, decisionID).Return(content, nil)
	mockRepo.On("OptionExists", modelPath, decisionID, "Option Z").Return(false, fmt.Errorf("exist check error"))

	err := service.(*DecisionServiceImplementation).appendOptions(modelPath, decisionID, newOptions)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "exist check error")
}
