package model

import (
	in_mocks "adg/mocks/inputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidateCommand_Success(t *testing.T) {
	mockInput := new(in_mocks.ModelValidate)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedModelPath")
	mockInput.On("Validate", "resolvedModelPath").Return(nil)

	cmd := NewValidateCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{}) // simulate no --model flag

	err := cmd.Execute()
	assert.NoError(t, err)

	mockInput.AssertCalled(t, "Validate", "resolvedModelPath")
}

func TestNewValidateCommand_ConfigError(t *testing.T) {
	mockInput := new(in_mocks.ModelValidate)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("").Once()

	cmd := NewValidateCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	assert.Error(t, err)
}

func TestNewValidateCommand_InputReturnsError(t *testing.T) {
	mockInput := new(in_mocks.ModelValidate)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedModelPath")
	mockInput.On("Validate", "resolvedModelPath").Return(errors.New("validation failed"))

	cmd := NewValidateCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	assert.EqualError(t, err, "validation failed")
}
