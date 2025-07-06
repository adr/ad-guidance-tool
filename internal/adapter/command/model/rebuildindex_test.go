package model

import (
	in_mocks "adg/mocks/inputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRebuildIndexCommand_Success(t *testing.T) {
	mockInput := new(in_mocks.ModelRebuildIndex)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("RebuildIndex", "resolvedPath").Return(nil)

	cmd := NewRebuildIndexCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{}) // No model flag, use config fallback

	err := cmd.Execute()
	assert.NoError(t, err)

	mockInput.AssertCalled(t, "RebuildIndex", "resolvedPath")
}

func TestNewRebuildIndexCommand_ConfigError(t *testing.T) {
	mockInput := new(in_mocks.ModelRebuildIndex)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("").Once()

	cmd := NewRebuildIndexCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	assert.Error(t, err)
}

func TestNewRebuildIndexCommand_InputReturnsError(t *testing.T) {
	mockInput := new(in_mocks.ModelRebuildIndex)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("RebuildIndex", "resolvedPath").Return(errors.New("rebuild failed"))

	cmd := NewRebuildIndexCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	assert.EqualError(t, err, "rebuild failed")
}
