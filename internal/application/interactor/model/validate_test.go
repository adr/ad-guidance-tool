package model

import (
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate_Success(t *testing.T) {
	mockSvc := new(svc_mocks.ModelService)
	mockOutput := new(out_mocks.ModelValidate)

	modelPath := "my/model"

	mockSvc.On("ValidateIndexDataCorrectness", modelPath).Return(nil)
	mockSvc.On("ValidateDecisionDataCorrectness", modelPath).Return(nil)
	mockOutput.On("ModelValidated", modelPath, nil, nil).Return()

	interactor := NewModelValidateInteractor(mockSvc, mockOutput)

	err := interactor.Validate(modelPath)

	assert.NoError(t, err)
	mockSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestValidate_IndexError(t *testing.T) {
	mockSvc := new(svc_mocks.ModelService)
	mockOutput := new(out_mocks.ModelValidate)

	modelPath := "invalid/model"
	indexErr := errors.New("index is broken")

	mockSvc.On("ValidateIndexDataCorrectness", modelPath).Return(indexErr)
	// Data validation is skipped if indexErr != nil
	mockOutput.On("ModelValidated", modelPath, indexErr, nil).Return()

	interactor := NewModelValidateInteractor(mockSvc, mockOutput)

	err := interactor.Validate(modelPath)

	assert.NoError(t, err)
	mockSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestValidate_DataError(t *testing.T) {
	mockSvc := new(svc_mocks.ModelService)
	mockOutput := new(out_mocks.ModelValidate)

	modelPath := "model"
	dataErr := errors.New("data mismatch")

	mockSvc.On("ValidateIndexDataCorrectness", modelPath).Return(nil)
	mockSvc.On("ValidateDecisionDataCorrectness", modelPath).Return(dataErr)
	mockOutput.On("ModelValidated", modelPath, nil, dataErr).Return()

	interactor := NewModelValidateInteractor(mockSvc, mockOutput)

	err := interactor.Validate(modelPath)

	assert.NoError(t, err)
	mockSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}
