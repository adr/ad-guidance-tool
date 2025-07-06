package decision

import (
	"adg/internal/domain/decision"
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLink_Success(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionLink)

	src := &decision.Decision{ID: "001"}
	dst := &decision.Decision{ID: "002"}

	mockSvc.On("GetDecisionByID", "model", "001").Return(src, nil)
	mockSvc.On("GetDecisionByID", "model", "002").Return(dst, nil)
	mockSvc.On("Link", "model", src, dst, "precedes", "succeeds").Return(nil)
	mockOut.On("Linked", "001", "002", "precedes", "succeeds").Return(nil)

	i := NewLinkDecisionsInteractor(mockSvc, mockOut)
	err := i.Link("model", "001", "", "002", "", "precedes", "succeeds")

	assert.NoError(t, err)
	mockSvc.AssertExpectations(t)
	mockOut.AssertExpectations(t)
}

func TestLink_ResolveSourceFails(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionLink)

	mockSvc.On("GetDecisionByID", "model", "001").Return(nil, errors.New("not found"))

	i := NewLinkDecisionsInteractor(mockSvc, mockOut)
	err := i.Link("model", "001", "", "002", "", "precedes", "succeeds")

	assert.ErrorContains(t, err, "could not find source decision")
	mockSvc.AssertExpectations(t)
}

func TestLink_ResolveTargetFails(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionLink)

	src := &decision.Decision{ID: "001"}
	mockSvc.On("GetDecisionByID", "model", "001").Return(src, nil)
	mockSvc.On("GetDecisionByID", "model", "002").Return(nil, errors.New("not found"))

	i := NewLinkDecisionsInteractor(mockSvc, mockOut)
	err := i.Link("model", "001", "", "002", "", "precedes", "succeeds")

	assert.ErrorContains(t, err, "could not find target decision")
}

func TestLink_SelfLinkError(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionLink)

	shared := &decision.Decision{ID: "003"}
	mockSvc.On("GetDecisionByID", "model", "003").Return(shared, nil).Twice()

	i := NewLinkDecisionsInteractor(mockSvc, mockOut)
	err := i.Link("model", "003", "", "003", "", "relates", "relates")

	assert.ErrorContains(t, err, "cannot create a tag from a decision to itself")
}

func TestLink_ServiceFails(t *testing.T) {
	mockSvc := new(svc_mocks.DecisionService)
	mockOut := new(out_mocks.DecisionLink)

	src := &decision.Decision{ID: "004"}
	dst := &decision.Decision{ID: "005"}

	mockSvc.On("GetDecisionByID", "model", "004").Return(src, nil)
	mockSvc.On("GetDecisionByID", "model", "005").Return(dst, nil)
	mockSvc.On("Link", "model", src, dst, "relates", "").Return(errors.New("write error"))

	i := NewLinkDecisionsInteractor(mockSvc, mockOut)
	err := i.Link("model", "004", "", "005", "", "relates", "")

	assert.ErrorContains(t, err, "linking failed")
}
