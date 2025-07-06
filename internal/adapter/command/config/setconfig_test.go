package config

import (
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSetCommand_NoFlagsProvided(t *testing.T) {
	mockCfg := new(svc_mocks.ConfigService)
	cmd := NewSetCommand(mockCfg)
	cmd.SetArgs([]string{}) // No flags

	err := cmd.Execute()
	assert.ErrorContains(t, err, "at least one configuration flag must be provided")
}

func TestNewSetCommand_SetOnlyConfigPath(t *testing.T) {
	mockCfg := new(svc_mocks.ConfigService)
	mockCfg.On("SetConfigPath", "/tmp/myconfig.yaml").Return(nil)

	cmd := NewSetCommand(mockCfg)
	cmd.SetArgs([]string{"--config-path", "/tmp/myconfig.yaml"})

	err := cmd.Execute()
	assert.NoError(t, err)
	mockCfg.AssertCalled(t, "SetConfigPath", "/tmp/myconfig.yaml")
}

func TestNewSetCommand_SetTemplateAndSave(t *testing.T) {
	mockCfg := new(svc_mocks.ConfigService)

	expectedPath := "/mock/saved/config.yaml"
	mockCfg.On("Save", "Context", "Consequences", "", "", "Decision", "", "").Return(expectedPath, nil)

	cmd := NewSetCommand(mockCfg)
	cmd.SetArgs([]string{"--template", "nygard"})

	err := cmd.Execute()
	assert.NoError(t, err)
	mockCfg.AssertCalled(t, "Save", "Context", "Consequences", "", "", "Decision", "", "")
}

func TestNewSetCommand_UnknownTemplate(t *testing.T) {
	mockCfg := new(svc_mocks.ConfigService)

	cmd := NewSetCommand(mockCfg)
	cmd.SetArgs([]string{"--template", "unknown"})

	err := cmd.Execute()
	assert.Error(t, err)
}

func TestNewSetCommand_SaveFails(t *testing.T) {
	mockCfg := new(svc_mocks.ConfigService)
	mockCfg.On("Save", "q", "c", "o", "cmt", "out", "me", "path").Return("", errors.New("save failed"))

	cmd := NewSetCommand(mockCfg)
	cmd.SetArgs([]string{
		"--question", "q",
		"--criteria", "c",
		"--options", "o",
		"--comments", "cmt",
		"--outcome", "out",
		"--author", "me",
		"--model", "path",
	})

	err := cmd.Execute()
	assert.ErrorContains(t, err, "save failed")
	mockCfg.AssertCalled(t, "Save", "q", "c", "o", "cmt", "out", "me", "path")
}

func TestNewSetCommand_SetConfigPathFails(t *testing.T) {
	mockCfg := new(svc_mocks.ConfigService)
	mockCfg.On("SetConfigPath", "broken-path").Return(errors.New("path invalid"))

	cmd := NewSetCommand(mockCfg)
	cmd.SetArgs([]string{"--config-path", "broken-path"})

	err := cmd.Execute()
	assert.ErrorContains(t, err, "path invalid")
	mockCfg.AssertCalled(t, "SetConfigPath", "broken-path")
}
