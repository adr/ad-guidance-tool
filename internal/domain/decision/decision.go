package decision

import "fmt"

type Decision struct {
	ID       string    `yaml:"adr_id"`
	Title    string    `yaml:"title"`
	Status   string    `yaml:"status"`
	Tags     []string  `yaml:"tags"`
	Links    Links     `yaml:"links"`
	Comments []Comment `yaml:"comments"`
}

type Links struct {
	Precedes []string            `yaml:"precedes"`
	Succeeds []string            `yaml:"succeeds"`
	Custom   map[string][]string `yaml:"custom,omitempty"`
}

type Comment struct {
	Author  string `yaml:"author"`
	Date    string `yaml:"date"`
	Comment string `yaml:"comment"`
}

type DecisionContent struct {
	ID       string
	Question string
	Criteria string
	Options  string
	Outcome  string
	Comments string
}

// flattens custom links into the main links block
func (l Links) MarshalYAML() (any, error) {
	out := make(map[string]any)

	out["precedes"] = l.Precedes
	out["succeeds"] = l.Succeeds

	for key, value := range l.Custom {
		out[key] = value
	}

	return out, nil
}

// reconstructs fixed and custom links from flat YAML structure
func (l *Links) UnmarshalYAML(unmarshal func(any) error) error {
	raw := make(map[string]any)
	if err := unmarshal(&raw); err != nil {
		return err
	}

	if v, ok := raw["precedes"]; ok {
		l.Precedes = asStringSlice(v)
	}

	if v, ok := raw["succeeds"]; ok {
		l.Succeeds = asStringSlice(v)
	}

	l.Custom = make(map[string][]string)

	for k, v := range raw {
		if k != "precedes" && k != "succeeds" {
			l.Custom[k] = asStringSlice(v)
		}
	}

	return nil
}

func asStringSlice(v any) []string {
	if v == nil {
		return []string{}
	}

	switch t := v.(type) {
	case []any:
		result := make([]string, len(t))
		for i, val := range t {
			result[i] = fmt.Sprintf("%v", val)
		}
		return result
	case []string:
		return t
	default:
		return []string{}
	}
}
