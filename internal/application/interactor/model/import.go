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

type ImportModelInteractor struct {
	modelService    modeldomain.ModelService
	decisionService decisiondomain.DecisionService
	output          outputport.ModelImport
}

func NewImportModelInteractor(
	modelService modeldomain.ModelService,
	decisionService decisiondomain.DecisionService,
	output outputport.ModelImport,
) inputport.ModelImport {
	return &ImportModelInteractor{
		modelService:    modelService,
		decisionService: decisionService,
		output:          output,
	}
}

func (i *ImportModelInteractor) Import(sourcePath, targetPath string, filters map[string][]string) error {
	if !i.modelService.Exists(targetPath) {
		return fmt.Errorf("can not import model, target directory %q does not contain a model (use copy for creating a completely new model)", targetPath)
	}

	highestID, err := i.calculateHighestID(targetPath)
	if err != nil {
		return err
	}

	decisions, err := i.loadAndFilterDecisions(sourcePath, filters)
	if err != nil {
		return err
	}

	if len(decisions) == 0 {
		i.output.Imported(sourcePath, targetPath, 0)
		return nil
	}

	sort.Slice(decisions, func(a, b int) bool {
		return decisions[a].ID < decisions[b].ID
	})

	for _, d := range decisions {
		if err := i.importDecision(sourcePath, targetPath, d, highestID); err != nil {
			return err
		}
	}

	if err := i.modelService.RebuildIndex(targetPath); err != nil {
		return fmt.Errorf("failed to rebuild target model index: %w", err)
	}

	i.output.Imported(sourcePath, targetPath, len(decisions))
	return nil
}

// determines the max ID in the existing target model
func (i *ImportModelInteractor) calculateHighestID(targetPath string) (int, error) {
	decisions, err := i.decisionService.GetAllDecisions(targetPath)
	if err != nil {
		return 0, fmt.Errorf("failed to load decisions from source model: %w", err)
	}

	highest := 0
	for _, d := range decisions {
		id, err := strconv.Atoi(d.ID)
		if err != nil {
			return 0, fmt.Errorf("invalid decision ID %q: %w", d.ID, err)
		}
		if id > highest {
			highest = id
		}
	}
	return highest, nil
}

// fetches and filters decisions from the source model
func (i *ImportModelInteractor) loadAndFilterDecisions(sourcePath string, filters map[string][]string) ([]decisiondomain.Decision, error) {
	decisions, err := i.decisionService.GetAllDecisions(sourcePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load decisions from source model: %w", err)
	}

	if len(filters) > 0 {
		decisions, err = i.decisionService.FilterDecisions(decisions, filters)
		if err != nil {
			return nil, fmt.Errorf("failed to apply filters: %w", err)
		}

		// TODO: validate that filtered decisions do not have any links to a decision that will not be imported
	}

	return decisions, nil
}

// adds a single decision from source into target, adjusting IDs
func (i *ImportModelInteractor) importDecision(sourcePath, targetPath string, d decisiondomain.Decision, increment int) error {
	content, err := i.decisionService.GetDecisionContent(sourcePath, d.ID)
	if err != nil {
		return fmt.Errorf("failed to load decision content for import: %w", err)
	}

	if _, err := i.decisionService.AddExisting(sourcePath, targetPath, &d, content, increment); err != nil {
		return fmt.Errorf("failed to add decision into target model: %w", err)
	}
	return nil
}
