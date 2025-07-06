package model

import (
	in_mocks "adg/mocks/inputport"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewMergeModelsCommand_Success(t *testing.T) {
	mockInput := new(in_mocks.ModelMerge)
	mockInput.
		On("Merge", "modelA", "modelB", "target", mock.Anything).
		Return(nil)

	cmd := NewMergeModelsCommand(mockInput)
	cmd.SetArgs([]string{
		"--model1", "modelA",
		"--model2", "modelB",
		"--target", "target",
		"--tag", "architecture",
		"--status", "open",
		"--title", "Login",
		"--id", "0001",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
	mockInput.AssertCalled(t, "Merge", "modelA", "modelB", "target", mock.MatchedBy(func(filters map[string][]string) bool {
		return len(filters) == 4 &&
			filters["tag"][0] == "architecture" &&
			filters["status"][0] == "open" &&
			filters["title"][0] == "Login" &&
			filters["id"][0] == "0001"
	}))
}

func TestNewMergeModelsCommand_MissingRequiredFlags(t *testing.T) {
	mockInput := new(in_mocks.ModelMerge)

	cmd := NewMergeModelsCommand(mockInput)
	cmd.SetArgs([]string{}) // Missing --model1, --model2, and --target

	err := cmd.Execute()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "--model1, --model2 and --target must all be provided")
}

func TestNewMergeModelsCommand_InputReturnsError(t *testing.T) {
	mockInput := new(in_mocks.ModelMerge)
	mockInput.
		On("Merge", "modelA", "modelB", "target", mock.Anything).
		Return(errors.New("merge failed"))

	cmd := NewMergeModelsCommand(mockInput)
	cmd.SetArgs([]string{
		"--model1", "modelA",
		"--model2", "modelB",
		"--target", "target",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "merge failed")
}
