package model

import (
	"adg/internal/application/inputport"
	"adg/internal/application/outputport"
	domain "adg/internal/domain/model"
	"fmt"
)

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
	if err := i.service.RebuildIndex(modelPath); err != nil {
		return fmt.Errorf("failed to rebuild index: %w", err)
	}

	i.output.IndexRebuilt(modelPath)
	return nil
}
