package model

import (
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInit_ModelAlreadyExists(t *testing.T) {
	mockService := new(svc_mocks.ModelService)
	mockOutput := new(out_mocks.ModelInit)

	modelPath := "some/path"

	mockService.On("Exists", modelPath).Return(true)

	interactor := NewInitModelInteractor(mockService, mockOutput)
	err := interactor.Init(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already contains a model")
	mockService.AssertExpectations(t)
	mockOutput.AssertNotCalled(t, "Initialized", mock.Anything)
}

func TestInit_CreateModelFails(t *testing.T) {
	mockService := new(svc_mocks.ModelService)
	mockOutput := new(out_mocks.ModelInit)

	modelPath := "some/path"

	mockService.On("Exists", modelPath).Return(false)
	mockService.On("CreateModel", modelPath).Return(errors.New("disk error"))

	interactor := NewInitModelInteractor(mockService, mockOutput)
	err := interactor.Init(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create model")
	mockService.AssertExpectations(t)
	mockOutput.AssertNotCalled(t, "Initialized", mock.Anything)
}

func TestInit_Success(t *testing.T) {
	mockService := new(svc_mocks.ModelService)
	mockOutput := new(out_mocks.ModelInit)

	modelPath := "some/path"

	mockService.On("Exists", modelPath).Return(false)
	mockService.On("CreateModel", modelPath).Return(nil)
	mockOutput.On("Initialized", modelPath).Return(nil)

	interactor := NewInitModelInteractor(mockService, mockOutput)
	err := interactor.Init(modelPath)

	assert.NoError(t, err)
	mockService.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}
