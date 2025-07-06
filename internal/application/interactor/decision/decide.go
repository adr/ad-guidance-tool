package decision

import (
	"adg/internal/application/inputport"
	util "adg/internal/application/interactor"
	"adg/internal/application/outputport"
	domain "adg/internal/domain/decision"
	"fmt"
)

type DecideDecisionInteractor struct {
	service domain.DecisionService
	output  outputport.DecisionDecide
}

func NewDecideInteractor(service domain.DecisionService, output outputport.DecisionDecide) inputport.DecisionDecide {
	return &DecideDecisionInteractor{
		service: service,
		output:  output,
	}
}

func (i *DecideDecisionInteractor) Decide(modelPath, id, title, option, reason, author string, enforceOption bool) error {
	var (
		decision *domain.Decision
		err      error
	)

	decision, err = util.ResolveDecisionByIdOrTitle(modelPath, id, title, i.service)
	if err != nil {
		return err
	}

	if decision.Status == "decided" {
		return fmt.Errorf("decision has already been decided, revise the decision to create a copy that is still open")
	}

	if err := i.service.Decide(modelPath, decision, option, reason, enforceOption); err != nil {
		return err
	}

	if err := i.service.Comment(modelPath, decision, author, "marked decision as decided"); err != nil {
		return err
	}

	i.output.Decided(decision.ID)
	return nil
}
