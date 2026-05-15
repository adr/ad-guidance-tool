package ade

import "github.com/spf13/cobra"

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Manage ADE plugins.",
	Long:  `Install, uninstall, update, and list ADE plugins.`,
}

func init() {
	enforceCmd.AddCommand(pluginCmd)
}
