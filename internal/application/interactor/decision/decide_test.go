package decision

import (
	"adg/internal/domain/decision"
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecide_ByID_Success(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionDecide)

	d := &decision.Decision{ID: "0005", Status: "open"}

	mockService.On("GetDecisionByID", "model", "0005").Return(d, nil)
	mockService.On("Decide", "model", d, "Option A", "Clear reason", true).Return(nil)
	mockService.On("Comment", "model", d, "Alice", "marked decision as decided").Return(nil)
	mockOutput.On("Decided", "0005").Return(nil)

	interactor := NewDecideInteractor(mockService, mockOutput)
	err := interactor.Decide("model", "0005", "", "Option A", "Clear reason", "Alice", true)

	assert.NoError(t, err)
	mockService.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestDecide_ByTitle_Success(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionDecide)

	d := &decision.Decision{ID: "0020", Status: "open"}

	mockService.On("GetDecisionByTitle", "model", "Important").Return(d, nil)
	mockService.On("Decide", "model", d, "1", "", false).Return(nil)
	mockService.On("Comment", "model", d, "Bob", "marked decision as decided").Return(nil)
	mockOutput.On("Decided", "0020").Return(nil)

	interactor := NewDecideInteractor(mockService, mockOutput)
	err := interactor.Decide("model", "", "Important", "1", "", "Bob", false)

	assert.NoError(t, err)
	mockService.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestDecide_AlreadyDecided(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionDecide)

	d := &decision.Decision{ID: "0042", Status: "decided"}

	mockService.On("GetDecisionByID", "model", "0042").Return(d, nil)

	interactor := NewDecideInteractor(mockService, mockOutput)
	err := interactor.Decide("model", "0042", "", "Any", "", "Someone", true)

	assert.ErrorContains(t, err, "already been decided")
	mockService.AssertExpectations(t)
}

func TestDecide_GetDecisionFails(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionDecide)

	mockService.On("GetDecisionByID", "model", "1234").Return(nil, errors.New("not found"))

	interactor := NewDecideInteractor(mockService, mockOutput)
	err := interactor.Decide("model", "1234", "", "X", "", "Y", true)

	assert.ErrorContains(t, err, "not found")
	mockService.AssertExpectations(t)
}

func TestDecide_DecideFails(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionDecide)

	d := &decision.Decision{ID: "0100", Status: "open"}

	mockService.On("GetDecisionByID", "model", "0100").Return(d, nil)
	mockService.On("Decide", "model", d, "X", "", false).Return(errors.New("fail"))

	interactor := NewDecideInteractor(mockService, mockOutput)
	err := interactor.Decide("model", "0100", "", "X", "", "author", false)

	assert.ErrorContains(t, err, "fail")
	mockService.AssertExpectations(t)
}

func TestDecide_CommentFails(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionDecide)

	d := &decision.Decision{ID: "0777", Status: "open"}

	mockService.On("GetDecisionByID", "model", "0777").Return(d, nil)
	mockService.On("Decide", "model", d, "Y", "", false).Return(nil)
	mockService.On("Comment", "model", d, "Zed", "marked decision as decided").Return(errors.New("write failed"))

	interactor := NewDecideInteractor(mockService, mockOutput)
	err := interactor.Decide("model", "0777", "", "Y", "", "Zed", false)

	assert.ErrorContains(t, err, "write failed")
	mockService.AssertExpectations(t)
}
