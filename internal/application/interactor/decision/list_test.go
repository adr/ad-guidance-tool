package decision

import (
	"adg/internal/domain/decision"
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListDecisions_SuccessWithoutFilters(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionList)

	decisions := []decision.Decision{
		{ID: "001", Title: "First"},
		{ID: "002", Title: "Second"},
	}

	mockSvc.On("GetAllDecisions", "model").Return(decisions, nil)
	mockOut.On("Listed", decisions, "table").Return(nil)

	interactor := NewListDecisionsInteractor(mockSvc, mockOut)
	err := interactor.ListDecisions("model", map[string][]string{}, "table")

	assert.NoError(t, err)
	mockSvc.AssertExpectations(t)
	mockOut.AssertExpectations(t)
}

func TestListDecisions_SuccessWithFilters(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionList)

	allDecisions := []decision.Decision{
		{ID: "001", Title: "Alpha"},
		{ID: "002", Title: "Beta"},
	}
	filtered := []decision.Decision{
		{ID: "002", Title: "Beta"},
	}

	mockSvc.On("GetAllDecisions", "model").Return(allDecisions, nil)
	mockSvc.On("FilterDecisions", allDecisions, mock.Anything).Return(filtered, nil)
	mockOut.On("Listed", filtered, "json").Return(nil)

	interactor := NewListDecisionsInteractor(mockSvc, mockOut)
	err := interactor.ListDecisions("model", map[string][]string{"id": {"002"}}, "json")

	assert.NoError(t, err)
	mockSvc.AssertExpectations(t)
	mockOut.AssertExpectations(t)
}

func TestListDecisions_GetAllFails(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionList)

	mockSvc.On("GetAllDecisions", "model").Return(nil, errors.New("failed to load"))

	interactor := NewListDecisionsInteractor(mockSvc, mockOut)
	err := interactor.ListDecisions("model", nil, "any")

	assert.ErrorContains(t, err, "failed to load")
	mockSvc.AssertExpectations(t)
}

func TestListDecisions_FilterFails(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionList)

	raw := []decision.Decision{
		{ID: "001", Title: "X"},
	}

	mockSvc.On("GetAllDecisions", "model").Return(raw, nil)
	mockSvc.On("FilterDecisions", raw, mock.Anything).Return(nil, errors.New("bad filter"))

	interactor := NewListDecisionsInteractor(mockSvc, mockOut)
	err := interactor.ListDecisions("model", map[string][]string{"tag": {"urgent"}}, "yaml")

	assert.ErrorContains(t, err, "bad filter")
}
