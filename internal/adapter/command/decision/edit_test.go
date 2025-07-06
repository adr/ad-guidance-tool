package decision

import (
	in_mocks "adg/mocks/inputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewEditCommand_Success(t *testing.T) {
	mockInput := new(in_mocks.DecisionEdit)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("Edit", "resolvedPath", "0001", "", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	cmd := NewEditCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
		"--question", "updated question",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestNewEditCommand_MissingFields(t *testing.T) {
	mockInput := new(in_mocks.DecisionEdit)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")

	cmd := NewEditCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "at least one of --question, --option, or --criteria must be provided")
}

func TestNewEditCommand_EditFails(t *testing.T) {
	mockInput := new(in_mocks.DecisionEdit)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("Edit", "resolvedPath", "0001", "", mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("edit failed"))

	cmd := NewEditCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
		"--question", "Q?",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "edit failed")
}
