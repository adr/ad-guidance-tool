package model

import domain "adg/internal/domain/decision"

type ModelRepository interface {
	CreateModel(modelPath string) error
	CreateIndex(modelPath string) error
	RebuildIndex(modelPath string, decisions []domain.Decision) error
	Exists(modelPath string) bool
}
