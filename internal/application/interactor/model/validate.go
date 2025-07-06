package model

import (
	"adg/internal/application/inputport"
	"adg/internal/application/outputport"
	domain "adg/internal/domain/model"
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
