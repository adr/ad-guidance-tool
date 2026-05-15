package pluginstore

import (
	"fmt"
	"strings"
)

// SetDefault sets a dotted key (e.g. "defaults.compile.plugin") to value in the
// YAML config file at cfgPath. Intermediate map nodes are created as needed.
func SetDefault(cfgPath, dottedKey, value string) error {
	cfg, err := readConfigFile(cfgPath)
	if err != nil {
		return err
	}

	parts := strings.Split(dottedKey, ".")
	setNestedKey(cfg, parts, value)

	return writeConfigFile(cfgPath, cfg)
}

// UnsetDefault removes a dotted key from the YAML config file at cfgPath.
// Empty parent maps are pruned. It is not an error if the key does not exist.
func UnsetDefault(cfgPath, dottedKey string) error {
	cfg, err := readConfigFile(cfgPath)
	if err != nil {
		return err
	}

	parts := strings.Split(dottedKey, ".")
	unsetNestedKey(cfg, parts)

	return writeConfigFile(cfgPath, cfg)
}

// GetDefault reads a single dotted key from the YAML config file at cfgPath.
// Returns the value and true if found, or "" and false if not present.
func GetDefault(cfgPath, dottedKey string) (string, bool, error) {
	cfg, err := readConfigFile(cfgPath)
	if err != nil {
		return "", false, err
	}

	parts := strings.Split(dottedKey, ".")
	var current interface{} = cfg
	for _, p := range parts {
		m, ok := current.(map[string]interface{})
		if !ok {
			return "", false, nil
		}
		current, ok = m[p]
		if !ok {
			return "", false, nil
		}
	}

	s, ok := current.(string)
	if !ok {
		return fmt.Sprintf("%v", current), true, nil
	}
	return s, true, nil
}

// setNestedKey creates or overwrites the value at the path described by parts.
func setNestedKey(cfg map[string]interface{}, parts []string, value interface{}) {
	for i := 0; i < len(parts)-1; i++ {
		child, ok := cfg[parts[i]].(map[string]interface{})
		if !ok {
			child = make(map[string]interface{})
			cfg[parts[i]] = child
		}
		cfg = child
	}
	cfg[parts[len(parts)-1]] = value
}

// unsetNestedKey removes the leaf key and prunes empty parent maps.
func unsetNestedKey(cfg map[string]interface{}, parts []string) {
	if len(parts) == 0 {
		return
	}
	if len(parts) == 1 {
		delete(cfg, parts[0])
		return
	}

	child, ok := cfg[parts[0]].(map[string]interface{})
	if !ok {
		return
	}
	unsetNestedKey(child, parts[1:])
	if len(child) == 0 {
		delete(cfg, parts[0])
	}
}
