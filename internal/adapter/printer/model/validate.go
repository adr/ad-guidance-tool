package model

import (
	"adg/internal/application/outputport"
	"fmt"
)

type ModelValidatePresenter struct{}

func NewModelValidatePresenter() *ModelValidatePresenter {
	return &ModelValidatePresenter{}
}

// ModelValidated prints per-decision issue reports. On a clean run nothing is
// printed for OK decisions in this PR — that's expedient; the `--quiet` flag
// from §C.6 lands in PR 2 where the output style is finalized. A non-empty
// issue list still results in a non-zero command exit because the cmd layer
// inspects the printer's "had issues" state via the returned bool — for now
// the presenter just prints; the exit-code mapping is also a PR 2 concern.
func (p *ModelValidatePresenter) ModelValidated(modelName string, issues []outputport.ValidationIssue) {
	if len(issues) == 0 {
		fmt.Printf("%s model is valid\n", modelName)
		return
	}
	fmt.Printf("%s model has %d validation issue(s):\n", modelName, len(issues))
	for _, issue := range issues {
		fmt.Printf("  ID %s: %s\n", issue.ID, issue.Message)
	}
}
