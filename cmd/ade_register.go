package cmd

import (
	adecmd "github.com/phi42/ad-enforcement-tool/enforce"
)

func init() {
	rootCmd.AddCommand(adecmd.NewEnforceCommand())
}
