package decision

import (
	"adg/internal/application/inputport"
	"adg/internal/application/outputport"
	domain "adg/internal/domain/decision"
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
