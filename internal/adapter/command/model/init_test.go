package model

import (
	in_mocks "adg/mocks/inputport"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInitCommand_Success(t *testing.T) {
	mockInput := new(in_mocks.ModelInit)
	mockInput.On("Init", "new-model").Return(nil)

	cmd := NewInitCommand(mockInput)
	cmd.SetArgs([]string{"new-model"})

	err := cmd.Execute()
	assert.NoError(t, err)
	mockInput.AssertCalled(t, "Init", "new-model")
}

func TestNewInitCommand_InputReturnsError(t *testing.T) {
	mockInput := new(in_mocks.ModelInit)
	mockInput.On("Init", "new-model").Return(errors.New("init failed"))

	cmd := NewInitCommand(mockInput)
	cmd.SetArgs([]string{"new-model"})

	err := cmd.Execute()
	assert.EqualError(t, err, "init failed")
	mockInput.AssertCalled(t, "Init", "new-model")
}

func TestNewInitCommand_MissingArgument(t *testing.T) {
	mockInput := new(in_mocks.ModelInit)

	cmd := NewInitCommand(mockInput)
	cmd.SetArgs([]string{}) // No arguments

	err := cmd.Execute()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "accepts 1 arg(s)")
}
