package cmd

import (
	cmd "adg/internal/adapter/command/config"
)

func init() {
	rootCmd.AddCommand(
		cmd.NewResetCommand(configSvc),
		cmd.NewSetCommand(configSvc),
	)
}
