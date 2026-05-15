package ade

import (
	"fmt"
	"path/filepath"
	"strings"

	compileapp "github.com/adr/ad-guidance-tool/internal/ade/application/compile"
	"github.com/adr/ad-guidance-tool/internal/ade/application/shared"
	"github.com/adr/ad-guidance-tool/internal/ade/domain"
	"github.com/spf13/cobra"
)

const (
	FLAG_COMPILE_INPUT       = "input"
	FLAG_COMPILE_INPUT_SHORT = "i"
	FLAG_COMPILE_INPUT_USAGE = "input path containing decision (required)"

	FLAG_COMPILE_PLUGIN       = "plugin"
	FLAG_COMPILE_PLUGIN_SHORT = "p"
	FLAG_COMPILE_PLUGIN_USAGE = "plugin name or path (falls back to defaults.compile.plugin in config)"

	FLAG_COMPILE_OUTPUT       = "output"
	FLAG_COMPILE_OUTPUT_SHORT = "o"
	FLAG_COMPILE_OUTPUT_USAGE = "output directory for compiled artifacts (falls back to defaults.compile.output in config)"
)

var enforceCompileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile rules into an executable test artifact using the specified plugin.",
	Long: `Compile rules from an ADR rule file into an executable test artifact.

The plugin generates test code (e.g. a Go test file) in the output directory.
Run the generated tests separately to validate the rules.

Examples:
  ade enforce compile -i docs/0001.rule -p arch-go -o ./internal
  ade enforce compile -i docs/ -p arch-go -o ./internal`,
	Run: enforceCompileCommand,
}

func init() {
	enforceCmd.AddCommand(enforceCompileCmd)

	enforceCompileCmd.Flags().StringP(FLAG_COMPILE_INPUT, FLAG_COMPILE_INPUT_SHORT, "", FLAG_COMPILE_INPUT_USAGE)
	enforceCompileCmd.MarkFlagRequired(FLAG_COMPILE_INPUT)

	enforceCompileCmd.Flags().StringP(FLAG_COMPILE_PLUGIN, FLAG_COMPILE_PLUGIN_SHORT, "", FLAG_COMPILE_PLUGIN_USAGE)

	enforceCompileCmd.Flags().StringP(FLAG_COMPILE_OUTPUT, FLAG_COMPILE_OUTPUT_SHORT, "", FLAG_COMPILE_OUTPUT_USAGE)
}

func enforceCompileCommand(cmd *cobra.Command, args []string) {
	input, err := cmd.Flags().GetString(FLAG_COMPILE_INPUT)
	domain.CheckFatalError(err, "reading input flag")

	plugin, err := cmd.Flags().GetString(FLAG_COMPILE_PLUGIN)
	domain.CheckFatalError(err, "reading plugin flag")
	if plugin == "" {
		plugin = adeViper.GetString(domain.CONFIG_DEFAULT_COMPILE_PLUGIN)
	}
	if plugin == "" {
		domain.CheckFatalError(fmt.Errorf("--plugin is required (pass as flag or set %s in config)", domain.CONFIG_DEFAULT_COMPILE_PLUGIN), "resolving plugin")
	}
	if !filepath.IsAbs(plugin) && filepath.Dir(plugin) == "." {
		if configPath := adeViper.GetString(domain.CONFIG_PLUGIN_KEY_PREFIX + plugin); configPath != "" {
			plugin = configPath
		}
	}

	output, err := cmd.Flags().GetString(FLAG_COMPILE_OUTPUT)
	domain.CheckFatalError(err, "reading output flag")
	if strings.TrimSpace(output) == "" {
		output = adeViper.GetString(domain.CONFIG_DEFAULT_COMPILE_OUTPUT)
	}
	if strings.TrimSpace(output) == "" {
		domain.CheckFatalError(fmt.Errorf("--output is required (pass as flag or set %s in config)", domain.CONFIG_DEFAULT_COMPILE_OUTPUT), "resolving output")
	}

	domain.CheckFatalError(shared.ValidatePluginMode(plugin, "compile"), "validating plugin mode")

	compileapp.Compile(compileapp.CompileInfo{
		InputFile:  input,
		PluginName: plugin,
		OutputDir:  output,
	})
}
