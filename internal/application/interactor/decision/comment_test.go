package decision

import (
	"adg/internal/domain/decision"
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComment_ByID_Success(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionComment)

	d := &decision.Decision{ID: "0012"}

	mockService.On("GetDecisionByID", "model", "0012").Return(d, nil)
	mockService.On("Comment", "model", d, "John", "Nice!").Return(nil)
	mockOutput.On("Commented", "0012", "John", "Nice!").Return(nil)

	interactor := NewCommentDecisionInteractor(mockService, mockOutput)
	err := interactor.Comment("model", "0012", "", "John", "Nice!")

	assert.NoError(t, err)
	mockService.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestComment_ByTitle_Success(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionComment)

	d := &decision.Decision{ID: "0042"}

	mockService.On("GetDecisionByTitle", "model", "My Decision").Return(d, nil)
	mockService.On("Comment", "model", d, "Alice", "I agree.").Return(nil)
	mockOutput.On("Commented", "0042", "Alice", "I agree.").Return(nil)

	interactor := NewCommentDecisionInteractor(mockService, mockOutput)
	err := interactor.Comment("model", "", "My Decision", "Alice", "I agree.")

	assert.NoError(t, err)
	mockService.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestComment_GetDecisionFails(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionComment)

	mockService.On("GetDecisionByID", "model", "9999").Return(nil, errors.New("not found"))

	interactor := NewCommentDecisionInteractor(mockService, mockOutput)
	err := interactor.Comment("model", "9999", "", "Bob", "Feedback")

	assert.ErrorContains(t, err, "not found")
	mockService.AssertExpectations(t)
}

func TestComment_CommentFails(t *testing.T) {
	mockService := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionComment)

	d := &decision.Decision{ID: "1001"}

	mockService.On("GetDecisionByID", "model", "1001").Return(d, nil)
	mockService.On("Comment", "model", d, "Jane", "Oops").Return(errors.New("repo error"))

	interactor := NewCommentDecisionInteractor(mockService, mockOutput)
	err := interactor.Comment("model", "1001", "", "Jane", "Oops")

	assert.ErrorContains(t, err, "failed to add comment")
	mockService.AssertExpectations(t)
}
