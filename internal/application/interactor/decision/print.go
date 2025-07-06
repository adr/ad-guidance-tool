package decision

import (
	"adg/internal/application/inputport"
	"adg/internal/application/outputport"
	domain "adg/internal/domain/decision"
	"fmt"
)

type PrintDecisionsInteractor struct {
	service domain.DecisionService
	output  outputport.DecisionPrint
}

func NewPrintDecisionsInteractor(service domain.DecisionService, output outputport.DecisionPrint) inputport.DecisionPrint {
	return &PrintDecisionsInteractor{
		service: service,
		output:  output,
	}
}

func (i *PrintDecisionsInteractor) Print(modelPath string, ids []string, titles []string, sections map[string]bool) error {
	var contents []domain.DecisionContent

	for _, id := range ids {
		content, err := i.service.GetDecisionContent(modelPath, id)
		if err != nil {
			return fmt.Errorf("failed to load content for ID %q: %w", id, err)
		}
		contents = append(contents, *content)
	}

	for _, title := range titles {
		decision, err := i.service.GetDecisionByTitle(modelPath, title)
		if err != nil {
			return fmt.Errorf("failed to resolve title %q: %w", title, err)
		}
		content, err := i.service.GetDecisionContent(modelPath, decision.ID)
		if err != nil {
			return fmt.Errorf("failed to load content for title %q: %w", title, err)
		}
		contents = append(contents, *content)
	}

	i.output.Printed(contents, sections)
	return nil
}
