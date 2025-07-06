package decision

import (
	in_mocks "adg/mocks/inputport"
	svc_mocks "adg/mocks/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTagCommand_ValidExecution(t *testing.T) {
	mockInput := new(in_mocks.DecisionTag)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("Tag", "resolvedPath", "0001", "", []string{"architecture", "urgent"}).Return(nil)

	cmd := NewTagCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001",
		"--tag", "architecture",
		"--tag", "urgent",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestNewTagCommand_ErrorWhenNoTagsProvided(t *testing.T) {
	mockInput := new(in_mocks.DecisionTag)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")

	cmd := NewTagCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{"--id", "0001"})

	err := cmd.Execute()
	assert.ErrorContains(t, err, "at least one tag must be specified using --tag")
}

func TestNewTagCommand_ModelPathResolutionFails(t *testing.T) {
	mockInput := new(in_mocks.DecisionTag)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("")

	cmd := NewTagCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{"--id", "0001", "--tag", "core"})

	err := cmd.Execute()
	assert.Error(t, err)
}
