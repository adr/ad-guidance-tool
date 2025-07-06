package decision

import (
	"adg/internal/application/inputport"
	util "adg/internal/application/interactor"
	"adg/internal/application/outputport"
	domain "adg/internal/domain/decision"
	"fmt"
)

type LinkDecisionsInteractor struct {
	service domain.DecisionService
	output  outputport.DecisionLink
}

func NewLinkDecisionsInteractor(service domain.DecisionService, output outputport.DecisionLink) inputport.DecisionLink {
	return &LinkDecisionsInteractor{
		service: service,
		output:  output,
	}
}

func (i *LinkDecisionsInteractor) Link(
	modelPath string,
	sourceID, sourceTitle string,
	targetID, targetTitle string,
	tag, reverseTag string,
) error {
	var (
		source *domain.Decision
		target *domain.Decision
		err    error
	)

	source, err = util.ResolveDecisionByIdOrTitle(modelPath, sourceID, sourceTitle, i.service)
	if err != nil {
		return fmt.Errorf("could not find source decision: %w", err)
	}

	target, err = util.ResolveDecisionByIdOrTitle(modelPath, targetID, targetTitle, i.service)
	if err != nil {
		return fmt.Errorf("could not find target decision: %w", err)
	}

	if source.ID == target.ID {
		return fmt.Errorf("source and target decision are the same, cannot create a tag from a decision to itself")
	}

	if err := i.service.Link(modelPath, source, target, tag, reverseTag); err != nil {
		return fmt.Errorf("linking failed: %w", err)
	}

	i.output.Linked(source.ID, target.ID, tag, reverseTag)
	return nil
}
