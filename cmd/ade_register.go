package cmd

import (
	adecmd "github.com/adr/ad-guidance-tool/cmd/ade"
)

func init() {
	rootCmd.AddCommand(adecmd.NewEnforceCommand())
}
