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
	issues, err := i.service.Validate(modelPath)
	if err != nil {
		return err
	}

	out := make([]outputport.ValidationIssue, 0, len(issues))
	for _, issue := range issues {
		out = append(out, outputport.ValidationIssue{ID: issue.ID, Message: issue.Message})
	}
	i.output.ModelValidated(modelPath, out)
	return nil
}
