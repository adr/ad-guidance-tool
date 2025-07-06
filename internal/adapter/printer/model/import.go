package model

import "fmt"

type ImportModelPresenter struct{}

func NewImportPresenter() *ImportModelPresenter {
	return &ImportModelPresenter{}
}

func (p *ImportModelPresenter) Imported(sourcePath, targetPath string, importedDecisions int) error {
	fmt.Printf("Successfully imported model %s with %d decisions to: %s\n", sourcePath, importedDecisions, targetPath)
	return nil
}
