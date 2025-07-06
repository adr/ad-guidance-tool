package decision

import (
	"adg/internal/domain/decision"
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrint_SuccessWithIDsAndTitles(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionPrint)

	content1 := &decision.DecisionContent{ID: "001", Question: "Q1"}
	content2 := &decision.DecisionContent{ID: "002", Question: "Q2"}
	sections := map[string]bool{"question": true}

	mockSvc.On("GetDecisionContent", "model", "001").Return(content1, nil)
	mockSvc.On("GetDecisionByTitle", "model", "Decision 2").
		Return(&decision.Decision{ID: "002"}, nil)
	mockSvc.On("GetDecisionContent", "model", "002").Return(content2, nil)
	mockOut.On("Printed", []decision.DecisionContent{*content1, *content2}, sections).Return(nil)

	interactor := NewPrintDecisionsInteractor(mockSvc, mockOut)
	err := interactor.Print("model", []string{"001"}, []string{"Decision 2"}, sections)

	assert.NoError(t, err)
	mockSvc.AssertExpectations(t)
	mockOut.AssertExpectations(t)
}

func TestPrint_FailsOnInvalidID(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionPrint)

	mockSvc.On("GetDecisionContent", "model", "001").Return(nil, errors.New("not found"))

	interactor := NewPrintDecisionsInteractor(mockSvc, mockOut)
	err := interactor.Print("model", []string{"001"}, nil, nil)

	assert.ErrorContains(t, err, "failed to load content for ID")
	mockSvc.AssertExpectations(t)
}

func TestPrint_FailsOnInvalidTitle(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionPrint)

	mockSvc.On("GetDecisionByTitle", "model", "Missing").Return(nil, errors.New("not found"))

	interactor := NewPrintDecisionsInteractor(mockSvc, mockOut)
	err := interactor.Print("model", nil, []string{"Missing"}, nil)

	assert.ErrorContains(t, err, "failed to resolve title")
	mockSvc.AssertExpectations(t)
}

func TestPrint_FailsOnContentLoadFromResolvedTitle(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionPrint)

	mockSvc.On("GetDecisionByTitle", "model", "D1").
		Return(&decision.Decision{ID: "009"}, nil)
	mockSvc.On("GetDecisionContent", "model", "009").
		Return(nil, errors.New("bad content"))

	interactor := NewPrintDecisionsInteractor(mockSvc, mockOut)
	err := interactor.Print("model", nil, []string{"D1"}, nil)

	assert.ErrorContains(t, err, "failed to load content for title")
	mockSvc.AssertExpectations(t)
}
