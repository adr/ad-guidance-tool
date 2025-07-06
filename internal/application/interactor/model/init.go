package model

import (
	"adg/internal/application/inputport"
	"adg/internal/application/outputport"
	domain "adg/internal/domain/model"
	"fmt"
)

type CreateModelInteractor struct {
	service domain.ModelService
	output  outputport.ModelInit
}

func NewInitModelInteractor(
	service domain.ModelService,
	output outputport.ModelInit,
) inputport.ModelInit {
	return &CreateModelInteractor{
		service: service,
		output:  output,
	}
}

func (i *CreateModelInteractor) Init(modelPath string) error {
	if i.service.Exists(modelPath) {
		return fmt.Errorf("can not initialize new model, target directory %q already contains a model", modelPath)
	}

	if err := i.service.CreateModel(modelPath); err != nil {
		return fmt.Errorf("failed to create model: %w", err)
	}

	i.output.Initialized(modelPath)
	return nil
}
