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

func TestCopy_TargetModelExists(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelCopy)

	interactor := NewCopyModelInteractor(mockModelSvc, mockDecisionSvc, mockOutput)

	mockModelSvc.On("Exists", "target").Return(true)

	err := interactor.Copy("source", "target", nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already contains a model")
	mockModelSvc.AssertExpectations(t)
}

func TestCopy_GetAllDecisionsFails(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelCopy)

	interactor := NewCopyModelInteractor(mockModelSvc, mockDecisionSvc, mockOutput)

	mockModelSvc.On("Exists", "target").Return(false)
	mockDecisionSvc.On("GetAllDecisions", "source").Return(nil, errors.New("load error"))

	err := interactor.Copy("source", "target", nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load decisions")
}

func TestCopy_FilterDecisionsFails(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelCopy)

	interactor := NewCopyModelInteractor(mockModelSvc, mockDecisionSvc, mockOutput)

	mockModelSvc.On("Exists", "target").Return(false)
	mockDecisionSvc.On("GetAllDecisions", "source").Return([]decision.Decision{}, nil)
	mockDecisionSvc.On("FilterDecisions", mock.Anything, mock.Anything).Return(nil, errors.New("filter error"))

	err := interactor.Copy("source", "target", map[string][]string{"tag": {"core"}})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to apply filters")
}

func TestCopy_CreateModelFails(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelCopy)

	interactor := NewCopyModelInteractor(mockModelSvc, mockDecisionSvc, mockOutput)

	mockModelSvc.On("Exists", "target").Return(false)
	mockDecisionSvc.On("GetAllDecisions", "source").Return([]decision.Decision{{ID: "001"}}, nil)
	mockModelSvc.On("CreateModel", "target").Return(errors.New("create error"))

	err := interactor.Copy("source", "target", nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create model directory")
}

func TestCopy_CopyDecisionFails(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelCopy)

	interactor := NewCopyModelInteractor(mockModelSvc, mockDecisionSvc, mockOutput)

	decisions := []decision.Decision{{ID: "001"}}

	mockModelSvc.On("Exists", "target").Return(false)
	mockDecisionSvc.On("GetAllDecisions", "source").Return(decisions, nil)
	mockModelSvc.On("CreateModel", "target").Return(nil)
	mockDecisionSvc.On("Copy", "source", "target", "001").Return(errors.New("copy error"))

	err := interactor.Copy("source", "target", nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to copy decision")
}

func TestCopy_RebuildIndexFails(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelCopy)

	interactor := NewCopyModelInteractor(mockModelSvc, mockDecisionSvc, mockOutput)

	decisions := []decision.Decision{{ID: "001"}}

	mockModelSvc.On("Exists", "target").Return(false)
	mockDecisionSvc.On("GetAllDecisions", "source").Return(decisions, nil)
	mockModelSvc.On("CreateModel", "target").Return(nil)
	mockDecisionSvc.On("Copy", "source", "target", "001").Return(nil)
	mockModelSvc.On("RebuildIndex", "target").Return(errors.New("index error"))

	err := interactor.Copy("source", "target", nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to rebuild index")
}

func TestCopy_Success(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelCopy)

	interactor := NewCopyModelInteractor(mockModelSvc, mockDecisionSvc, mockOutput)

	decisions := []decision.Decision{{ID: "001"}, {ID: "002"}}

	mockModelSvc.On("Exists", "target").Return(false)
	mockDecisionSvc.On("GetAllDecisions", "source").Return(decisions, nil)
	mockModelSvc.On("CreateModel", "target").Return(nil)
	mockDecisionSvc.On("Copy", "source", "target", "001").Return(nil)
	mockDecisionSvc.On("Copy", "source", "target", "002").Return(nil)
	mockModelSvc.On("RebuildIndex", "target").Return(nil)
	mockOutput.On("Copied", "source", "target", 2).Return()

	err := interactor.Copy("source", "target", nil)

	assert.NoError(t, err)
	mockOutput.AssertCalled(t, "Copied", "source", "target", 2)
}
