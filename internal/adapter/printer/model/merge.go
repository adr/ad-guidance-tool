package model

import "fmt"

type MergeModelsPresenter struct{}

func NewMergePresenter() *MergeModelsPresenter {
	return &MergeModelsPresenter{}
}

func (p *MergeModelsPresenter) Merged(modelAPath, modelBPath, targetPath string, mergedDecisions int) error {
	fmt.Printf("Successfully merged %d decisions from models %s and %s to new directory: %s\n", mergedDecisions, modelAPath, modelBPath, targetPath)
	return nil
}
