package cmd

import (
	cmd "adg/internal/adapter/command/model"
	print "adg/internal/adapter/printer/model"
	interactor "adg/internal/application/interactor/model"
)

func init() {
	rootCmd.AddCommand(
		cmd.NewCopyCommand(interactor.NewCopyModelInteractor(modelSvc, decisionSvc, print.NewCopyPresenter()), configSvc),
		cmd.NewImportCommand(interactor.NewImportModelInteractor(modelSvc, decisionSvc, print.NewImportPresenter()), configSvc),
		cmd.NewInitCommand(interactor.NewInitModelInteractor(modelSvc, print.NewInitPresenter())),
		cmd.NewMergeModelsCommand(interactor.NewMergeModelsInteractor(modelSvc, decisionSvc, print.NewMergePresenter())),
		cmd.NewRebuildIndexCommand(interactor.NewRebuildIndexInteractor(modelSvc, print.NewRebuildIndexPresenter()), configSvc),
		cmd.NewValidateCommand(interactor.NewModelValidateInteractor(modelSvc, print.NewModelValidatePresenter()), configSvc),
	)
}
