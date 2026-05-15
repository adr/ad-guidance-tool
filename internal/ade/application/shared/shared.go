package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/exec"

	"github.com/adr/ad-guidance-tool/internal/ade/domain"
	"google.golang.org/protobuf/proto"
)

// pluginInfo is the JSON shape returned by a plugin when invoked with --info.
type pluginInfo struct {
	Modes []string `json:"modes"`
}

// QueryPluginInfo runs the plugin with the --info flag and returns the list
// of invocation modes it advertises (e.g. "compile", "verify").
func QueryPluginInfo(plugin string) ([]string, error) {
	path, err := domain.ResolvePluginPath(plugin)
	if err != nil {
		return nil, err
	}

	var stdout bytes.Buffer
	cmd := exec.Command(path, "--info")
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("querying plugin %q with --info: %w", plugin, err)
	}

	var info pluginInfo
	if err := json.Unmarshal(bytes.TrimSpace(stdout.Bytes()), &info); err != nil {
		return nil, fmt.Errorf("parsing --info response from plugin %q: %w", plugin, err)
	}
	return info.Modes, nil
}

// ValidatePluginMode ensures the plugin advertises the requested mode. It
// returns a descriptive error if the plugin does not support that mode.
func ValidatePluginMode(plugin, requestedMode string) error {
	modes, err := QueryPluginInfo(plugin)
	if err != nil {
		return err
	}
	for _, m := range modes {
		if m == requestedMode {
			return nil
		}
	}
	return fmt.Errorf("plugin %q supports modes %v and cannot be used with \"enforce %s\"", plugin, modes, requestedMode)
}

// CompileSpec reads and parses a DSL rule file, returning the SpecIR.
func CompileSpec(path string) (*domain.SpecIR, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	ir, err := domain.ParseDSL(string(data))
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	slog.Debug("loaded and validated spec", "file", path)
	return ir, nil
}

// RunPlugin serialises ir as protobuf, pipes it into the plugin process and
// streams stdout/stderr back to the caller's console.
func RunPlugin(plugin string, ir *domain.SpecIR) error {
	path, err := domain.ResolvePluginPath(plugin)
	if err != nil {
		return err
	}

	payload, err := proto.Marshal(ir)
	if err != nil {
		return fmt.Errorf("unable to marshal SpecIR to protobuf: %s", err.Error())
	}

	pluginCmd := exec.Command(path)
	pluginCmd.Stdin = bytes.NewReader(payload)
	pluginCmd.Stdout = os.Stdout
	pluginCmd.Stderr = os.Stderr

	if err := pluginCmd.Run(); err != nil {
		return fmt.Errorf("plugin failed: %s", err.Error())
	}
	return nil
}
