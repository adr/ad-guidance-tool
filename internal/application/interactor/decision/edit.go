package decision

import (
	"adg/internal/application/inputport"
	util "adg/internal/application/interactor"
	"adg/internal/application/outputport"
	domain "adg/internal/domain/decision"
)

type EditDecisionInteractor struct {
	service domain.DecisionService
	output  outputport.DecisionEdit
}

func NewEditDecisionInteractor(service domain.DecisionService, output outputport.DecisionEdit) inputport.DecisionEdit {
	return &EditDecisionInteractor{
		service: service,
		output:  output,
	}
}

func (i *EditDecisionInteractor) Edit(modelPath, id, title string, question *string, options *[]string, criteria *string) error {
	var (
		decision *domain.Decision
		err      error
	)

	decision, err = util.ResolveDecisionByIdOrTitle(modelPath, id, title, i.service)
	if err != nil {
		return err
	}

	if err := i.service.Edit(modelPath, decision, question, options, criteria); err != nil {
		return err
	}

	i.output.Edited(decision.ID)
	return nil
}
