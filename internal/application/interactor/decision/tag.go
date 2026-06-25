package decision

import (
	"github.com/adr/ad-guidance-tool/internal/application/inputport"
	util "github.com/adr/ad-guidance-tool/internal/application/interactor"
	"github.com/adr/ad-guidance-tool/internal/application/outputport"
	domain "github.com/adr/ad-guidance-tool/internal/domain/decision"
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
