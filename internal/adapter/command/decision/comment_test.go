package decision

import (
	in_mocks "adg/mocks/inputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommentCommand_Success(t *testing.T) {
	mockInput := new(in_mocks.DecisionComment)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("GetAuthor").Return("auto-author")
	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")

	mockInput.On("Comment", "resolvedPath", "0001", "", "auto-author", "Great decision").Return(nil)

	cmd := NewCommentCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
		"--text", "Great decision",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestNewCommentCommand_MissingText(t *testing.T) {
	mockInput := new(in_mocks.DecisionComment)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")

	cmd := NewCommentCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{"--id", "0001"})

	err := cmd.Execute()
	assert.EqualError(t, err, "--text is required to provide the comment")
}

func TestNewCommentCommand_MissingAuthor(t *testing.T) {
	mockInput := new(in_mocks.DecisionComment)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("GetAuthor").Return("")
	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")

	cmd := NewCommentCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
		"--text", "Hello world",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "author must be provided using --author or set in config")
}

func TestNewCommentCommand_InputReturnsError(t *testing.T) {
	mockInput := new(in_mocks.DecisionComment)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("GetAuthor").Return("alice")
	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")

	mockInput.On("Comment", "resolvedPath", "0001", "", "alice", "bad comment").Return(errors.New("failure"))

	cmd := NewCommentCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
		"--text", "bad comment",
		"--author", "alice",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "failure")
}
