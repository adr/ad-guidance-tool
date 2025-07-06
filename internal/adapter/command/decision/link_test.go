package decision

import (
	in_mocks "adg/mocks/inputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkCommand_SuccessWithDefaultTags(t *testing.T) {
	mockInput := new(in_mocks.DecisionLink)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("Link", "resolvedPath", "0001", "", "0002", "", "precedes", "succeeds").Return(nil)

	cmd := NewLinkCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--from", "0001",
		"--to", "0002",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestNewLinkCommand_SuccessWithCustomTags(t *testing.T) {
	mockInput := new(in_mocks.DecisionLink)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("Link", "resolvedPath", "", "A", "", "B", "blocks", "blocked by").Return(nil)

	cmd := NewLinkCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--from", "A",
		"--to", "B",
		"--tag", "blocks",
		"--reverse-tag", "blocked by",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestNewLinkCommand_RejectsReservedTags(t *testing.T) {
	mockInput := new(in_mocks.DecisionLink)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")

	cmd := NewLinkCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--from", "0001",
		"--to", "0002",
		"--tag", "precedes",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, `you cannot use "precedes" or "succeeds" as custom tags; omit --tag and --reverse-tag to use them implicitly`)
}

func TestNewLinkCommand_InputReturnsError(t *testing.T) {
	mockInput := new(in_mocks.DecisionLink)
	mockConfig := new(svc_mocks.ConfigService)

	mockConfig.On("IsLoaded").Return(true)
	mockConfig.On("GetDefaultModelPath").Return("resolvedPath")
	mockInput.On("Link", "resolvedPath", "", "X", "", "Y", "precedes", "succeeds").Return(errors.New("link failed"))

	cmd := NewLinkCommand(mockInput, mockConfig)
	cmd.SetArgs([]string{
		"--from", "X",
		"--to", "Y",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "link failed")
}
