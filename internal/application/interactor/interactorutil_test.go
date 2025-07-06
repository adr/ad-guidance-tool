package interactor

import (
	"adg/internal/domain/decision"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveDecisionByIdOrTitle_ResolveByID(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	expected := &decision.Decision{ID: "0001"}

	mockSvc.On("GetDecisionByID", "test-model", "0001").Return(expected, nil)

	result, err := ResolveDecisionByIdOrTitle("test-model", "0001", "", mockSvc)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockSvc.AssertExpectations(t)
}

func TestResolveDecisionByIdOrTitle_ResolveByTitle(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	expected := &decision.Decision{ID: "0002", Title: "Test Decision"}

	mockSvc.On("GetDecisionByTitle", "test-model", "Test Decision").Return(expected, nil)

	result, err := ResolveDecisionByIdOrTitle("test-model", "", "Test Decision", mockSvc)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockSvc.AssertExpectations(t)
}

func TestResolveDecisionByIdOrTitle_ResolveByID_Error(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)

	mockSvc.On("GetDecisionByID", "test-model", "0001").Return(nil, errors.New("not found"))

	result, err := ResolveDecisionByIdOrTitle("test-model", "0001", "", mockSvc)

	assert.Nil(t, result)
	assert.ErrorContains(t, err, "not found")
	mockSvc.AssertExpectations(t)
}

func TestResolveDecisionByIdOrTitle_ResolveByTitle_Error(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)

	mockSvc.On("GetDecisionByTitle", "test-model", "Nonexistent").Return(nil, errors.New("not found"))

	result, err := ResolveDecisionByIdOrTitle("test-model", "", "Nonexistent", mockSvc)

	assert.Nil(t, result)
	assert.ErrorContains(t, err, "not found")
	mockSvc.AssertExpectations(t)
}
