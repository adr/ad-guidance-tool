package ade

import (
	"fmt"
	"path/filepath"

	"github.com/adr/ad-guidance-tool/internal/ade/domain"
	"github.com/adr/ad-guidance-tool/internal/ade/pluginstore"
	"github.com/spf13/cobra"
)

const (
	FLAG_INSTALL_PATH       = "path"
	FLAG_INSTALL_PATH_SHORT = "p"
	FLAG_INSTALL_PATH_USAGE = "path to an existing local binary to register"

	FLAG_INSTALL_REPO       = "repo"
	FLAG_INSTALL_REPO_SHORT = "r"
	FLAG_INSTALL_REPO_USAGE = "GitHub module path to download from (e.g. github.com/owner/repo)"
)

var installCmd = &cobra.Command{
	Use:   "install <name>",
	Short: "Install a plugin into the global plugin directory.",
	Long: `Install a plugin and register it in the global config.

The global config is stored in the platform config directory:
  Linux/macOS: $XDG_CONFIG_HOME/ade/ade.yaml (default: ~/.config/ade/ade.yaml)
  Windows:     %APPDATA%\ade\ade.yaml

The <name> argument is always the name under which the plugin is registered.
Exactly one of --path or --repo must be provided.

Local mode: register an existing binary you have already built or downloaded:

  adg enforce plugin install filecheck --path /tmp/downloads/filecheck

Remote mode: download directly from a GitHub release:

  adg enforce plugin install arch-go --repo github.com/phi42/adplugin-arch-go

The release must contain assets whose filenames include the OS and architecture,
for example:

  my-plugin-linux-amd64
  my-plugin-darwin-arm64
  my-plugin-windows-amd64.exe

The binary is copied to the platform data directory:
  Linux/macOS: $XDG_DATA_HOME/ade/plugins/<name> (default: ~/.local/share/ade/plugins/<name>)
  Windows:     %APPDATA%\ade\plugins\<name>

An entry is added under plugins.<name> in the global config file.

Authentication

If the GITHUB_TOKEN environment variable is set, it is forwarded as a Bearer
token on every GitHub API and download request. This is required for private
repositories. For public repositories the variable is not needed. On GitHub
Actions runners GITHUB_TOKEN is set automatically; on a developer machine you
must set it manually (e.g. in your shell profile).`,
	Args: cobra.ExactArgs(1),
	Run:  installCommand,
}

func init() {
	pluginCmd.AddCommand(installCmd)
	installCmd.Flags().StringP(FLAG_INSTALL_PATH, FLAG_INSTALL_PATH_SHORT, "", FLAG_INSTALL_PATH_USAGE)
	installCmd.Flags().StringP(FLAG_INSTALL_REPO, FLAG_INSTALL_REPO_SHORT, "", FLAG_INSTALL_REPO_USAGE)
	installCmd.MarkFlagsMutuallyExclusive(FLAG_INSTALL_PATH, FLAG_INSTALL_REPO)
}

func installCommand(cmd *cobra.Command, args []string) {
	name := args[0]
	localPath, err := cmd.Flags().GetString(FLAG_INSTALL_PATH)
	domain.CheckFatalError(err, "reading path flag")
	repoURL, err := cmd.Flags().GetString(FLAG_INSTALL_REPO)
	domain.CheckFatalError(err, "reading repo flag")

	if localPath == "" && repoURL == "" {
		domain.CheckFatalError(fmt.Errorf("one of --path or --repo is required"), "validating flags")
	}

	pluginDir, err := pluginstore.GlobalPluginDir()
	domain.CheckFatalError(err, "resolving plugin directory")

	binaryName := pluginstore.PluginBinaryName(name)
	dst := filepath.Join(pluginDir, binaryName)

	var source string

	if localPath != "" {
		// Local mode: copy the binary from --path.
		if err := pluginstore.CopyBinary(localPath, dst); err != nil {
			domain.CheckFatalError(err, "copying binary")
		}
		if err := pluginstore.SetExecutable(dst); err != nil {
			domain.CheckFatalError(err, "setting executable permission")
		}
		source = ""
	} else {
		// Remote mode: download from GitHub release at --repo.
		repoURL = pluginstore.NormaliseModuleURL(repoURL)
		if err := pluginstore.FetchRelease(repoURL, dst); err != nil {
			domain.CheckFatalError(err, "fetching release")
		}
		if err := pluginstore.SetExecutable(dst); err != nil {
			domain.CheckFatalError(err, "setting executable permission")
		}
		source = repoURL
	}

	if err := pluginstore.UpdateGlobalConfig(name, dst, source); err != nil {
		domain.CheckFatalError(err, "updating global config")
	}
	fmt.Printf("installed plugin %q to %s\n", name, dst)
}
