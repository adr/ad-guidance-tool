package cmd

import (
	cmd "github.com/adr/ad-guidance-tool/internal/adapter/command/decision"
	print "github.com/adr/ad-guidance-tool/internal/adapter/printer/decision"
	interactor "github.com/adr/ad-guidance-tool/internal/application/interactor/decision"
	adecmd "github.com/phi42/ad-enforcement-tool/cmd"
)

func init() {
	enforceCmd := adecmd.NewEnforceCommand()
	enforceCmd.AddCommand(
		cmd.NewRuleCommand(interactor.NewRuleInteractor(decisionSvc, print.NewRulePresenter()), configSvc),
	)
	rootCmd.AddCommand(enforceCmd)
}
