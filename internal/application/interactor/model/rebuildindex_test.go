package model

import (
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRebuildIndex_Success(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockOutput := new(out_mocks.ModelRebuildIndex)

	modelPath := "some/model"

	mockModelSvc.On("RebuildIndex", modelPath).Return(nil)
	mockOutput.On("IndexRebuilt", modelPath).Return()

	interactor := NewRebuildIndexInteractor(mockModelSvc, mockOutput)

	err := interactor.RebuildIndex(modelPath)

	assert.NoError(t, err)
	mockModelSvc.AssertExpectations(t)
	mockOutput.AssertExpectations(t)
}

func TestRebuildIndex_Failure(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockOutput := new(out_mocks.ModelRebuildIndex)

	modelPath := "some/model"
	mockModelSvc.On("RebuildIndex", modelPath).Return(errors.New("fs error"))

	interactor := NewRebuildIndexInteractor(mockModelSvc, mockOutput)

	err := interactor.RebuildIndex(modelPath)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to rebuild index")
	mockModelSvc.AssertExpectations(t)
	mockOutput.AssertNotCalled(t, "IndexRebuilt", mock.Anything)
}
