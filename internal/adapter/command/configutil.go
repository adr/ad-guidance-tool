package commands

import (
	domain "adg/internal/domain/config"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func ResolveModelPathOrDefault(flagValue string, config domain.ConfigService) (string, error) {
	if flagValue != "" {
		return flagValue, nil
	}
	if !config.IsLoaded() || config.GetDefaultModelPath() == "" {
		return "", fmt.Errorf("model path must be provided via --model or config")
	}
	return config.GetDefaultModelPath(), nil
}

func ResolveIdOrTitle(idOrTitle string, id, title *string) error {
	if idOrTitle == "" {
		return fmt.Errorf("you must specify the decisions via --id by either providing the numbered id (e.g., 0001) or the name of the decision (e.g, 'my-decision')")
	}

	if matched, _ := regexp.MatchString(`^\d{4}$`, idOrTitle); matched {
		*id = idOrTitle // dereference and assign
		*title = ""     // clear title
		return nil
	}

	if matched, _ := regexp.MatchString(`[a-zA-Z]`, idOrTitle); matched {
		*title = idOrTitle // dereference and assign
		*id = ""           // clear id
		return nil
	}

	return errors.New("input must be either a 4-digit ID or a title containing at least one letter")
}

func GetTemplateSections(template string) (map[string]string, error) {
	switch strings.ToLower(template) {
	case "nygard":
		return map[string]string{
			"question": "Context",
			"criteria": "Consequences",
			"outcome":  "Decision",
		}, nil
	case "madr":
		return map[string]string{
			"question": "Context and Problem Statement",
			"options":  "Considered Options",
			"criteria": "Decision Drivers",
			"outcome":  "Decision Outcome",
		}, nil
	default:
		return nil, fmt.Errorf("unknown template: %q (available: Nygard, MADR)", template)
	}
}
