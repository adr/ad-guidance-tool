package decision

import (
	"adg/internal/application/inputport"
	util "adg/internal/application/interactor"
	"adg/internal/application/outputport"
	domain "adg/internal/domain/decision"
	"fmt"
)

type ReviseDecisionInteractor struct {
	service domain.DecisionService
	output  outputport.DecisionRevise
}

func NewReviseDecisionInteractor(service domain.DecisionService, output outputport.DecisionRevise) inputport.DecisionRevise {
	return &ReviseDecisionInteractor{
		service: service,
		output:  output,
	}
}

func (i *ReviseDecisionInteractor) ReviseDecision(modelPath, id, title string) error {
	original, err := util.ResolveDecisionByIdOrTitle(modelPath, id, title, i.service)
	if err != nil {
		return fmt.Errorf("failed to find original decision: %w", err)
	}

	revised, err := i.service.Revise(modelPath, original)
	if err != nil {
		return fmt.Errorf("failed to revise decision: %w", err)
	}

	if err := i.service.Link(modelPath, original, revised, "revised by", "revises"); err != nil {
		return fmt.Errorf("failed to link revised decision to original: %w", err)
	}

	i.output.Revised(original.ID, revised.ID)
	return nil
}
