package decision

import (
	in_mocks "github.com/adr/ad-guidance-tool/mocks/inputport"
	svc_mocks "github.com/adr/ad-guidance-tool/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRuleCommand_SuccessWithID(t *testing.T) {
	mockInput := new(in_mocks.DecisionRule)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("Rule", "resolvedPath", "0001", "", "").Return(nil)

	cmd := NewRuleCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{"--id", "0001"})

	err := cmd.Execute()
	assert.NoError(t, err)

	mockInput.AssertCalled(t, "Rule", "resolvedPath", "0001", "", "")
}

func TestNewRuleCommand_SuccessWithTitle(t *testing.T) {
	mockInput := new(in_mocks.DecisionRule)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("Rule", "resolvedPath", "", "test-decision", "").Return(nil)

	cmd := NewRuleCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{"--title", "test-decision"})

	err := cmd.Execute()
	assert.NoError(t, err)

	mockInput.AssertCalled(t, "Rule", "resolvedPath", "", "test-decision", "")
}

func TestNewRuleCommand_SuccessWithCustomOutput(t *testing.T) {
	mockInput := new(in_mocks.DecisionRule)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("Rule", "resolvedPath", "0001", "", "custom.rule").Return(nil)

	cmd := NewRuleCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{"--id", "0001", "--output", "custom.rule"})

	err := cmd.Execute()
	assert.NoError(t, err)

	mockInput.AssertCalled(t, "Rule", "resolvedPath", "0001", "", "custom.rule")
}

func TestNewRuleCommand_InputReturnsError(t *testing.T) {
	mockInput := new(in_mocks.DecisionRule)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("Rule", "resolvedPath", "0001", "", "").Return(errors.New("rule failed"))

	cmd := NewRuleCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{"--id", "0001"})

	err := cmd.Execute()
	assert.EqualError(t, err, "rule failed")
}

func TestNewRuleCommand_ConfigFails(t *testing.T) {
	mockInput := new(in_mocks.DecisionRule)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(false)
	mockConfig.On("GetDefaultModelPath").Return("")

	cmd := NewRuleCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{"--id", "0001"})

	err := cmd.Execute()
	assert.Error(t, err)
}
