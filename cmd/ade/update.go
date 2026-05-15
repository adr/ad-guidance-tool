package ade

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/adr/ad-guidance-tool/internal/ade/domain"
	"github.com/adr/ad-guidance-tool/internal/ade/pluginstore"
)

var updateCmd = &cobra.Command{
	Use:   "update <name>",
	Short: "Re-fetch the latest release for a remotely installed plugin.",
	Long: `Download the latest GitHub release for a plugin that was previously installed
with 'ade install github.com/owner/repo' and overwrite the existing binary.

  ade plugin update filecheck

Plugins installed with --path (local mode) cannot be updated this way because
no remote source was recorded. Re-run 'ade install <name> --path <new-path>'
to replace a locally installed plugin.

Authentication

If the GITHUB_TOKEN environment variable is set, it is forwarded as a Bearer
token on every GitHub API and download request. This is required for private
repositories and is ignored for public ones.`,
	Args: cobra.ExactArgs(1),
	Run:  updateCommand,
}

func init() {
	pluginCmd.AddCommand(updateCmd)
}

func updateCommand(cmd *cobra.Command, args []string) {
	name := args[0]

	plugins, sources, err := pluginstore.ReadGlobalConfig()
	domain.CheckFatalError(err, "reading global config")

	binaryPath, registered := plugins[name]
	if !registered {
		fmt.Fprintf(os.Stderr, "plugin %q is not registered; run 'ade install' first\n", name)
		os.Exit(1)
	}

	source := sources[name]
	if source == "" {
		fmt.Fprintf(os.Stderr,
			"plugin %q was installed locally (no remote source recorded); "+
				"use 'ade install %s --path <new-path>' to update it\n", name, name)
		os.Exit(1)
	}

	if err := pluginstore.FetchRelease(source, binaryPath); err != nil {
		domain.CheckFatalError(err, "fetching release")
	}
	if err := pluginstore.SetExecutable(binaryPath); err != nil {
		domain.CheckFatalError(err, "setting executable permission")
	}

	fmt.Printf("updated plugin %q at %s\n", name, binaryPath)
}
