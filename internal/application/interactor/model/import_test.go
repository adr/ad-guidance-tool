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

func TestImport_TargetModelDoesNotExist(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelImport)

	mockModelSvc.On("Exists", "target").Return(false)

	interactor := NewImportModelInteractor(mockModelSvc, mockDecisionSvc, mockOutput)
	err := interactor.Import("source", "target", nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "does not contain a model")
	mockModelSvc.AssertExpectations(t)
}

func TestImport_GetTargetDecisionsError(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelImport)

	mockModelSvc.On("Exists", "target").Return(true)
	mockDecisionSvc.On("GetAllDecisions", "target").Return(nil, errors.New("load error"))

	interactor := NewImportModelInteractor(mockModelSvc, mockDecisionSvc, mockOutput)
	err := interactor.Import("source", "target", nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load decisions from source model")
}

func TestImport_FilterAndAddSuccess(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.ModelImport)

	target := []decision.Decision{
		{ID: "0003"},
	}
	source := []decision.Decision{
		{ID: "0001"},
		{ID: "0002"},
	}
	content := &decision.DecisionContent{}

	mockModelSvc.On("Exists", "target").Return(true)
	mockDecisionSvc.On("GetAllDecisions", "target").Return(target, nil)
	mockDecisionSvc.On("GetAllDecisions", "source").Return(source, nil)
	mockDecisionSvc.On("FilterDecisions", source, mock.Anything).Return(source, nil)
	mockDecisionSvc.On("GetDecisionContent", "source", mock.Anything).Return(content, nil)
	mockDecisionSvc.On("AddExisting", "source", "target", mock.Anything, content, 3).Return(&decision.Decision{}, nil).Twice()
	mockModelSvc.On("RebuildIndex", "target").Return(nil)
	mockOutput.On("Imported", "source", "target", 2).Return(nil)

	interactor := NewImportModelInteractor(mockModelSvc, mockDecisionSvc, mockOutput)
	err := interactor.Import("source", "target", map[string][]string{"id": {"0001", "0002"}})

	assert.NoError(t, err)
	mockModelSvc.AssertExpectations(t)
	mockDecisionSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}
