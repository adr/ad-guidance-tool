package model

import (
	"errors"
	"strings"
	"testing"

	"adg/internal/domain"
	"adg/internal/domain/decision"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateModel_Success(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	service := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "test/path"

	// Expectations
	mockModelRepo.On("CreateModel", modelPath).Return(nil)
	mockModelRepo.On("CreateIndex", modelPath).Return(nil)

	// Test
	err := service.CreateModel(modelPath)

	assert.NoError(t, err)
	mockModelRepo.AssertCalled(t, "CreateModel", modelPath)
	mockModelRepo.AssertCalled(t, "CreateIndex", modelPath)
}

func TestCreateModel_CreateModelFails(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	service := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "test/path"
	expectedErr := errors.New("create model failed")

	// Expectations
	mockModelRepo.On("CreateModel", modelPath).Return(expectedErr)

	// Test
	err := service.CreateModel(modelPath)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	mockModelRepo.AssertCalled(t, "CreateModel", modelPath)
	mockModelRepo.AssertNotCalled(t, "CreateIndex", modelPath)
}

func TestCreateModel_CreateIndexFails(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	service := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "test/path"
	expectedErr := errors.New("create index failed")

	// Expectations
	mockModelRepo.On("CreateModel", modelPath).Return(nil)
	mockModelRepo.On("CreateIndex", modelPath).Return(expectedErr)

	// Test
	err := service.CreateModel(modelPath)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	mockModelRepo.AssertCalled(t, "CreateModel", modelPath)
	mockModelRepo.AssertCalled(t, "CreateIndex", modelPath)
}

func TestRebuildIndex_Success(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	service := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "some/path"
	decisions := []decision.Decision{
		{ID: "d1"},
		{ID: "d2"},
	}

	mockDecisionRepo.On("LoadAllByData", modelPath).Return(decisions, nil)
	mockModelRepo.On("RebuildIndex", modelPath, decisions).Return(nil)

	err := service.RebuildIndex(modelPath)

	assert.NoError(t, err)
	mockDecisionRepo.AssertCalled(t, "LoadAllByData", modelPath)
	mockModelRepo.AssertCalled(t, "RebuildIndex", modelPath, decisions)
}

func TestRebuildIndex_DuplicateID(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	service := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "some/path"
	decisions := []decision.Decision{
		{ID: "duplicate"},
		{ID: "duplicate"},
	}

	mockDecisionRepo.On("LoadAllByData", modelPath).Return(decisions, nil)

	err := service.RebuildIndex(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate decision ID")
	mockModelRepo.AssertNotCalled(t, "RebuildIndex", mock.Anything, mock.Anything)
}

func TestRebuildIndex_LoadDecisionsFails(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	service := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "some/path"
	expectedErr := errors.New("load failure")

	mockDecisionRepo.On("LoadAllByData", modelPath).Return(nil, expectedErr)

	err := service.RebuildIndex(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load decisions")
	mockModelRepo.AssertNotCalled(t, "RebuildIndex", mock.Anything, mock.Anything)
}

func TestExists_ReturnsTrue(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	service := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "some/path"

	mockModelRepo.On("Exists", modelPath).Return(true)

	exists := service.Exists(modelPath)

	assert.True(t, exists)
	mockModelRepo.AssertCalled(t, "Exists", modelPath)
}

func TestExists_ReturnsFalse(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	service := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "some/path"

	mockModelRepo.On("Exists", modelPath).Return(false)

	exists := service.Exists(modelPath)

	assert.False(t, exists)
	mockModelRepo.AssertCalled(t, "Exists", modelPath)
}

func TestValidateIndexDataCorrectness_Success(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	svc := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "some/path"

	data := []decision.Decision{
		{ID: "d1", Title: "T1", Status: "Open", Tags: []string{"tag"}, Links: decision.Links{}, Comments: nil},
	}
	index := []decision.Decision{
		{ID: "d1", Title: "T1", Status: "Open", Tags: []string{"tag"}, Links: decision.Links{}, Comments: nil},
	}

	mockDecisionRepo.On("LoadAllByData", modelPath).Return(data, nil)
	mockDecisionRepo.On("LoadAllByIndex", modelPath).Return(index, nil)

	err := svc.ValidateIndexDataCorrectness(modelPath)

	assert.NoError(t, err)
}

func TestValidateIndexDataCorrectness_IDMissingInData(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	service := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "some/path"

	data := []decision.Decision{} // no decisions in data
	index := []decision.Decision{
		{
			ID:       "d1",
			Title:    "T1",
			Status:   "Open",
			Tags:     []string{"tag"},
			Links:    decision.Links{},
			Comments: []decision.Comment{},
		},
	}

	mockDecisionRepo.On("LoadAllByData", modelPath).Return(data, nil)
	mockDecisionRepo.On("LoadAllByIndex", modelPath).Return(index, nil)

	err := service.ValidateIndexDataCorrectness(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation of metadata completed with mismatches")
}

func TestValidateIndexDataCorrectness_DuplicateIDInData(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	svc := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "some/path"
	data := []decision.Decision{
		{ID: "d1"},
		{ID: "d1"},
	}

	mockDecisionRepo.On("LoadAllByData", modelPath).Return(data, nil)

	err := svc.ValidateIndexDataCorrectness(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate decision ID found: d1")
}

func TestValidateIndexDataCorrectness_LoadAllByDataFails(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	svc := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "some/path"
	mockDecisionRepo.On("LoadAllByData", modelPath).Return(nil, errors.New("failed to read"))

	err := svc.ValidateIndexDataCorrectness(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load decisions from files")
}

func TestValidateIndexDataCorrectness_LoadAllByIndexFails(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	svc := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "some/path"
	mockDecisionRepo.On("LoadAllByData", modelPath).Return([]decision.Decision{{ID: "d1"}}, nil)
	mockDecisionRepo.On("LoadAllByIndex", modelPath).Return(nil, errors.New("index fail"))

	err := svc.ValidateIndexDataCorrectness(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load decisions from index")
}

func TestValidateIndexDataCorrectness_MetadataMismatch(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	svc := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "some/path"

	data := []decision.Decision{{ID: "d1", Title: "A"}}
	index := []decision.Decision{{ID: "d1", Title: "B"}}

	mockDecisionRepo.On("LoadAllByData", modelPath).Return(data, nil)
	mockDecisionRepo.On("LoadAllByIndex", modelPath).Return(index, nil)

	err := svc.ValidateIndexDataCorrectness(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation of metadata completed with mismatches")
}

func TestValidateIndexDataCorrectness_MissingInIndex(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)

	svc := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "some/path"

	data := []decision.Decision{{ID: "d1"}}
	index := []decision.Decision{}

	mockDecisionRepo.On("LoadAllByData", modelPath).Return(data, nil)
	mockDecisionRepo.On("LoadAllByIndex", modelPath).Return(index, nil)

	err := svc.ValidateIndexDataCorrectness(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation of metadata completed with mismatches")
}

func TestValidateDecisionDataCorrectness_Success(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)
	svc := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "test/path"
	content := strings.Join([]string{
		domain.AnchorForSection(domain.AnchorSectionQuestion),
		domain.AnchorForSection(domain.AnchorSectionOptions),
		domain.AnchorForSection(domain.AnchorSectionCriteria),
	}, "\n")

	decisions := []decision.Decision{
		{ID: "0002"},
		{ID: "0001"},
	}

	mockDecisionRepo.On("LoadAllByIndex", modelPath).Return(decisions, nil)
	mockDecisionRepo.On("LoadDecisionContentRaw", modelPath, "0001").Return(content, nil)
	mockDecisionRepo.On("LoadDecisionContentRaw", modelPath, "0002").Return(content, nil)

	err := svc.ValidateDecisionDataCorrectness(modelPath)

	assert.NoError(t, err)
}

func TestValidateDecisionDataCorrectness_MissingAnchors(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)
	svc := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "test/path"
	content := domain.AnchorForSection(domain.AnchorSectionQuestion) // missing options and criteria

	mockDecisionRepo.On("LoadAllByIndex", modelPath).Return([]decision.Decision{{ID: "d2"}}, nil)
	mockDecisionRepo.On("LoadDecisionContentRaw", modelPath, "d2").Return(content, nil)

	err := svc.ValidateDecisionDataCorrectness(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation of file contents completed with errors")
}

func TestValidateDecisionDataCorrectness_LoadIndexFails(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)
	svc := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "test/path"
	mockDecisionRepo.On("LoadAllByIndex", modelPath).Return(nil, errors.New("index failed"))

	err := svc.ValidateDecisionDataCorrectness(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load decisions from index")
}

func TestValidateDecisionDataCorrectness_LoadContentFails(t *testing.T) {
	mockModelRepo := new(MockModelRepository)
	mockDecisionRepo := new(decision.MockDecisionRepository)
	svc := NewModelService(mockModelRepo, mockDecisionRepo)

	modelPath := "test/path"
	mockDecisionRepo.On("LoadAllByIndex", modelPath).Return([]decision.Decision{{ID: "d3"}}, nil)
	mockDecisionRepo.On("LoadDecisionContentRaw", modelPath, "d3").Return("", errors.New("read error"))

	err := svc.ValidateDecisionDataCorrectness(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load decision content from files")
}
