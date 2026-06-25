package decision

import (
	"errors"
	"testing"

	in_mocks "github.com/adr/ad-guidance-tool/mocks/inputport"
	svc_mocks "github.com/adr/ad-guidance-tool/mocks/service"

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
	assert.EqualError(t, err, "comment text must be provided (via arguments or --text flag)")
}

func TestNewCommentCommand_WithPositionalArgs(t *testing.T) {
	mockInput := new(in_mocks.DecisionComment)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("GetAuthor").Return("auto-author")
	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")

	mockInput.On("Comment", "resolvedPath", "0001", "", "auto-author", "This is my comment text").Return(nil)

	cmd := NewCommentCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
		"This", "is", "my", "comment", "text",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestNewCommentCommand_FlagsOverrideArgs(t *testing.T) {
	mockInput := new(in_mocks.DecisionComment)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("GetAuthor").Return("auto-author")
	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")

	mockInput.On("Comment", "resolvedPath", "0001", "", "auto-author", "Flag text").Return(nil)

	cmd := NewCommentCommand(mockInput, mockConfig)
	// When --text is provided, positional args should be ignored
	cmd.SetArgs([]string{
		"--id", "0001",
		"--text", "Flag text",
		"Ignored", "Args",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
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
