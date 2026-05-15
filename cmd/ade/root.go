package ade

import (
	"log/slog"
	"os"

	"github.com/adr/ad-guidance-tool/internal/ade/domain"
	"github.com/adr/ad-guidance-tool/internal/ade/pluginstore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var adeViper = viper.New()

var cfgFile string
var debugFlag bool
var quietFlag bool
var noWarnFlag bool
var configFileUsed string

var enforceCmd = &cobra.Command{
	Use:   "enforce",
	Short: "Architectural Decision Enforcement (ADE) commands.",
	Long:  `Commands for enforcing architectural decisions: validate rule files, compile rules into test artifacts, verify rules against the codebase, and manage plugins and configuration.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initADEConfig()
		setupLogger(debugFlag, quietFlag, noWarnFlag)
		if configFileUsed != "" {
			slog.Debug("using config file", "path", configFileUsed)
		}
	},
}

// NewEnforceCommand returns the 'enforce' subcommand for use in the ADG root.
func NewEnforceCommand() *cobra.Command {
	return enforceCmd
}

func init() {
	enforceCmd.CompletionOptions.HiddenDefaultCmd = true
	enforceCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (overrides default hierarchy: global ade.yaml then .ade.yaml in current directory)")
	enforceCmd.PersistentFlags().BoolVar(&debugFlag, "debug", false, "enable debug output")
	enforceCmd.PersistentFlags().BoolVar(&quietFlag, "quiet", false, "suppress all output except errors")
	enforceCmd.PersistentFlags().BoolVar(&noWarnFlag, "no-warnings", false, "suppress warnings, show info and errors")
}

func initADEConfig() {
	if cfgFile != "" {
		adeViper.SetConfigFile(cfgFile)
		adeViper.AutomaticEnv()
		if err := adeViper.ReadInConfig(); err == nil {
			configFileUsed = adeViper.ConfigFileUsed()
		}
		return
	}

	adeViper.AutomaticEnv()

	// Load the global config from the XDG config directory as the fallback base.
	if globalCfg, err := pluginstore.GlobalConfigPath(); err == nil {
		if _, err := os.Stat(globalCfg); err == nil {
			adeViper.SetConfigFile(globalCfg)
			if err := adeViper.ReadInConfig(); err == nil {
				configFileUsed = adeViper.ConfigFileUsed()
			}
		}
	}

	// Merge the project-level config on top; its values take precedence.
	if cwd, err := os.Getwd(); err == nil {
		pv := viper.New()
		pv.AddConfigPath(cwd)
		pv.SetConfigType(domain.CONFIG_FILE_EXT)
		pv.SetConfigName(domain.CONFIG_FILE_NAME)
		if err := pv.ReadInConfig(); err == nil {
			if mergeErr := adeViper.MergeConfigMap(pv.AllSettings()); mergeErr == nil {
				configFileUsed = pv.ConfigFileUsed()
			}
		}
	}
}

func setupLogger(debug, quiet, noWarn bool) {
	level := slog.LevelInfo
	skipWarn := noWarn
	switch {
	case debug:
		level = slog.LevelDebug
		os.Setenv("ADE_LOG_LEVEL", "debug")
	case quiet:
		level = slog.LevelError
		os.Setenv("ADE_LOG_LEVEL", "quiet")
	case noWarn:
		os.Setenv("ADE_LOG_LEVEL", "no-warnings")
	}
	slog.SetDefault(slog.New(newCLIHandler(os.Stderr, level, skipWarn)))
}
