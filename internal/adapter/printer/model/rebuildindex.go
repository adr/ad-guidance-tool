package model

import "fmt"

type RebuildIndexPresenter struct{}

func NewRebuildIndexPresenter() *RebuildIndexPresenter {
	return &RebuildIndexPresenter{}
}

func (p *RebuildIndexPresenter) IndexRebuilt(modelName string) {
	fmt.Printf("Index successfully updated for model: %s\n", modelName)
}
