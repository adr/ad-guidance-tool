package model

import (
	"adg/internal/domain"
	decisiondomain "adg/internal/domain/decision"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type ModelService interface {
	CreateModel(modelPath string) error
	RebuildIndex(modelPath string) error
	Exists(modelPath string) bool
	ValidateIndexDataCorrectness(modelPath string) error
	ValidateDecisionDataCorrectness(modelPath string) error
}

type ModelServiceImplementation struct {
	modelRepo    ModelRepository
	decisionRepo decisiondomain.DecisionRepository
}

func NewModelService(modelRepo ModelRepository, decisionRepo decisiondomain.DecisionRepository) ModelService {
	return &ModelServiceImplementation{
		modelRepo:    modelRepo,
		decisionRepo: decisionRepo,
	}
}

func (s *ModelServiceImplementation) CreateModel(modelPath string) error {
	if err := s.modelRepo.CreateModel(modelPath); err != nil {
		return err
	}
	return s.modelRepo.CreateIndex(modelPath)
}

func (s *ModelServiceImplementation) RebuildIndex(modelPath string) error {
	decisions, err := s.decisionRepo.LoadAllByData(modelPath)
	if err != nil {
		return fmt.Errorf("failed to load decisions: %w", err)
	}

	// check for duplicate decision IDs
	seen := make(map[string]bool)
	for _, d := range decisions {
		if seen[d.ID] {
			return fmt.Errorf("cannot rebuild index, duplicate decision ID %s detected in directory or sub directories", d.ID)
		}
		seen[d.ID] = true
	}

	return s.modelRepo.RebuildIndex(modelPath, decisions)
}

func (s *ModelServiceImplementation) Exists(modelPath string) bool {
	return s.modelRepo.Exists(modelPath)
}

func (s *ModelServiceImplementation) ValidateIndexDataCorrectness(modelPath string) error {
	dataDecisions, err := s.decisionRepo.LoadAllByData(modelPath)
	if err != nil {
		return fmt.Errorf("failed to load decisions from files: %w", err)
	}

	if dupID, ok := findDuplicateID(dataDecisions); ok {
		return fmt.Errorf("duplicate decision ID found: %s", dupID)
	}

	indexDecisions, err := s.decisionRepo.LoadAllByIndex(modelPath)
	if err != nil {
		return fmt.Errorf("failed to load decisions from index: %w (run 'rebuild' to recreate the index file)", err)
	}

	dataMap := indexByID(dataDecisions)
	indexMap := indexByID(indexDecisions)

	allIDs := uniqueKeys(dataMap, indexMap)
	var errorsFound bool

	for _, id := range allIDs {
		data, inData := dataMap[id]
		index, inIndex := indexMap[id]

		switch {
		case !inData:
			fmt.Printf("ID %s exists in index but not in data (recreate the decision or fix the metadata)\n", id)
			errorsFound = true
		case !inIndex:
			fmt.Printf("ID %s exists in data but not in index (run 'rebuild' to update index file)\n", id)
			errorsFound = true
		case !metadataEqual(data, index):
			fmt.Printf("ID %s metadata mismatch between file and index (run 'rebuild' to update index file)\n", id)
			errorsFound = true
		default:
			fmt.Printf("ID %s metadata is consistent with index\n", id)
		}
	}

	if errorsFound {
		return fmt.Errorf("validation of metadata completed with mismatches")
	}

	return nil
}

func (s *ModelServiceImplementation) ValidateDecisionDataCorrectness(modelPath string) error {
	decisions, err := s.decisionRepo.LoadAllByIndex(modelPath)
	if err != nil {
		return fmt.Errorf("failed to load decisions from index: %w", err)
	}

	sort.Slice(decisions, func(i, j int) bool {
		return decisions[i].ID < decisions[j].ID
	})

	var errorsFound bool

	for _, d := range decisions {
		content, err := s.decisionRepo.LoadDecisionContentRaw(modelPath, d.ID)
		if err != nil {
			return fmt.Errorf("failed to load decision content from files: %w", err)
		}

		missing := validateRequiredAnchors(content)
		if len(missing) > 0 {
			errorsFound = true
			for _, tag := range missing {
				fmt.Printf("ID %s is missing required tag: %s\n", d.ID, tag)
			}
		} else {
			fmt.Printf("ID %s has valid section tags\n", d.ID)
		}
	}

	if errorsFound {
		return fmt.Errorf("validation of file contents completed with errors")
	}

	return nil
}

// Helper methods
func findDuplicateID(decisions []decisiondomain.Decision) (string, bool) {
	seen := make(map[string]bool)
	for _, d := range decisions {
		if seen[d.ID] {
			return d.ID, true
		}
		seen[d.ID] = true
	}
	return "", false
}

func indexByID(decisions []decisiondomain.Decision) map[string]decisiondomain.Decision {
	m := make(map[string]decisiondomain.Decision)
	for _, d := range decisions {
		m[d.ID] = d
	}
	return m
}

func uniqueKeys(m1, m2 map[string]decisiondomain.Decision) []string {
	keys := make(map[string]struct{})
	for k := range m1 {
		keys[k] = struct{}{}
	}
	for k := range m2 {
		keys[k] = struct{}{}
	}
	all := make([]string, 0, len(keys))
	for k := range keys {
		all = append(all, k)
	}
	sort.Strings(all)
	return all
}

func metadataEqual(a, b decisiondomain.Decision) bool {
	return a.ID == b.ID &&
		a.Title == b.Title &&
		a.Status == b.Status &&
		reflect.DeepEqual(a.Tags, b.Tags) &&
		reflect.DeepEqual(a.Links, b.Links) &&
		reflect.DeepEqual(a.Comments, b.Comments)
}

func validateRequiredAnchors(content string) []string {
	required := []string{
		domain.AnchorForSection(domain.AnchorSectionQuestion),
		domain.AnchorForSection(domain.AnchorSectionOptions),
		domain.AnchorForSection(domain.AnchorSectionCriteria),
	}

	var missing []string
	for _, anchor := range required {
		if !strings.Contains(content, anchor) {
			missing = append(missing, anchor)
		}
	}
	return missing
}
