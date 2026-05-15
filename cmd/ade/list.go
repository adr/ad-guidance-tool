package ade

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/adr/ad-guidance-tool/internal/ade/domain"
	"github.com/adr/ad-guidance-tool/internal/ade/pluginstore"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all globally registered plugins.",
	Long: `Print all plugins registered in the global config file with their paths and status.

The global config is stored in the platform config directory:
  Linux/macOS: $XDG_CONFIG_HOME/ade/ade.yaml (default: ~/.config/ade/ade.yaml)
  Windows:     %APPDATA%\ade\ade.yaml

The STATUS column shows:
  ok       — binary exists at the registered path
  missing  — binary not found (may need to reinstall)

The SOURCE column shows the GitHub module URL for remotely installed plugins,
or "(local)" for plugins installed with --path.`,
	Args: cobra.NoArgs,
	Run:  listCommand,
}

func init() {
	pluginCmd.AddCommand(listCmd)
}

func listCommand(cmd *cobra.Command, args []string) {
	plugins, sources, err := pluginstore.ReadGlobalConfig()
	domain.CheckFatalError(err, "reading global config")

	if len(plugins) == 0 {
		fmt.Println("no plugins registered; use 'ade install' to add one")
		return
	}

	// Sort names for deterministic output.
	names := make([]string, 0, len(plugins))
	for name := range plugins {
		names = append(names, name)
	}
	sort.Strings(names)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "PLUGIN\tPATH\tSTATUS\tSOURCE")
	for _, name := range names {
		path := plugins[name]
		status := "ok"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			status = "missing"
		}
		source := sources[name]
		if source == "" {
			source = "(local)"
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", name, path, status, source)
	}
	w.Flush()
}
