package decision

import (
	"github.com/adr/ad-guidance-tool/internal/application/inputport"
	util "github.com/adr/ad-guidance-tool/internal/application/interactor"
	"github.com/adr/ad-guidance-tool/internal/application/outputport"
	domain "github.com/adr/ad-guidance-tool/internal/domain/decision"
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
