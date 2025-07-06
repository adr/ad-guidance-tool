package decision

import (
	"adg/internal/domain/decision"
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTag_Success(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionTag)

	modelPath := "model"
	decision := &decision.Decision{ID: "0012"}
	tags := []string{"critical", "backend"}

	mockService.On("GetDecisionByID", modelPath, "0012").Return(decision, nil)
	mockService.On("Tag", modelPath, decision, "critical").Return(nil)
	mockService.On("Tag", modelPath, decision, "backend").Return(nil)
	mockOutput.On("Tagged", "0012", tags).Return()

	interactor := NewTagDecisionInteractor(mockService, mockOutput)
	err := interactor.Tag(modelPath, "0012", "", tags)

	assert.NoError(t, err)
	mockService.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestTag_ResolveError(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionTag)

	modelPath := "model"
	tags := []string{"urgent"}

	mockService.On("GetDecisionByID", modelPath, "0012").Return(nil, errors.New("not found"))

	interactor := NewTagDecisionInteractor(mockService, mockOutput)
	err := interactor.Tag(modelPath, "0012", "", tags)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
	mockService.AssertExpectations(t)
}

func TestTag_TagFails(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionTag)

	modelPath := "model"
	decision := &decision.Decision{ID: "0012"}
	tags := []string{"duplicate"}

	mockService.On("GetDecisionByID", modelPath, "0012").Return(decision, nil)
	mockService.On("Tag", modelPath, decision, "duplicate").Return(errors.New("already exists"))

	interactor := NewTagDecisionInteractor(mockService, mockOutput)
	err := interactor.Tag(modelPath, "0012", "", tags)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to tag decision")
	mockService.AssertExpectations(t)
}
