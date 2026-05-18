package model

import (
	"errors"
	"testing"

	"adg/internal/application/outputport"
	modeldomain "adg/internal/domain/model"
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"

	"github.com/stretchr/testify/assert"
)

func TestValidate_NoIssues(t *testing.T) {
	mockSvc := new(svc_mocks.ModelService)
	mockOutput := new(out_mocks.ModelValidate)

	modelPath := "model"

	mockSvc.On("Validate", modelPath).Return([]modeldomain.ValidationIssue{}, nil)
	mockOutput.On("ModelValidated", modelPath, []outputport.ValidationIssue{}).Return()

	interactor := NewModelValidateInteractor(mockSvc, mockOutput)

	err := interactor.Validate(modelPath)

	assert.NoError(t, err)
	mockSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestValidate_MapsIssues(t *testing.T) {
	mockSvc := new(svc_mocks.ModelService)
	mockOutput := new(out_mocks.ModelValidate)

	modelPath := "model"

	domainIssues := []modeldomain.ValidationIssue{
		{ID: "0001", Message: "filename does not match NNNN-slug.md"},
		{ID: "0002", Message: "missing required section: Decision Outcome"},
	}
	expectedOut := []outputport.ValidationIssue{
		{ID: "0001", Message: "filename does not match NNNN-slug.md"},
		{ID: "0002", Message: "missing required section: Decision Outcome"},
	}

	mockSvc.On("Validate", modelPath).Return(domainIssues, nil)
	mockOutput.On("ModelValidated", modelPath, expectedOut).Return()

	interactor := NewModelValidateInteractor(mockSvc, mockOutput)

	err := interactor.Validate(modelPath)

	assert.NoError(t, err)
	mockSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestValidate_PropagatesError(t *testing.T) {
	mockSvc := new(svc_mocks.ModelService)
	mockOutput := new(out_mocks.ModelValidate)

	modelPath := "model"

	mockSvc.On("Validate", modelPath).Return(nil, errors.New("read error"))

	interactor := NewModelValidateInteractor(mockSvc, mockOutput)

	err := interactor.Validate(modelPath)

	assert.ErrorContains(t, err, "read error")
	mockSvc.AssertExpectations(t)
	mockOutput.AssertNotCalled(t, "ModelValidated")
}
