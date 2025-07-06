package decision

import (
	"adg/internal/application/inputport"
	"adg/internal/application/outputport"
	decisiondomain "adg/internal/domain/decision"
	modeldomain "adg/internal/domain/model"
	"fmt"
)

type AddDecisionsInteractor struct {
	modelService    modeldomain.ModelService
	decisionService decisiondomain.DecisionService
	output          outputport.DecisionAdd
}

func NewAddDecisionsInteractor(
	modelService modeldomain.ModelService,
	decisionService decisiondomain.DecisionService,
	output outputport.DecisionAdd,
) inputport.DecisionAdd {
	return &AddDecisionsInteractor{
		modelService:    modelService,
		decisionService: decisionService,
		output:          output,
	}
}

func (i *AddDecisionsInteractor) Add(modelPath string, titles []string) error {
	var successes []*decisiondomain.Decision
	failures := make(map[string]error)

	if !i.modelService.Exists(modelPath) {
		return fmt.Errorf("can not add decisions, index of model %q does not exist (use rebuild to recreate index file if directory actually does contain decisions)", modelPath)
	}

	for _, title := range titles {
		decision, err := i.decisionService.AddNew(modelPath, title)
		if err != nil {
			failures[title] = err
			continue
		}
		successes = append(successes, decision)
	}

	i.output.Added(successes, failures)
	return nil
}
