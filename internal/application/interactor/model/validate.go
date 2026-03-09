package model

import (
	"github.com/adr/ad-guidance-tool/internal/application/inputport"
	"github.com/adr/ad-guidance-tool/internal/application/outputport"
	domain "github.com/adr/ad-guidance-tool/internal/domain/model"
)

type ModelValidateInteractor struct {
	service domain.ModelService
	output  outputport.ModelValidate
}

func NewModelValidateInteractor(
	service domain.ModelService,
	output outputport.ModelValidate,
) inputport.ModelValidate {
	return &ModelValidateInteractor{
		service: service,
		output:  output,
	}
}

func (i *ModelValidateInteractor) Validate(modelPath string) error {
	indexErr := i.service.ValidateIndexDataCorrectness(modelPath)

	// TODO: also validate that decision metadata is correct (all required fields are available)

	var dataErr error
	if indexErr == nil {
		dataErr = i.service.ValidateDecisionDataCorrectness(modelPath)
	}

	i.output.ModelValidated(modelPath, indexErr, dataErr)

	return nil
}
