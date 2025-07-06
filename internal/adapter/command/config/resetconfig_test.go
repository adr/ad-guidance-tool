package config

import (
	svc_mocks "adg/mocks/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResetCommand_FullReset(t *testing.T) {
	mockConfig := new(svc_mocks.ConfigService)
	mockConfig.On("ResetAll").Return(nil)

	cmd := NewResetCommand(mockConfig)
	cmd.SetArgs([]string{}) // no flags

	err := cmd.Execute()
	assert.NoError(t, err)
	mockConfig.AssertCalled(t, "ResetAll")
}

func TestNewResetCommand_TemplateReset(t *testing.T) {
	mockConfig := new(svc_mocks.ConfigService)
	mockConfig.On("ResetTemplateHeaders").Return(nil)

	cmd := NewResetCommand(mockConfig)
	cmd.SetArgs([]string{"--template"})

	err := cmd.Execute()
	assert.NoError(t, err)
	mockConfig.AssertCalled(t, "ResetTemplateHeaders")
}

func TestNewResetCommand_FullResetFails(t *testing.T) {
	mockConfig := new(svc_mocks.ConfigService)
	mockConfig.On("ResetAll").Return(assert.AnError)

	cmd := NewResetCommand(mockConfig)
	cmd.SetArgs([]string{}) // no flags

	err := cmd.Execute()
	assert.Error(t, err)
	mockConfig.AssertCalled(t, "ResetAll")
}

func TestNewResetCommand_TemplateResetFails(t *testing.T) {
	mockConfig := new(svc_mocks.ConfigService)
	mockConfig.On("ResetTemplateHeaders").Return(assert.AnError)

	cmd := NewResetCommand(mockConfig)
	cmd.SetArgs([]string{"--template"})

	err := cmd.Execute()
	assert.Error(t, err)
	mockConfig.AssertCalled(t, "ResetTemplateHeaders")
}
