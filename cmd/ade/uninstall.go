package ade

import (
	"fmt"
	"os"

	"github.com/adr/ad-guidance-tool/internal/ade/domain"
	"github.com/adr/ad-guidance-tool/internal/ade/pluginstore"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall <name>",
	Short: "Remove a plugin from the global plugin directory.",
	Long: `Remove a globally installed plugin binary and its entry in the global config file.

  adg enforce plugin uninstall filecheck`,

	Args: cobra.ExactArgs(1),
	Run:  uninstallCommand,
}

func init() {
	pluginCmd.AddCommand(uninstallCmd)
}

func uninstallCommand(cmd *cobra.Command, args []string) {
	name := args[0]

	plugins, _, err := pluginstore.ReadGlobalConfig()
	domain.CheckFatalError(err, "reading global config")

	binaryPath, registered := plugins[name]
	if !registered {
		fmt.Fprintf(os.Stderr, "plugin %q is not registered in the global config\n", name)
		os.Exit(1)
	}

	if err := os.Remove(binaryPath); err != nil && !os.IsNotExist(err) {
		domain.CheckFatalError(err, "removing plugin binary")
	}

	if err := pluginstore.RemoveFromGlobalConfig(name); err != nil {
		domain.CheckFatalError(err, "updating global config")
	}

	fmt.Printf("uninstalled plugin %q\n", name)
}
