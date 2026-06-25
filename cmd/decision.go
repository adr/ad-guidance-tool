package cmd

import (
	cmd "github.com/adr/ad-guidance-tool/internal/adapter/command/decision"
	print "github.com/adr/ad-guidance-tool/internal/adapter/printer/decision"
	interactor "github.com/adr/ad-guidance-tool/internal/application/interactor/decision"
)

func init() {
	rootCmd.AddCommand(
		cmd.NewAddCommand(interactor.NewAddDecisionsInteractor(modelSvc, decisionSvc, print.NewAddPresenter()), configSvc),
		cmd.NewCommentCommand(interactor.NewCommentDecisionInteractor(decisionSvc, print.NewCommentPresenter()), configSvc),
		cmd.NewDecideCommand(interactor.NewDecideInteractor(decisionSvc, print.NewDecidePresenter()), configSvc),
		cmd.NewEditCommand(interactor.NewEditDecisionInteractor(decisionSvc, print.NewEditPresenter()), configSvc),
		cmd.NewLinkCommand(interactor.NewLinkDecisionsInteractor(decisionSvc, print.NewLinkPresenter()), configSvc),
		cmd.NewListCommand(interactor.NewListDecisionsInteractor(decisionSvc, print.NewListPresenter()), configSvc),
		cmd.NewPrintCommand(interactor.NewPrintDecisionsInteractor(decisionSvc, print.NewPrintPresenter(configSvc)), configSvc),
		cmd.NewReviseCommand(interactor.NewReviseDecisionInteractor(decisionSvc, print.NewRevisePresenter()), configSvc),
		cmd.NewTagCommand(interactor.NewTagDecisionInteractor(decisionSvc, print.NewTagPresenter()), configSvc),
	)
}
