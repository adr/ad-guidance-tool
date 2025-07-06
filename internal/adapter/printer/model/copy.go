package model

import (
	"fmt"
)

type CopyModelPresenter struct{}

func NewCopyPresenter() *CopyModelPresenter {
	return &CopyModelPresenter{}
}

func (p *CopyModelPresenter) Copied(source, target string, copiedDecisions int) {
	fmt.Printf("Successfully copied %d decisions from model %s to new model %s\n", copiedDecisions, source, target)
}
