package model

// ModelRepository is the persistence boundary for the model directory itself.
//
// index.yaml is dropped in this fork; CreateIndex and RebuildIndex from the
// legacy interface are gone. The repository's only remaining job is creating
// the directory and checking existence — ADR files within are managed by
// DecisionRepository (decision package).
type ModelRepository interface {
	CreateModel(modelPath string) error
	Exists(modelPath string) bool
}
