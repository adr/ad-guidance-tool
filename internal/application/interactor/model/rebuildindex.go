package model

import (
	"adg/internal/application/inputport"
	"adg/internal/application/outputport"
	domain "adg/internal/domain/model"
)

// RebuildIndexInteractor exists in PR 1b as a no-op shim. The fork drops
// index.yaml entirely; ADR files are the only source of truth. The interactor
// is retained so the `adg rebuild` CLI surface continues to exist for
// backwards-compat with anyone scripting against it — it simply reports
// success and does nothing. PR 2 (command port) may remove the command
// entirely.
type RebuildIndexInteractor struct {
	service domain.ModelService
	output  outputport.ModelRebuildIndex
}

func NewRebuildIndexInteractor(
	service domain.ModelService,
	output outputport.ModelRebuildIndex,
) inputport.ModelRebuildIndex {
	return &RebuildIndexInteractor{
		service: service,
		output:  output,
	}
}

func (i *RebuildIndexInteractor) RebuildIndex(modelPath string) error {
	// No-op: index.yaml is no longer used.
	i.output.IndexRebuilt(modelPath)
	return nil
}
