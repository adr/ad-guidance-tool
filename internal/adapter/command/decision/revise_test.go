package decision

import (
	in_mocks "adg/mocks/inputport"
	svc_mocks "adg/mocks/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReviseCommand_ValidExecution(t *testing.T) {
	mockInput := new(in_mocks.DecisionRevise)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("ReviseDecision", "resolvedPath", "0001", "").Return(nil)

	cmd := NewReviseCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestNewReviseCommand_ModelPathResolutionFails(t *testing.T) {
	mockInput := new(in_mocks.DecisionRevise)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("")

	cmd := NewReviseCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
	})

	err := cmd.Execute()
	assert.Error(t, err)
}
