package ade

import (
	"fmt"
	"path/filepath"

	"github.com/adr/ad-guidance-tool/internal/ade/application/shared"
	verifyapp "github.com/adr/ad-guidance-tool/internal/ade/application/verify"
	"github.com/adr/ad-guidance-tool/internal/ade/domain"
	"github.com/spf13/cobra"
)

const (
	FLAG_VERIFY_INPUT       = "input"
	FLAG_VERIFY_INPUT_SHORT = "i"
	FLAG_VERIFY_INPUT_USAGE = "input path containing decision (required)"

	FLAG_VERIFY_PLUGIN       = "plugin"
	FLAG_VERIFY_PLUGIN_SHORT = "p"
	FLAG_VERIFY_PLUGIN_USAGE = "plugin name or path (falls back to defaults.verify.plugin in config)"

	FLAG_VERIFY_ROOT       = "root"
	FLAG_VERIFY_ROOT_SHORT = "r"
	FLAG_VERIFY_ROOT_USAGE = "root directory that file paths in rules are resolved against (default: current working directory)"
)

var enforceVerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify rules immediately against the target using the specified plugin.",
	Long: `Execute rules from an ADR rule file directly against the target and report results.

Unlike 'enforce compile', which generates test code for later execution, 'enforce verify'
runs assertions immediately using the plugin and exits non-zero if any rule is violated.

Examples:
  adg enforce verify -i docs/0003.rule -p filecheck
  adg enforce verify -i docs/ -p filecheck -r ./src`,
	Run: enforceVerifyCommand,
}

func init() {
	enforceCmd.AddCommand(enforceVerifyCmd)

	enforceVerifyCmd.Flags().StringP(FLAG_VERIFY_INPUT, FLAG_VERIFY_INPUT_SHORT, "", FLAG_VERIFY_INPUT_USAGE)
	enforceVerifyCmd.MarkFlagRequired(FLAG_VERIFY_INPUT)

	enforceVerifyCmd.Flags().StringP(FLAG_VERIFY_PLUGIN, FLAG_VERIFY_PLUGIN_SHORT, "", FLAG_VERIFY_PLUGIN_USAGE)

	enforceVerifyCmd.Flags().StringP(FLAG_VERIFY_ROOT, FLAG_VERIFY_ROOT_SHORT, "", FLAG_VERIFY_ROOT_USAGE)
}

func enforceVerifyCommand(cmd *cobra.Command, args []string) {
	input, err := cmd.Flags().GetString(FLAG_VERIFY_INPUT)
	domain.CheckFatalError(err, "reading input flag")

	plugin, err := cmd.Flags().GetString(FLAG_VERIFY_PLUGIN)
	domain.CheckFatalError(err, "reading plugin flag")
	if plugin == "" {
		plugin = adeViper.GetString(domain.CONFIG_DEFAULT_VERIFY_PLUGIN)
	}
	if plugin == "" {
		domain.CheckFatalError(fmt.Errorf("--plugin is required (pass as flag or set %s in config)", domain.CONFIG_DEFAULT_VERIFY_PLUGIN), "resolving plugin")
	}
	if !filepath.IsAbs(plugin) && filepath.Dir(plugin) == "." {
		if configPath := adeViper.GetString(domain.CONFIG_PLUGIN_KEY_PREFIX + plugin); configPath != "" {
			plugin = configPath
		}
	}

	root, err := cmd.Flags().GetString(FLAG_VERIFY_ROOT)
	domain.CheckFatalError(err, "reading root flag")

	domain.CheckFatalError(shared.ValidatePluginMode(plugin, "verify"), "validating plugin mode")

	ruleFiles, err := collectRuleFilePaths(input)
	domain.CheckFatalError(err, "resolving input path")

	for _, f := range ruleFiles {
		verifyapp.Verify(verifyapp.VerifyInfo{
			InputFile:  f,
			PluginName: plugin,
			RootDir:    root,
		})
	}
}
