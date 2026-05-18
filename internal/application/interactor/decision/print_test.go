package decision

import (
	"adg/internal/application/outputport"
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

	body1 := "# A\n\n## Context and Problem Statement\n\nQ1\n"
	body2 := "# B\n\n## Context and Problem Statement\n\nQ2\n"
	sections := map[string]bool{"context": true}

	mockSvc.On("GetBody", "model", "001").Return(body1, nil)
	mockSvc.On("GetDecisionByTitle", "model", "Decision 2").
		Return(&decision.Decision{ID: "002"}, nil)
	mockSvc.On("GetBody", "model", "002").Return(body2, nil)
	mockOut.On("Printed", []outputport.DecisionBody{
		{ID: "001", Body: body1},
		{ID: "002", Body: body2},
	}, sections).Return(nil)

	interactor := NewPrintDecisionsInteractor(mockSvc, mockOut)
	err := interactor.Print("model", []string{"001"}, []string{"Decision 2"}, sections)

	assert.NoError(t, err)
	mockSvc.AssertExpectations(t)
	mockOut.AssertExpectations(t)
}

func TestPrint_FailsOnInvalidID(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionPrint)

	mockSvc.On("GetBody", "model", "001").Return("", errors.New("not found"))

	interactor := NewPrintDecisionsInteractor(mockSvc, mockOut)
	err := interactor.Print("model", []string{"001"}, nil, nil)

	assert.ErrorContains(t, err, "failed to load body for ID")
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

func TestPrint_FailsOnBodyLoadFromResolvedTitle(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionPrint)

	mockSvc.On("GetDecisionByTitle", "model", "D1").
		Return(&decision.Decision{ID: "009"}, nil)
	mockSvc.On("GetBody", "model", "009").
		Return("", errors.New("bad body"))

	interactor := NewPrintDecisionsInteractor(mockSvc, mockOut)
	err := interactor.Print("model", nil, []string{"D1"}, nil)

	assert.ErrorContains(t, err, "failed to load body for title")
	mockSvc.AssertExpectations(t)
}
