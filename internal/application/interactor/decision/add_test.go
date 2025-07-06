package decision

import (
	decisiondomain "adg/internal/domain/decision"
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd_ModelDoesNotExist(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionAdd)

	modelPath := "nonexistent"
	mockModelSvc.On("Exists", modelPath).Return(false)

	interactor := NewAddDecisionsInteractor(mockModelSvc, mockDecisionSvc, mockOutput)

	// Act/When
	err := interactor.Add(modelPath, []string{"Decision A"})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "can not add decisions")
	mockModelSvc.AssertExpectations(t)
	mockOutput.AssertNotCalled(t, "Added", mock.Anything, mock.Anything)
}

func TestAdd_AllSuccess(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionAdd)

	modelPath := "model"
	titles := []string{"Decision A", "Decision B"}

	mockModelSvc.On("Exists", modelPath).Return(true)
	mockDecisionSvc.On("AddNew", modelPath, "Decision A").Return(&decisiondomain.Decision{Title: "Decision A"}, nil)
	mockDecisionSvc.On("AddNew", modelPath, "Decision B").Return(&decisiondomain.Decision{Title: "Decision B"}, nil)
	mockOutput.On("Added", mock.AnythingOfType("[]*decision.Decision"), mock.AnythingOfType("map[string]error")).Return()

	interactor := NewAddDecisionsInteractor(mockModelSvc, mockDecisionSvc, mockOutput)

	err := interactor.Add(modelPath, titles)

	assert.NoError(t, err)
	mockModelSvc.AssertExpectations(t)
	mockDecisionSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestAdd_PartialFailures(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionAdd)

	modelPath := "model"
	titles := []string{"Decision A", "Decision B"}
	addErr := errors.New("invalid title")

	mockModelSvc.On("Exists", modelPath).Return(true)
	mockDecisionSvc.On("AddNew", modelPath, "Decision A").Return(nil, addErr)
	mockDecisionSvc.On("AddNew", modelPath, "Decision B").Return(&decisiondomain.Decision{Title: "Decision B"}, nil)
	mockOutput.On("Added", mock.Anything, mock.Anything).Return()

	interactor := NewAddDecisionsInteractor(mockModelSvc, mockDecisionSvc, mockOutput)

	err := interactor.Add(modelPath, titles)

	assert.NoError(t, err)
	mockModelSvc.AssertExpectations(t)
	mockDecisionSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}
