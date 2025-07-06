package model

import (
	"adg/internal/application/inputport"
	"adg/internal/application/outputport"
	decisiondomain "adg/internal/domain/decision"
	modeldomain "adg/internal/domain/model"
	"fmt"
)

type CopyModelInteractor struct {
	modelService    modeldomain.ModelService
	decisionService decisiondomain.DecisionService
	output          outputport.ModelCopy
}

func NewCopyModelInteractor(
	modelService modeldomain.ModelService,
	decisionService decisiondomain.DecisionService,
	output outputport.ModelCopy,
) inputport.ModelCopy {
	return &CopyModelInteractor{
		modelService:    modelService,
		decisionService: decisionService,
		output:          output,
	}
}

func (i *CopyModelInteractor) Copy(modelPath, targetPath string, filters map[string][]string) error {
	if i.modelService.Exists(targetPath) {
		return fmt.Errorf("can not copy model, target directory %q already contains a model (use import for copying decisions to an existing model)", targetPath)
	}

	decisions, err := i.decisionService.GetAllDecisions(modelPath)
	if err != nil {
		return fmt.Errorf("failed to load decisions: %w", err)
	}

	if len(filters) > 0 {
		decisions, err = i.decisionService.FilterDecisions(decisions, filters)
		if err != nil {
			return fmt.Errorf("failed to apply filters: %w", err)
		}

		// TODO: validate that filtered decisions do not have any links to a decision that will not be imported
	}

	if err := i.modelService.CreateModel(targetPath); err != nil {
		return fmt.Errorf("failed to create model directory: %w", err)
	}

	for _, d := range decisions {
		if err := i.decisionService.Copy(modelPath, targetPath, d.ID); err != nil {
			return fmt.Errorf("failed to copy decision %s: %w", d.ID, err)
		}
	}

	if err := i.modelService.RebuildIndex(targetPath); err != nil {
		return fmt.Errorf("failed to rebuild index: %w", err)
	}

	i.output.Copied(modelPath, targetPath, len(decisions))
	return nil
}
