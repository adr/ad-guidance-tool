package commands

import (
	svc_mocks "adg/mocks/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveModelPathOrDefault_WithFlagValue(t *testing.T) {
	mockCfg := new(svc_mocks.ConfigService)
	path, err := ResolveModelPathOrDefault("/explicit/path", mockCfg)
	assert.NoError(t, err)
	assert.Equal(t, "/explicit/path", path)
}

func TestResolveModelPathOrDefault_FromConfig(t *testing.T) {
	mockCfg := new(svc_mocks.ConfigService)
	mockCfg.On("IsLoaded").Return(true)
	mockCfg.On("GetDefaultModelPath").Return("/default/from/config")

	path, err := ResolveModelPathOrDefault("", mockCfg)
	assert.NoError(t, err)
	assert.Equal(t, "/default/from/config", path)
}

func TestResolveModelPathOrDefault_MissingAll(t *testing.T) {
	mockCfg := new(svc_mocks.ConfigService)
	mockCfg.On("IsLoaded").Return(false)

	path, err := ResolveModelPathOrDefault("", mockCfg)
	assert.Error(t, err)
	assert.Empty(t, path)
	assert.Contains(t, err.Error(), "model path must be provided")
}

func TestResolveIdOrTitle_ValidID(t *testing.T) {
	var id, title string
	err := ResolveIdOrTitle("0001", &id, &title)
	assert.NoError(t, err)
	assert.Equal(t, "0001", id)
	assert.Empty(t, title)
}

func TestResolveIdOrTitle_ValidTitle(t *testing.T) {
	var id, title string
	err := ResolveIdOrTitle("my-decision", &id, &title)
	assert.NoError(t, err)
	assert.Equal(t, "my-decision", title)
	assert.Empty(t, id)
}

func TestResolveIdOrTitle_EmptyInput(t *testing.T) {
	var id, title string
	err := ResolveIdOrTitle("", &id, &title)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "you must specify the decisions via --id")
}

func TestResolveIdOrTitle_InvalidFormat(t *testing.T) {
	var id, title string
	err := ResolveIdOrTitle("####", &id, &title)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "input must be either a 4-digit ID")
}

func TestGetTemplateSections_Nygard(t *testing.T) {
	sections, err := GetTemplateSections("Nygard")
	assert.NoError(t, err)
	assert.Equal(t, "Context", sections["question"])
	assert.Equal(t, "Consequences", sections["criteria"])
	assert.Equal(t, "Decision", sections["outcome"])
}

func TestGetTemplateSections_MADR(t *testing.T) {
	sections, err := GetTemplateSections("madr")
	assert.NoError(t, err)
	assert.Equal(t, "Context and Problem Statement", sections["question"])
	assert.Equal(t, "Considered Options", sections["options"])
	assert.Equal(t, "Decision Drivers", sections["criteria"])
	assert.Equal(t, "Decision Outcome", sections["outcome"])
}

func TestGetTemplateSections_UnknownTemplate(t *testing.T) {
	sections, err := GetTemplateSections("random")
	assert.Error(t, err)
	assert.Nil(t, sections)
	assert.Contains(t, err.Error(), "unknown template")
}
