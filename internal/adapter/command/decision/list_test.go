package decision

import (
	in_mocks "adg/mocks/inputport"
	svc_mocks "adg/mocks/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewListCommand_SuccessWithAllFilters(t *testing.T) {
	mockInput := new(in_mocks.DecisionList)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("ListDecisions", "resolvedPath", map[string][]string{
		"tag":    {"archived"},
		"status": {"open"},
		"title":  {"core"},
		"id":     {"0001"},
	}, "yaml").Return(nil)

	cmd := NewListCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--tag", "archived",
		"--status", "open",
		"--title", "core",
		"--id", "0001",
		"--format", "yaml",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestNewListCommand_MinimalValidInput(t *testing.T) {
	mockInput := new(in_mocks.DecisionList)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("ListDecisions", "resolvedPath", map[string][]string{}, "simple").Return(nil)

	cmd := NewListCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestNewListCommand_ConfigResolutionFails(t *testing.T) {
	mockInput := new(in_mocks.DecisionList)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("").Once()

	cmd := NewListCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	assert.Error(t, err)
}
