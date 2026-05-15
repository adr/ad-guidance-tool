package pluginstore

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

const pluginSourcesKeyPrefix = "plugin_sources."

// xdgConfigHome returns the XDG config home for the current platform.
// It respects $XDG_CONFIG_HOME on all platforms.
// On Windows it falls back to %APPDATA%, then $HOME/AppData/Roaming.
// On other platforms it falls back to $HOME/.config.
func xdgConfigHome() (string, error) {
	if dir := os.Getenv("XDG_CONFIG_HOME"); dir != "" {
		return dir, nil
	}
	if runtime.GOOS == "windows" {
		if dir := os.Getenv("APPDATA"); dir != "" {
			return dir, nil
		}
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to determine home directory: %w", err)
	}
	if runtime.GOOS == "windows" {
		return filepath.Join(home, "AppData", "Roaming"), nil
	}
	return filepath.Join(home, ".config"), nil
}

// xdgDataHome returns the XDG data home for the current platform.
// It respects $XDG_DATA_HOME on all platforms.
// On Windows it falls back to %APPDATA%, then $HOME/AppData/Roaming.
// On other platforms it falls back to $HOME/.local/share.
func xdgDataHome() (string, error) {
	if dir := os.Getenv("XDG_DATA_HOME"); dir != "" {
		return dir, nil
	}
	if runtime.GOOS == "windows" {
		if dir := os.Getenv("APPDATA"); dir != "" {
			return dir, nil
		}
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to determine home directory: %w", err)
	}
	if runtime.GOOS == "windows" {
		return filepath.Join(home, "AppData", "Roaming"), nil
	}
	return filepath.Join(home, ".local", "share"), nil
}

// GlobalPluginDir returns the path to the platform data directory for ADE plugins.
// Linux/macOS: $XDG_DATA_HOME/ade/plugins (default: ~/.local/share/ade/plugins)
// Windows:     %APPDATA%\ade\plugins
func GlobalPluginDir() (string, error) {
	base, err := xdgDataHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "ade", "plugins"), nil
}

// GlobalConfigPath returns the path to the platform config file for ADE.
// Linux/macOS: $XDG_CONFIG_HOME/ade/ade.yaml (default: ~/.config/ade/ade.yaml)
// Windows:     %APPDATA%\ade\ade.yaml
func GlobalConfigPath() (string, error) {
	base, err := xdgConfigHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "ade", "ade.yaml"), nil
}

// PluginBinaryName returns the OS-appropriate filename for a plugin binary.
// On Windows it appends ".exe" if the name does not already have that extension.
func PluginBinaryName(name string) string {
	if runtime.GOOS == "windows" && !strings.HasSuffix(name, ".exe") {
		return name + ".exe"
	}
	return name
}

// CopyBinary copies the file at src to dst, creating parent directories if needed.
func CopyBinary(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return fmt.Errorf("creating plugin directory: %w", err)
	}
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("opening source file: %w", err)
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0750)
	if err != nil {
		return fmt.Errorf("creating destination file: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copying binary: %w", err)
	}
	return nil
}

// SetExecutable ensures the file at path is executable.
// On Windows this is a no-op; executability is determined by file extension.
func SetExecutable(path string) error {
	if runtime.GOOS == "windows" {
		return nil
	}
	return os.Chmod(path, 0750)
}

// NormaliseModuleURL strips any https:// or http:// prefix from a module URL.
func NormaliseModuleURL(moduleURL string) string {
	moduleURL = strings.TrimPrefix(moduleURL, "https://")
	moduleURL = strings.TrimPrefix(moduleURL, "http://")
	return moduleURL
}

// ParseModuleURL splits a "github.com/<owner>/<repo>" module URL into its parts.
// An optional "https://" prefix is stripped before parsing.
func ParseModuleURL(moduleURL string) (owner, repo string, err error) {
	moduleURL = NormaliseModuleURL(moduleURL)
	parts := strings.SplitN(moduleURL, "/", 3)
	if len(parts) != 3 || parts[0] != "github.com" {
		return "", "", fmt.Errorf("invalid module URL %q: expected github.com/<owner>/<repo>", moduleURL)
	}
	return parts[1], parts[2], nil
}

// FetchRelease downloads the binary for the current OS and architecture from the
// latest GitHub release of moduleURL (e.g. "github.com/owner/repo") into dst.
func FetchRelease(moduleURL, dst string) error {
	owner, repo, err := ParseModuleURL(moduleURL)
	if err != nil {
		return err
	}

	release, err := fetchLatestRelease(owner, repo)
	if err != nil {
		return err
	}

	asset, err := matchAsset(release)
	if err != nil {
		return err
	}

	return downloadAsset(owner, repo, asset.ID, dst)
}

// UpdateGlobalConfig writes or updates the plugin entry in ~/.ade.yaml.
// source is the module URL for remotely installed plugins; pass an empty string
// for locally installed plugins.
func UpdateGlobalConfig(name, binaryPath, source string) error {
	cfgPath, err := GlobalConfigPath()
	if err != nil {
		return err
	}

	cfg, err := readConfigFile(cfgPath)
	if err != nil {
		return err
	}

	plugins, _ := cfg["plugins"].(map[string]interface{})
	if plugins == nil {
		plugins = make(map[string]interface{})
		cfg["plugins"] = plugins
	}
	plugins[name] = binaryPath

	if source != "" {
		sources, _ := cfg["plugin_sources"].(map[string]interface{})
		if sources == nil {
			sources = make(map[string]interface{})
			cfg["plugin_sources"] = sources
		}
		sources[name] = source
	}

	return writeConfigFile(cfgPath, cfg)
}

// RemoveFromGlobalConfig deletes the plugin and its source entry from ~/.ade.yaml.
// It is not an error if the plugin is not registered.
func RemoveFromGlobalConfig(name string) error {
	cfgPath, err := GlobalConfigPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("reading global config: %w", err)
	}

	var cfg map[string]interface{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("parsing global config: %w", err)
	}
	if cfg == nil {
		return nil
	}

	if plugins, ok := cfg["plugins"].(map[string]interface{}); ok {
		delete(plugins, name)
	}
	if sources, ok := cfg["plugin_sources"].(map[string]interface{}); ok {
		delete(sources, name)
	}

	return writeConfigFile(cfgPath, cfg)
}

// ReadGlobalConfig returns the plugins map and plugin_sources map from ~/.ade.yaml.
func ReadGlobalConfig() (plugins map[string]string, sources map[string]string, err error) {
	plugins = make(map[string]string)
	sources = make(map[string]string)

	cfgPath, err := GlobalConfigPath()
	if err != nil {
		return nil, nil, err
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		if os.IsNotExist(err) {
			return plugins, sources, nil
		}
		return nil, nil, fmt.Errorf("reading global config: %w", err)
	}

	var cfg map[string]interface{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, nil, fmt.Errorf("parsing global config: %w", err)
	}

	if raw, ok := cfg["plugins"].(map[string]interface{}); ok {
		for k, v := range raw {
			if s, ok := v.(string); ok {
				plugins[k] = s
			}
		}
	}
	if raw, ok := cfg["plugin_sources"].(map[string]interface{}); ok {
		for k, v := range raw {
			if s, ok := v.(string); ok {
				sources[k] = s
			}
		}
	}
	return plugins, sources, nil
}

// — internal helpers —

type githubRelease struct {
	Assets []githubAsset `json:"assets"`
}

type githubAsset struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func fetchLatestRelease(owner, repo string) (*githubRelease, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("building GitHub API request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "ade-tool")
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("calling GitHub API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned %s for %s/%s", resp.Status, owner, repo)
	}

	var release githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("decoding GitHub API response: %w", err)
	}
	return &release, nil
}

func matchAsset(release *githubRelease) (*githubAsset, error) {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	for i := range release.Assets {
		name := strings.ToLower(release.Assets[i].Name)
		if strings.Contains(name, goos) && strings.Contains(name, goarch) {
			return &release.Assets[i], nil
		}
	}
	return nil, fmt.Errorf(
		"no release asset found for %s/%s; asset filenames must include OS and architecture (e.g., my-plugin-%s-%s)",
		goos, goarch, goos, goarch,
	)
}

// downloadAsset downloads a release asset via the GitHub API endpoint.
// Using the API URL (rather than browser_download_url) ensures that the
// Authorization header is only sent to api.github.com; the subsequent CDN
// redirect receives a pre-signed URL and does not require authentication.
// This is the correct approach for private repository assets.
func downloadAsset(owner, repo string, assetID int64, dst string) error {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/assets/%d", owner, repo, assetID)

	if err := os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return fmt.Errorf("creating plugin directory: %w", err)
	}

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return fmt.Errorf("building download request: %w", err)
	}
	req.Header.Set("Accept", "application/octet-stream")
	req.Header.Set("User-Agent", "ade-tool")
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("downloading: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download returned %s", resp.Status)
	}

	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0750)
	if err != nil {
		return fmt.Errorf("creating destination file %s: %w", dst, err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("writing downloaded binary: %w", err)
	}
	return nil
}

func readConfigFile(cfgPath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]interface{}), nil
		}
		return nil, fmt.Errorf("reading global config: %w", err)
	}
	var cfg map[string]interface{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing global config: %w", err)
	}
	if cfg == nil {
		return make(map[string]interface{}), nil
	}
	return cfg, nil
}

func writeConfigFile(cfgPath string, cfg map[string]interface{}) error {
	if err := os.MkdirAll(filepath.Dir(cfgPath), 0750); err != nil {
		return fmt.Errorf("creating config directory: %w", err)
	}
	updated, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("serializing config: %w", err)
	}
	return os.WriteFile(cfgPath, updated, 0600)
}
