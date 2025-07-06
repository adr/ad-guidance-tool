package model

import (
	domain "adg/internal/domain/decision"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type FileModelRepository struct{}

func NewFileModelRepository() *FileModelRepository {
	return &FileModelRepository{}
}

func (r *FileModelRepository) CreateModel(modelPath string) error {
	if err := os.MkdirAll(modelPath, 0755); err != nil {
		return fmt.Errorf("failed to create model directory: %w", err)
	}
	return nil
}

func (r *FileModelRepository) CreateIndex(modelPath string) error {
	indexPath := filepath.Join(modelPath, "index.yaml")
	indexData := map[string]any{
		"decisions": map[string]any{},
	}
	content, err := yaml.Marshal(indexData)
	if err != nil {
		return fmt.Errorf("failed to serialize empty index: %w", err)
	}
	return os.WriteFile(indexPath, content, 0644)
}

func (r *FileModelRepository) RebuildIndex(modelPath string, decisions []domain.Decision) error {
	indexPath := filepath.Join(modelPath, "index.yaml")

	// directly map all decisions
	indexData := make(map[string]domain.Decision)
	for _, decision := range decisions {
		indexData[decision.ID] = decision
	}

	// marshal and save
	yamlOut, err := yaml.Marshal(map[string]interface{}{"decisions": indexData})
	if err != nil {
		return err
	}
	return os.WriteFile(indexPath, yamlOut, 0644)
}

func (r *FileModelRepository) Exists(modelPath string) bool {
	indexPath := filepath.Join(modelPath, "index.yaml")
	_, err := os.Stat(indexPath)
	return err == nil
}
