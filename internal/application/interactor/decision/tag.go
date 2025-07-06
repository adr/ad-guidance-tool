package decision

import (
	"adg/internal/application/inputport"
	util "adg/internal/application/interactor"
	"adg/internal/application/outputport"
	domain "adg/internal/domain/decision"
	"fmt"
)

type TagDecisionInteractor struct {
	service domain.DecisionService
	output  outputport.DecisionTag
}

func NewTagDecisionInteractor(service domain.DecisionService, output outputport.DecisionTag) inputport.DecisionTag {
	return &TagDecisionInteractor{
		service: service,
		output:  output,
	}
}

func (i *TagDecisionInteractor) Tag(modelPath, id, title string, tags []string) error {
	decision, err := util.ResolveDecisionByIdOrTitle(modelPath, id, title, i.service)
	if err != nil {
		return err
	}

	for _, tag := range tags {
		if err := i.service.Tag(modelPath, decision, tag); err != nil {
			return fmt.Errorf("failed to tag decision with tag %q: %w", tag, err)
		}
	}

	i.output.Tagged(decision.ID, tags)
	return nil
}
