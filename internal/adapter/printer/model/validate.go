package model

import "fmt"

type ModelValidatePresenter struct{}

func NewModelValidatePresenter() *ModelValidatePresenter {
	return &ModelValidatePresenter{}
}

func (p *ModelValidatePresenter) ModelValidated(modelName string, indexErr, dataErr error) {
	if indexErr == nil {
		fmt.Printf("%s model metadata is valid and index is up to date\n", modelName)

		if dataErr == nil {
			fmt.Printf("%s model file content is valid with correct anchors\n", modelName)
		} else {
			fmt.Printf("%s model file content is invalid: %s\n", modelName, dataErr)
		}
	} else {
		fmt.Printf("%s model metadata is invalid: %s\n", modelName, indexErr)
	}

	// todo: print total number of valid/invalid decisions if available
}
