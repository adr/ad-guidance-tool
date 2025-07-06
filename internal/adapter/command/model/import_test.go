package model

import (
	in_mocks "adg/mocks/inputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewImportCommand_MissingSource(t *testing.T) {
	mockInput := new(in_mocks.ModelImport)
	mockCfg := new(svc_mocks.ConfigService)

	cmd := NewImportCommand(mockInput, mockCfg)
	cmd.SetArgs([]string{}) // No --source

	err := cmd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "--target must both be provided")
}

func TestNewImportCommand_SuccessfulImport(t *testing.T) {
	mockInput := new(in_mocks.ModelImport)
	mockCfg := new(svc_mocks.ConfigService)

	mockCfg.On("IsLoaded").Return(true)
	mockCfg.On("GetDefaultModelPath").Return("target/model")
	mockInput.On("Import", "source/model", "target/model", map[string][]string{
		"tag":    {"critical"},
		"status": {"open"},
		"title":  {"Pattern"},
		"id":     {"0005"},
	}).Return(nil)

	cmd := NewImportCommand(mockInput, mockCfg)
	cmd.SetArgs([]string{
		"--source", "source/model",
		"--tag", "critical",
		"--status", "open",
		"--title", "Pattern",
		"--id", "0005",
	})

	err := cmd.Execute()
	assert.NoError(t, err)

	mockInput.AssertExpectations(t)
	mockCfg.AssertExpectations(t)
}

func TestNewImportCommand_InputReturnsError(t *testing.T) {
	mockInput := new(in_mocks.ModelImport)
	mockCfg := new(svc_mocks.ConfigService)

	mockCfg.On("IsLoaded").Return(true)
	mockCfg.On("GetDefaultModelPath").Return("target/model")
	mockInput.On("Import", "source/model", "target/model", mock.Anything).Return(errors.New("import failed"))

	cmd := NewImportCommand(mockInput, mockCfg)
	cmd.SetArgs([]string{"--source", "source/model"})

	err := cmd.Execute()
	assert.EqualError(t, err, "import failed")

	mockInput.AssertExpectations(t)
	mockCfg.AssertExpectations(t)
}
