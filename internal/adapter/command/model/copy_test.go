package model

import (
	in_mocks "adg/mocks/inputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewCopyCommand_MissingTarget(t *testing.T) {
	mockInput := new(in_mocks.ModelCopy)
	mockCfg := new(svc_mocks.ConfigService)

	cmd := NewCopyCommand(mockInput, mockCfg)
	cmd.SetArgs([]string{}) // no target

	err := cmd.Execute()

	assert.Error(t, err)
	assert.EqualError(t, err, "--target is required")
}

func TestNewCopyCommand_UsesDefaultModelFromConfig(t *testing.T) {
	mockInput := new(in_mocks.ModelCopy)
	mockCfg := new(svc_mocks.ConfigService)

	mockCfg.On("IsLoaded").Return(true)
	mockCfg.On("GetDefaultModelPath").Return("default/model")
	mockInput.On("Copy", "default/model", "target/path", map[string][]string{
		"tag":    {"foo"},
		"status": {"open"},
		"title":  {"Example"},
		"id":     {"0001"},
	}).Return(nil)

	cmd := NewCopyCommand(mockInput, mockCfg)
	cmd.SetArgs([]string{
		"--target", "target/path",
		"--tag", "foo",
		"--status", "open",
		"--title", "Example",
		"--id", "0001",
	})

	err := cmd.Execute()
	assert.NoError(t, err)

	mockInput.AssertExpectations(t)
	mockCfg.AssertExpectations(t)
}

func TestNewCopyCommand_InputReturnsError(t *testing.T) {
	mockInput := new(in_mocks.ModelCopy)
	mockCfg := new(svc_mocks.ConfigService)

	mockCfg.On("IsLoaded").Return(true)
	mockCfg.On("GetDefaultModelPath").Return("default/model")
	mockInput.On("Copy", "default/model", "target/path", mock.Anything).Return(errors.New("copy failed"))

	cmd := NewCopyCommand(mockInput, mockCfg)
	cmd.SetArgs([]string{"--target", "target/path"})

	err := cmd.Execute()

	assert.EqualError(t, err, "copy failed")
	mockInput.AssertExpectations(t)
	mockCfg.AssertExpectations(t)
}
