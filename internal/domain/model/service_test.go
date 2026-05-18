package model

import (
	"errors"
	"strings"
	"testing"

	"adg/internal/domain/decision"

	"github.com/stretchr/testify/assert"
)

func TestExists_DelegatesToRepo(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	service := NewModelService(mockModelRepo, mockDecisionRepo)

	mockModelRepo.On("Exists", "model").Return(true).Once()
	assert.True(t, service.Exists("model"))

	mockModelRepo.On("Exists", "missing").Return(false).Once()
	assert.False(t, service.Exists("missing"))

	mockModelRepo.AssertExpectations(t)
}

func TestCreateModel_DelegatesToRepo(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	service := NewModelService(mockModelRepo, mockDecisionRepo)

	mockModelRepo.On("CreateModel", "model").Return(nil)

	err := service.CreateModel("model")
	assert.NoError(t, err)
	mockModelRepo.AssertExpectations(t)
}

func TestCreateModel_PropagatesError(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	service := NewModelService(mockModelRepo, mockDecisionRepo)

	mockModelRepo.On("CreateModel", "model").Return(errors.New("disk full"))

	err := service.CreateModel("model")
	assert.ErrorContains(t, err, "disk full")
}

// TestValidate_FilenameMismatch verifies rule 1: filenames must match NNNN-slug.md.
func TestValidate_FilenameMismatch(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)
	service := NewModelService(mockModelRepo, mockDecisionRepo)

	mockDecisionRepo.On("LoadAll", "model").Return([]decision.Decision{
		{ID: "abc", Title: "T", Status: "proposed"},
	}, nil)
	mockDecisionRepo.On("LoadBody", "model", "abc").Return(validBody("T"), nil)

	issues, err := service.Validate("model")

	assert.NoError(t, err)
	assert.NotEmpty(t, issues)
	assert.Equal(t, "abc", issues[0].ID)
	assert.Contains(t, issues[0].Message, "filename does not match NNNN-slug.md")
}

// TestValidate_SupersessionReverseIntegrity verifies rule 8: a Decision listed
// in another's Supersedes must have a status that points back at the successor.
func TestValidate_SupersessionReverseIntegrity(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)
	service := NewModelService(mockModelRepo, mockDecisionRepo)

	// ADR-0002 claims to supersede ADR-0001, but ADR-0001's status doesn't
	// agree (it's still "proposed" instead of "superseded by ADR-0002").
	decisions := []decision.Decision{
		{ID: "0001", Title: "First", Status: "proposed"},
		{ID: "0002", Title: "Second", Status: "accepted", Supersedes: []string{"0001"}, LegacyOutcome: true},
	}
	mockDecisionRepo.On("LoadAll", "model").Return(decisions, nil)
	mockDecisionRepo.On("LoadBody", "model", "0001").Return(validBody("First"), nil)
	mockDecisionRepo.On("LoadBody", "model", "0002").Return(validBody("Second"), nil)

	issues, err := service.Validate("model")

	assert.NoError(t, err)
	found := false
	for _, iss := range issues {
		if iss.ID == "0002" && containsAll(iss.Message, "supersedes 0001", "proposed") {
			found = true
		}
	}
	assert.True(t, found, "expected supersession reverse-integrity issue, got: %#v", issues)
}

// TestValidate_SupersessionForwardIntegrity verifies rule 7: status "superseded
// by ADR-X" implies ADR-X exists and has self in its Supersedes list.
func TestValidate_SupersessionForwardIntegrity(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)
	service := NewModelService(mockModelRepo, mockDecisionRepo)

	// ADR-0001 says it's been superseded by ADR-0002, but ADR-0002 doesn't
	// list 0001 in its Supersedes.
	decisions := []decision.Decision{
		{ID: "0001", Title: "First", Status: "superseded by ADR-0002"},
		{ID: "0002", Title: "Second", Status: "accepted", LegacyOutcome: true},
	}
	mockDecisionRepo.On("LoadAll", "model").Return(decisions, nil)
	mockDecisionRepo.On("LoadBody", "model", "0001").Return(validBody("First"), nil)
	mockDecisionRepo.On("LoadBody", "model", "0002").Return(validBody("Second"), nil)

	issues, err := service.Validate("model")

	assert.NoError(t, err)
	found := false
	for _, iss := range issues {
		if iss.ID == "0001" && containsAll(iss.Message, "ADR-0002", "supersedes list does not include 0001") {
			found = true
		}
	}
	assert.True(t, found, "expected supersession forward-integrity issue, got: %#v", issues)
}

// TestValidate_PlaceholderNumericCommentText defends §A.1: numeric comment text
// indicates an unrecovered legacy placeholder.
func TestValidate_PlaceholderNumericCommentText(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)
	service := NewModelService(mockModelRepo, mockDecisionRepo)

	decisions := []decision.Decision{
		{
			ID: "0001", Title: "T", Status: "proposed",
			Comments: []decision.Comment{
				{Author: "Jane", Date: "2026-01-01 12:00:00", Text: "1"},
			},
		},
	}
	mockDecisionRepo.On("LoadAll", "model").Return(decisions, nil)
	mockDecisionRepo.On("LoadBody", "model", "0001").Return(validBody("T"), nil)

	issues, err := service.Validate("model")

	assert.NoError(t, err)
	found := false
	for _, iss := range issues {
		if iss.ID == "0001" && containsAll(iss.Message, "comment 1", "placeholder") {
			found = true
		}
	}
	assert.True(t, found, "expected placeholder-comment issue, got: %#v", issues)
}

// TestValidate_AllValid_NoIssues ensures a well-formed decision produces zero issues.
func TestValidate_AllValid_NoIssues(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)
	service := NewModelService(mockModelRepo, mockDecisionRepo)

	decisions := []decision.Decision{
		{ID: "0001", Title: "T", Status: "proposed"},
	}
	mockDecisionRepo.On("LoadAll", "model").Return(decisions, nil)
	mockDecisionRepo.On("LoadBody", "model", "0001").Return(validBody("T"), nil)

	issues, err := service.Validate("model")

	assert.NoError(t, err)
	assert.Empty(t, issues)
}

func TestValidate_PropagatesLoadError(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)
	service := NewModelService(mockModelRepo, mockDecisionRepo)

	mockDecisionRepo.On("LoadAll", "model").Return(nil, errors.New("read error"))

	_, err := service.Validate("model")
	assert.ErrorContains(t, err, "read error")
}

// validBody returns a MADR-shaped body containing the three required sections,
// at least one bullet under Considered Options, and the matching H1 title.
func validBody(title string) string {
	return "# " + title + `

## Context and Problem Statement

Some context.

## Considered Options

* Option A

## Decision Outcome

{...}
`
}

func containsAll(s string, substrs ...string) bool {
	for _, sub := range substrs {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}
