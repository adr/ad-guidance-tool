package model

import (
	out_mocks "adg/mocks/outputport"
	svc_mocks "adg/mocks/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRebuildIndex_NoOp verifies the PR 1b shim behavior: index.yaml is gone,
// so the interactor simply reports success and never touches the model service.
// The full removal of `adg rebuild` is deferred to PR 2.
func TestRebuildIndex_NoOp(t *testing.T) {
	mockModelSvc := new(svc_mocks.ModelService)
	mockOutput := new(out_mocks.ModelRebuildIndex)

	modelPath := "some/model"

	mockOutput.On("IndexRebuilt", modelPath).Return()

	interactor := NewRebuildIndexInteractor(mockModelSvc, mockOutput)

	err := interactor.RebuildIndex(modelPath)

	assert.NoError(t, err)
	mockOutput.AssertExpectations(t)
	mockModelSvc.AssertNotCalled(t, "RebuildIndex")
}
