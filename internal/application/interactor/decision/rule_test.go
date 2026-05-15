package decision

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	decisiondomain "github.com/adr/ad-guidance-tool/internal/domain/decision"
	out_mocks "github.com/adr/ad-guidance-tool/mocks/outputport"
	svc_mocks "github.com/adr/ad-guidance-tool/mocks/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRule_SuccessWithID(t *testing.T) {
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionRule)

	// Create a temp directory for testing
	tempDir := t.TempDir()
	modelPath := filepath.Join(tempDir, "model")
	os.MkdirAll(modelPath, 0755)

	id := "0001"
	adrFilePath := filepath.Join(modelPath, "AD0001-test-decision.md")
	decision := &decisiondomain.Decision{
		ID:    "0001",
		Title: "test-decision",
	}

	mockDecisionSvc.On("GetDecisionByID", modelPath, id).Return(decision, nil)
	mockDecisionSvc.On("GetDecisionFilePath", modelPath, id).Return(adrFilePath, nil)
	mockOutput.On("RuleGenerated", id, mock.AnythingOfType("string")).Return()

	interactor := NewRuleInteractor(mockDecisionSvc, mockOutput)

	err := interactor.Rule(modelPath, id, "", "")

	assert.NoError(t, err)
	mockDecisionSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestRule_SuccessWithTitle(t *testing.T) {
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionRule)

	// Create a temp directory for testing
	tempDir := t.TempDir()
	modelPath := filepath.Join(tempDir, "model")
	os.MkdirAll(modelPath, 0755)

	title := "test-decision"
	adrFilePath := filepath.Join(modelPath, "AD0001-test-decision.md")
	decision := &decisiondomain.Decision{
		ID:    "0001",
		Title: "test-decision",
	}

	mockDecisionSvc.On("GetDecisionByTitle", modelPath, title).Return(decision, nil)
	mockDecisionSvc.On("GetDecisionFilePath", modelPath, decision.ID).Return(adrFilePath, nil)
	mockOutput.On("RuleGenerated", decision.ID, mock.AnythingOfType("string")).Return()

	interactor := NewRuleInteractor(mockDecisionSvc, mockOutput)

	err := interactor.Rule(modelPath, "", title, "")

	assert.NoError(t, err)
	mockDecisionSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestRule_WithCustomOutputPath(t *testing.T) {
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionRule)

	modelPath := "model"
	id := "0001"
	customOutput := filepath.Join(t.TempDir(), "custom.rule")
	decision := &decisiondomain.Decision{
		ID:    "0001",
		Title: "test-decision",
	}

	mockDecisionSvc.On("GetDecisionByID", modelPath, id).Return(decision, nil)
	mockOutput.On("RuleGenerated", id, customOutput).Return()

	interactor := NewRuleInteractor(mockDecisionSvc, mockOutput)

	err := interactor.Rule(modelPath, id, "", customOutput)

	assert.NoError(t, err)
	mockDecisionSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)

	// Verify file was created
	assert.FileExists(t, customOutput)

	// Verify content
	content, err := os.ReadFile(customOutput)
	assert.NoError(t, err)
	assert.Contains(t, string(content), `adr "0001" "test-decision"`)
	assert.Contains(t, string(content), `code "rule_name"`)
	assert.Contains(t, string(content), `file "rule_name"`)
}

func TestRule_WithCustomOutputDirectory(t *testing.T) {
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionRule)

	modelPath := "model"
	id := "0001"
	tempDir := t.TempDir()
	adrFilePath := filepath.Join(modelPath, "AD0001-test-decision.md")
	expectedOutput := filepath.Join(tempDir, "AD0001-test-decision.rule")
	decision := &decisiondomain.Decision{
		ID:    "0001",
		Title: "test-decision",
	}

	mockDecisionSvc.On("GetDecisionByID", modelPath, id).Return(decision, nil)
	mockDecisionSvc.On("GetDecisionFilePath", modelPath, id).Return(adrFilePath, nil)
	mockOutput.On("RuleGenerated", id, expectedOutput).Return()

	interactor := NewRuleInteractor(mockDecisionSvc, mockOutput)

	err := interactor.Rule(modelPath, id, "", tempDir)

	assert.NoError(t, err)
	mockDecisionSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)

	// Verify file was created with correct name
	assert.FileExists(t, expectedOutput)
}

func TestRule_NoIDOrTitle(t *testing.T) {
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionRule)

	interactor := NewRuleInteractor(mockDecisionSvc, mockOutput)

	err := interactor.Rule("model", "", "", "")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "either --id or --title must be provided")
}

func TestRule_DecisionNotFoundByID(t *testing.T) {
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionRule)

	modelPath := "model"
	id := "9999"

	mockDecisionSvc.On("GetDecisionByID", modelPath, id).Return(nil, errors.New("not found"))

	interactor := NewRuleInteractor(mockDecisionSvc, mockOutput)

	err := interactor.Rule(modelPath, id, "", "")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get decision by ID")
	mockDecisionSvc.AssertExpectations(t)
}

func TestRule_DecisionNotFoundByTitle(t *testing.T) {
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionRule)

	modelPath := "model"
	title := "nonexistent"

	mockDecisionSvc.On("GetDecisionByTitle", modelPath, title).Return(nil, errors.New("not found"))

	interactor := NewRuleInteractor(mockDecisionSvc, mockOutput)

	err := interactor.Rule(modelPath, "", title, "")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get decision by title")
	mockDecisionSvc.AssertExpectations(t)
}

func TestRule_FilePathRetrievalFails(t *testing.T) {
	mockDecisionSvc := new(svc_mocks.DecisionService)
	mockOutput := new(out_mocks.DecisionRule)

	modelPath := "model"
	id := "0001"
	decision := &decisiondomain.Decision{
		ID:    "0001",
		Title: "test-decision",
	}

	mockDecisionSvc.On("GetDecisionByID", modelPath, id).Return(decision, nil)
	mockDecisionSvc.On("GetDecisionFilePath", modelPath, id).Return("", errors.New("file not found"))

	interactor := NewRuleInteractor(mockDecisionSvc, mockOutput)

	err := interactor.Rule(modelPath, id, "", "")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get decision file path")
	mockDecisionSvc.AssertExpectations(t)
}

func TestRule_TemplateGeneration(t *testing.T) {
	// Test the template function directly
	id := "0001"
	title := "test-decision"

	content := generateRuleTemplate(id, title)

	assert.Contains(t, content, `adr "0001" "test-decision"`)
	assert.Contains(t, content, `code "rule_name"`)
	assert.Contains(t, content, `file "rule_name"`)
	assert.Contains(t, content, "severity error")
}
