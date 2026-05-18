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
	var bodies []outputport.DecisionBody

	for _, id := range ids {
		body, err := i.service.GetBody(modelPath, id)
		if err != nil {
			return fmt.Errorf("failed to load body for ID %q: %w", id, err)
		}
		bodies = append(bodies, outputport.DecisionBody{ID: id, Body: body})
	}

	for _, title := range titles {
		decision, err := i.service.GetDecisionByTitle(modelPath, title)
		if err != nil {
			return fmt.Errorf("failed to resolve title %q: %w", title, err)
		}
		body, err := i.service.GetBody(modelPath, decision.ID)
		if err != nil {
			return fmt.Errorf("failed to load body for title %q: %w", title, err)
		}
		bodies = append(bodies, outputport.DecisionBody{ID: decision.ID, Body: body})
	}

	i.output.Printed(bodies, sections)
	return nil
}
