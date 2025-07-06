package model

import "fmt"

type InitModelPresenter struct{}

func NewInitPresenter() *InitModelPresenter {
	return &InitModelPresenter{}
}

func (p *InitModelPresenter) Initialized(modelPath string) {
	fmt.Printf("Successfully created model directory: %s\n", modelPath)
}
