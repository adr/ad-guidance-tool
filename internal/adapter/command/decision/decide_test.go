package decision

import (
	in_mocks "adg/mocks/inputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDecideCommand_Success(t *testing.T) {
	mockInput := new(in_mocks.DecisionDecide)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("GetAuthor").Return("jane")
	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")

	mockInput.On("Decide", "resolvedPath", "0001", "", "2", "best option", "jane", false).Return(nil)

	cmd := NewDecideCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
		"--option", "2",
		"--rationale", "best option",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestNewDecideCommand_MissingOption(t *testing.T) {
	mockInput := new(in_mocks.DecisionDecide)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockConfig.On("GetAuthor").Return("author")

	cmd := NewDecideCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "--option must be provided (either its name or a positive integer (1-based index)")
}

func TestNewDecideCommand_InputFails(t *testing.T) {
	mockInput := new(in_mocks.DecisionDecide)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("GetAuthor").Return("kate")
	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")

	mockInput.On("Decide", "resolvedPath", "0001", "", "A", "", "kate", false).Return(errors.New("decision error"))

	cmd := NewDecideCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
		"--option", "A",
		"--author", "kate",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "decision error")
}
