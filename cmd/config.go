package cmd

import (
	cmd "github.com/adr/ad-guidance-tool/internal/adapter/command/config"
)

func init() {
	rootCmd.AddCommand(
		cmd.NewResetCommand(configSvc),
		cmd.NewSetCommand(configSvc),
	)
}
