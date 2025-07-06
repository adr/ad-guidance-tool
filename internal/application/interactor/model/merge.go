package model

import (
	"adg/internal/application/inputport"
	"adg/internal/application/outputport"
	decisiondomain "adg/internal/domain/decision"
	modeldomain "adg/internal/domain/model"
	"fmt"
	"sort"
	"strconv"
)

type MergeModelsInteractor struct {
	modelService    modeldomain.ModelService
	decisionService decisiondomain.DecisionService
	output          outputport.ModelMerge
}

func NewMergeModelsInteractor(
	modelService modeldomain.ModelService,
	decisionService decisiondomain.DecisionService,
	output outputport.ModelMerge,
) inputport.ModelMerge {
	return &MergeModelsInteractor{
		modelService:    modelService,
		decisionService: decisionService,
		output:          output,
	}
}

func (i *MergeModelsInteractor) Merge(modelAPath, modelBPath, targetPath string, filters map[string][]string) error {
	if i.modelService.Exists(targetPath) {
		return fmt.Errorf("can not merge models, target directory %q already contains a model (use import for copying decisions to an existing model)", targetPath)
	}

	if err := i.modelService.CreateModel(targetPath); err != nil {
		return fmt.Errorf("failed to create model: %w", err)
	}

	modelADecisions, err := i.decisionService.GetAllDecisions(modelAPath)
	if err != nil {
		return fmt.Errorf("failed to load decisions from model %q: %w", modelAPath, err)
	}

	highestID, err := getHighestID(modelADecisions)
	if err != nil {
		return fmt.Errorf("invalid ID in model %q: %w", modelAPath, err)
	}

	amountA, err := i.copyDecisions(modelAPath, targetPath, filters, 0)
	if err != nil {
		return fmt.Errorf("failed to merge decisions from model %q: %w", modelAPath, err)
	}

	amountB, err := i.copyDecisions(modelBPath, targetPath, filters, highestID)
	if err != nil {
		return fmt.Errorf("failed to merge decisions from model %q: %w", modelBPath, err)
	}

	if err := i.modelService.RebuildIndex(targetPath); err != nil {
		return fmt.Errorf("failed to rebuild index: %w", err)
	}

	i.output.Merged(modelAPath, modelBPath, targetPath, amountA+amountB)
	return nil
}

func (i *MergeModelsInteractor) copyDecisions(fromModel, toModel string, filters map[string][]string, increment int) (int, error) {
	decisions, err := i.decisionService.GetAllDecisions(fromModel)
	if err != nil {
		return 0, fmt.Errorf("failed to load decisions from %q: %w", fromModel, err)
	}

	if len(filters) > 0 {
		decisions, err = i.decisionService.FilterDecisions(decisions, filters)
		if err != nil {
			return 0, fmt.Errorf("failed to apply filters: %w", err)
		}
		// TODO: validate that filtered decisions do not have any links to a decision that will not be imported
	}

	sort.Slice(decisions, func(a, b int) bool {
		return decisions[a].ID < decisions[b].ID
	})

	for _, d := range decisions {
		content, err := i.decisionService.GetDecisionContent(fromModel, d.ID)
		if err != nil {
			return 0, fmt.Errorf("failed to load decision content from %q: %w", fromModel, err)
		}

		if _, err := i.decisionService.AddExisting(fromModel, toModel, &d, content, increment); err != nil {
			return 0, fmt.Errorf("failed to add decision from %q: %w", fromModel, err)
		}
	}

	return len(decisions), nil
}

func getHighestID(decisions []decisiondomain.Decision) (int, error) {
	highest := 0
	for _, d := range decisions {
		id, err := strconv.Atoi(d.ID)
		if err != nil {
			return 0, err
		}
		if id > highest {
			highest = id
		}
	}
	return highest, nil
}
