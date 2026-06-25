package decision

import (
	"github.com/adr/ad-guidance-tool/internal/application/inputport"
	"github.com/adr/ad-guidance-tool/internal/application/outputport"
	domain "github.com/adr/ad-guidance-tool/internal/domain/decision"
)

type ListDecisionsInteractor struct {
	service domain.DecisionService
	output  outputport.DecisionList
}

func NewListDecisionsInteractor(service domain.DecisionService, output outputport.DecisionList) inputport.DecisionList {
	return &ListDecisionsInteractor{
		service: service,
		output:  output,
	}
}

func (i *ListDecisionsInteractor) ListDecisions(modelPath string, filters map[string][]string, format string) error {
	decisions, err := i.service.GetAllDecisions(modelPath)
	if err != nil {
		return err
	}

	if len(filters) > 0 {
		decisions, err = i.service.FilterDecisions(decisions, filters)
		if err != nil {
			return err
		}
	}

	i.output.Listed(decisions, format)
	return nil
}
