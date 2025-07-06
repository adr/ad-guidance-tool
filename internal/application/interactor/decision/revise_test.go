package decision

import (
	"adg/internal/domain/decision"
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReviseDecision_Success(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionRevise)

	original := &decision.Decision{ID: "001", Title: "Decide A"}
	revised := &decision.Decision{ID: "002", Title: "Decide A (Revised)"}

	mockSvc.On("GetDecisionByID", "model", "001").Return(original, nil)
	mockSvc.On("Revise", "model", original).Return(revised, nil)
	mockSvc.On("Link", "model", original, revised, "revised by", "revises").Return(nil)
	mockOut.On("Revised", "001", "002").Return(nil)

	interactor := NewReviseDecisionInteractor(mockSvc, mockOut)
	err := interactor.ReviseDecision("model", "001", "")

	assert.NoError(t, err)
	mockSvc.AssertExpectations(t)
	mockOut.AssertExpectations(t)
}

func TestReviseDecision_FailsToFindOriginal(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionRevise)

	mockSvc.On("GetDecisionByID", "model", "999").Return(nil, errors.New("not found"))

	interactor := NewReviseDecisionInteractor(mockSvc, mockOut)
	err := interactor.ReviseDecision("model", "999", "")

	assert.ErrorContains(t, err, "failed to find original decision")
	mockSvc.AssertExpectations(t)
}

func TestReviseDecision_FailsToRevise(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionRevise)

	original := &decision.Decision{ID: "001", Title: "Old"}
	mockSvc.On("GetDecisionByID", "model", "001").Return(original, nil)
	mockSvc.On("Revise", "model", original).Return(nil, errors.New("revision failed"))

	interactor := NewReviseDecisionInteractor(mockSvc, mockOut)
	err := interactor.ReviseDecision("model", "001", "")

	assert.ErrorContains(t, err, "failed to revise decision")
	mockSvc.AssertExpectations(t)
}

func TestReviseDecision_FailsToLink(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionRevise)

	original := &decision.Decision{ID: "001"}
	revised := &decision.Decision{ID: "002"}

	mockSvc.On("GetDecisionByID", "model", "001").Return(original, nil)
	mockSvc.On("Revise", "model", original).Return(revised, nil)
	mockSvc.On("Link", "model", original, revised, "revised by", "revises").Return(errors.New("link error"))

	interactor := NewReviseDecisionInteractor(mockSvc, mockOut)
	err := interactor.ReviseDecision("model", "001", "")

	assert.ErrorContains(t, err, "failed to link revised decision")
	mockSvc.AssertExpectations(t)
}
