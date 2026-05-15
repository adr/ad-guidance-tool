package decision

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adr/ad-guidance-tool/internal/application/inputport"
	"github.com/adr/ad-guidance-tool/internal/application/outputport"
	decisiondomain "github.com/adr/ad-guidance-tool/internal/domain/decision"
)

type RuleInteractor struct {
	decisionService decisiondomain.DecisionService
	output          outputport.DecisionRule
}

func NewRuleInteractor(
	decisionService decisiondomain.DecisionService,
	output outputport.DecisionRule,
) inputport.DecisionRule {
	return &RuleInteractor{
		decisionService: decisionService,
		output:          output,
	}
}

func (i *RuleInteractor) Rule(modelPath, id, title, outputPath string) error {
	// Resolve the decision by ID or title
	var decision *decisiondomain.Decision
	var err error

	if id != "" {
		decision, err = i.decisionService.GetDecisionByID(modelPath, id)
		if err != nil {
			return fmt.Errorf("failed to get decision by ID %q: %w", id, err)
		}
	} else if title != "" {
		decision, err = i.decisionService.GetDecisionByTitle(modelPath, title)
		if err != nil {
			return fmt.Errorf("failed to get decision by title %q: %w", title, err)
		}
	} else {
		return fmt.Errorf("either --id or --title must be provided")
	}

	// Generate the rule file content
	ruleContent := generateRuleTemplate(decision.ID, decision.Title)

	// Determine output path
	var ruleFilePath string
	if outputPath != "" {
		// Custom output path provided
		info, err := os.Stat(outputPath)
		if err == nil && info.IsDir() {
			// If it's a directory, we still need to get the ADR filename
			adrFilePath, err := i.decisionService.GetDecisionFilePath(modelPath, decision.ID)
			if err != nil {
				return fmt.Errorf("failed to get decision file path: %w", err)
			}
			adrBasename := filepath.Base(adrFilePath)
			ruleFilename := strings.TrimSuffix(adrBasename, filepath.Ext(adrBasename)) + ".rule"
			ruleFilePath = filepath.Join(outputPath, ruleFilename)
		} else {
			ruleFilePath = outputPath
		}
	} else {
		// Default: same directory as the ADR, same filename with .rule extension
		adrFilePath, err := i.decisionService.GetDecisionFilePath(modelPath, decision.ID)
		if err != nil {
			return fmt.Errorf("failed to get decision file path: %w", err)
		}
		adrDir := filepath.Dir(adrFilePath)
		adrBasename := filepath.Base(adrFilePath)
		ruleFilename := strings.TrimSuffix(adrBasename, filepath.Ext(adrBasename)) + ".rule"
		ruleFilePath = filepath.Join(adrDir, ruleFilename)
	}

	// Write the file
	err = os.WriteFile(ruleFilePath, []byte(ruleContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write rule file: %w", err)
	}

	i.output.RuleGenerated(decision.ID, ruleFilePath)
	return nil
}

func generateRuleTemplate(id, title string) string {
	return fmt.Sprintf(`adr "%s" "%s"

# component "MyComponent" = "com.example.mypackage"
# path "MyPath" = "src/mypackage"

code "rule_name" {
  # MyComponent must not depend on MyOtherComponent
  severity error
}

file "rule_name" {
  # path "**/*.go" must exist
  severity error
}
`, id, title)
}
