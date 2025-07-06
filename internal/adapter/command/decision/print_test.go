package decision

import (
	in_mocks "adg/mocks/inputport"
	svc_mocks "adg/mocks/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPrintCommand_AllSectionsDefaulted(t *testing.T) {
	mockInput := new(in_mocks.DecisionPrint)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("Print", "resolvedPath", []string{"0001"}, []string{"My Decision"}, map[string]bool{
		"question": true,
		"options":  true,
		"criteria": true,
		"comments": true,
		"outcome":  true,
	}).Return(nil)

	cmd := NewPrintCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001", "--id", "My Decision",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestNewPrintCommand_CustomSectionsSelected(t *testing.T) {
	mockInput := new(in_mocks.DecisionPrint)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("Print", "resolvedPath", []string{"0001"}, []string(nil), map[string]bool{
		"question": true,
		"options":  false,
		"criteria": false,
		"comments": false,
		"outcome":  false,
	}).Return(nil)

	cmd := NewPrintCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--id", "0001", "--question",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestNewPrintCommand_ErrorWhenNoIdsProvided(t *testing.T) {
	mockInput := new(in_mocks.DecisionPrint)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")

	cmd := NewPrintCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	assert.ErrorContains(t, err, "at least one --id or --title must be provided")
}

func TestNewPrintCommand_ModelPathResolutionFails(t *testing.T) {
	mockInput := new(in_mocks.DecisionPrint)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("").Once()

	cmd := NewPrintCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{"--id", "0001"})

	err := cmd.Execute()
	assert.Error(t, err)
}
