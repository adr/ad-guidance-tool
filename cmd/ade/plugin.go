package ade

import "github.com/spf13/cobra"

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Manage enforcement plugins.",
	Long:  `Install, uninstall, update, and list enforcement plugins.`,
}

func init() {
	enforceCmd.AddCommand(pluginCmd)
}
