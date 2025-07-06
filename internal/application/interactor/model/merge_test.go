package model

import (
	"adg/internal/domain/decision"
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMerge_TargetExists(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelMerge)

	mockModelSvc.On("Exists", "target").Return(true)

	interactor := NewMergeModelsInteractor(mockModelSvc, mockDecisionSvc, mockOutput)
	err := interactor.Merge("modelA", "modelB", "target", nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already contains a model")
}

func TestMerge_SuccessfulMerge(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelMerge)

	modelADecisions := []decision.Decision{
		{ID: "0002"},
	}
	modelBDecisions := []decision.Decision{
		{ID: "0001"},
	}
	content := &decision.DecisionContent{}

	filters := map[string][]string{"id": {"0001", "0002"}}

	mockModelSvc.On("Exists", "target").Return(false)
	mockModelSvc.On("CreateModel", "target").Return(nil)
	mockModelSvc.On("RebuildIndex", "target").Return(nil)

	mockDecisionSvc.On("GetAllDecisions", "modelA").Return(modelADecisions, nil)
	mockDecisionSvc.On("GetAllDecisions", "modelB").Return(modelBDecisions, nil)
	mockDecisionSvc.On("FilterDecisions", modelADecisions, filters).Return(modelADecisions, nil)
	mockDecisionSvc.On("FilterDecisions", modelBDecisions, filters).Return(modelBDecisions, nil)

	mockDecisionSvc.On("GetDecisionContent", "modelA", "0002").Return(content, nil)
	mockDecisionSvc.On("AddExisting", "modelA", "target", &modelADecisions[0], content, 0).Return(&decision.Decision{}, nil)

	mockDecisionSvc.On("GetDecisionContent", "modelB", "0001").Return(content, nil)
	mockDecisionSvc.On("AddExisting", "modelB", "target", &modelBDecisions[0], content, 2).Return(&decision.Decision{}, nil)

	mockOutput.On("Merged", "modelA", "modelB", "target", 2).Return(nil)

	interactor := NewMergeModelsInteractor(mockModelSvc, mockDecisionSvc, mockOutput)
	err := interactor.Merge("modelA", "modelB", "target", filters)

	assert.NoError(t, err)
	mockModelSvc.AssertExpectations(t)
	mockDecisionSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestMerge_TargetModelExists(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelMerge)

	mockModelSvc.On("Exists", "target").Return(true)

	interactor := NewMergeModelsInteractor(mockModelSvc, mockDecisionSvc, mockOutput)
	err := interactor.Merge("modelA", "modelB", "target", map[string][]string{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already contains a model")
}

func TestMerge_CreateModelFails(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelMerge)

	mockModelSvc.On("Exists", "target").Return(false)
	mockModelSvc.On("CreateModel", "target").Return(errors.New("disk error"))

	interactor := NewMergeModelsInteractor(mockModelSvc, mockDecisionSvc, mockOutput)
	err := interactor.Merge("modelA", "modelB", "target", map[string][]string{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create model")
}

func TestMerge_LoadModelAFails(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelMerge)

	mockModelSvc.On("Exists", "target").Return(false)
	mockModelSvc.On("CreateModel", "target").Return(nil)
	mockDecisionSvc.On("GetAllDecisions", "modelA").Return(nil, errors.New("read error"))

	interactor := NewMergeModelsInteractor(mockModelSvc, mockDecisionSvc, mockOutput)
	err := interactor.Merge("modelA", "modelB", "target", map[string][]string{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load decisions from model")
}

func TestMerge_InvalidIDInModelA(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelMerge)

	mockModelSvc.On("Exists", "target").Return(false)
	mockModelSvc.On("CreateModel", "target").Return(nil)
	mockDecisionSvc.On("GetAllDecisions", "modelA").Return([]decision.Decision{{ID: "bad-id"}}, nil)

	interactor := NewMergeModelsInteractor(mockModelSvc, mockDecisionSvc, mockOutput)
	err := interactor.Merge("modelA", "modelB", "target", map[string][]string{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid ID")
}

func TestMerge_CopyModelAFails(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelMerge)

	mockModelSvc.On("Exists", "target").Return(false)
	mockModelSvc.On("CreateModel", "target").Return(nil)
	mockDecisionSvc.On("GetAllDecisions", "modelA").Return([]decision.Decision{{ID: "0001"}}, nil)
	mockDecisionSvc.On("GetDecisionContent", "modelA", "0001").Return(nil, errors.New("no content"))

	interactor := NewMergeModelsInteractor(mockModelSvc, mockDecisionSvc, mockOutput)
	err := interactor.Merge("modelA", "modelB", "target", map[string][]string{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to merge decisions from model")
}

func TestMerge_CopyModelBFails(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelMerge)

	mockModelSvc.On("Exists", "target").Return(false)
	mockModelSvc.On("CreateModel", "target").Return(nil)
	mockDecisionSvc.On("GetAllDecisions", "modelA").Return([]decision.Decision{{ID: "0001"}}, nil)
	mockDecisionSvc.On("GetAllDecisions", "modelB").Return([]decision.Decision{{ID: "0002"}}, nil)
	mockDecisionSvc.On("GetDecisionContent", "modelA", "0001").Return(&decision.DecisionContent{}, nil)
	mockDecisionSvc.On("AddExisting", "modelA", "target", mock.Anything, mock.Anything, 0).Return(&decision.Decision{}, nil)
	mockDecisionSvc.On("GetDecisionContent", "modelB", "0002").Return(nil, errors.New("no content"))

	interactor := NewMergeModelsInteractor(mockModelSvc, mockDecisionSvc, mockOutput)
	err := interactor.Merge("modelA", "modelB", "target", map[string][]string{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to merge decisions from model")
}

func TestMerge_RebuildIndexFails(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelMerge)

	mockModelSvc.On("Exists", "target").Return(false)
	mockModelSvc.On("CreateModel", "target").Return(nil)
	mockDecisionSvc.On("GetAllDecisions", "modelA").Return([]decision.Decision{{ID: "0001"}}, nil)
	mockDecisionSvc.On("GetDecisionContent", "modelA", "0001").Return(&decision.DecisionContent{}, nil)
	mockDecisionSvc.On("AddExisting", "modelA", "target", mock.Anything, mock.Anything, 0).Return(&decision.Decision{}, nil)
	mockDecisionSvc.On("GetAllDecisions", "modelB").Return([]decision.Decision{}, nil)
	mockModelSvc.On("RebuildIndex", "target").Return(errors.New("write error"))

	interactor := NewMergeModelsInteractor(mockModelSvc, mockDecisionSvc, mockOutput)
	err := interactor.Merge("modelA", "modelB", "target", map[string][]string{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to rebuild index")
}
