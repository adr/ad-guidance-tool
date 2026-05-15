package domain

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

// Constants for app and configuration variables.
const (
	CONFIG_FILE_NAME         string = ".ade"
	CONFIG_FILE_EXT          string = "yaml"
	CONFIG_PLUGIN_KEY_PREFIX string = "plugins."

	CONFIG_DEFAULT_COMPILE_PLUGIN string = "defaults.compile.plugin"
	CONFIG_DEFAULT_COMPILE_OUTPUT string = "defaults.compile.output"
	CONFIG_DEFAULT_VERIFY_PLUGIN  string = "defaults.verify.plugin"
)

// KnownDefaults lists all configuration keys that can be set via 'ade config set'.
var KnownDefaults = []string{
	CONFIG_DEFAULT_COMPILE_PLUGIN,
	CONFIG_DEFAULT_COMPILE_OUTPUT,
	CONFIG_DEFAULT_VERIFY_PLUGIN,
}

// ResolvePluginPath resolves a plugin name or path to an executable path.
// If the input looks like a path (absolute or contains separators) it is used directly.
// Otherwise the current working directory is searched.
func ResolvePluginPath(pluginNameOrPath string) (string, error) {
	// check if input looks like a path (contains path separators)
	if filepath.IsAbs(pluginNameOrPath) || filepath.Dir(pluginNameOrPath) != "." {
		cleanPath := filepath.Clean(pluginNameOrPath)
		if _, err := os.Stat(cleanPath); err != nil {
			if os.IsNotExist(err) {
				return "", fmt.Errorf("plugin path does not exist: %s", cleanPath)
			}
			return "", fmt.Errorf("unable to access plugin path %s: %w", cleanPath, err)
		}
		return cleanPath, nil
	}

	// search in current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("unable to get current working directory: %w", err)
	}

	if path := tryFindExecutable(cwd, pluginNameOrPath); path != "" {
		return path, nil
	}

	cwdPath := filepath.Join(cwd, pluginNameOrPath)
	return "", fmt.Errorf("plugin not found: tried '%s'", cwdPath)
}

// tryFindExecutable attempts to find an executable in the given directory.
func tryFindExecutable(dir, name string) string {
	path := filepath.Join(dir, name)
	if _, err := os.Stat(path); err == nil {
		return path
	}

	if filepath.Ext(name) == "" {
		pathWithExe := filepath.Join(dir, name+".exe")
		if _, err := os.Stat(pathWithExe); err == nil {
			return pathWithExe
		}
	}

	return ""
}

// Calls [slog.Error] and exits with code 1 if error is not nil.
func CheckFatalError(err error, activity string) {
	if err != nil {
		slog.Error("fatal error", "activity", activity, "error", err.Error())
		os.Exit(1)
	}
}
