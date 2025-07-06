package decision

import (
	domain "adg/internal/domain/decision"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type ListDecisionsPresenter struct{}

func NewListPresenter() *ListDecisionsPresenter {
	return &ListDecisionsPresenter{}
}

func (p *ListDecisionsPresenter) Listed(decisions []domain.Decision, format string) {
	if len(decisions) == 0 {
		fmt.Println("Model is empty, no decisions to list.")
		return
	}

	sortDecisionsByID(decisions)

	var output string
	var err error

	switch strings.ToLower(format) {
	case "json":
		output, err = p.renderJSON(decisions)
	case "yaml", "yml":
		output, err = p.renderYAML(decisions)
	case "md", "markdown":
		output = p.renderMarkdown(decisions)
	default:
		output = p.renderSimple(decisions)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error rendering decisions: %v\n", err)
		return
	}

	fmt.Println(output)
}

func sortDecisionsByID(decisions []domain.Decision) {
	sort.SliceStable(decisions, func(i, j int) bool {
		id1, err1 := strconv.Atoi(decisions[i].ID)
		id2, err2 := strconv.Atoi(decisions[j].ID)

		if err1 != nil || err2 != nil {
			return decisions[i].ID < decisions[j].ID
		}
		return id1 < id2
	})
}

func (p *ListDecisionsPresenter) renderJSON(decisions []domain.Decision) (string, error) {
	data, err := json.MarshalIndent(decisions, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal decisions to JSON: %w", err)
	}
	return string(data), nil
}

func (p *ListDecisionsPresenter) renderYAML(decisions []domain.Decision) (string, error) {
	data, err := yaml.Marshal(decisions)
	if err != nil {
		return "", fmt.Errorf("failed to marshal decisions to YAML: %w", err)
	}
	return string(data), nil
}

func (p *ListDecisionsPresenter) renderMarkdown(decisions []domain.Decision) string {
	var sb strings.Builder
	for _, d := range decisions {
		sb.WriteString(fmt.Sprintf("### %s - %s\n", d.ID, d.Title))
		sb.WriteString(fmt.Sprintf("- **Status:** %s\n", d.Status))
		if len(d.Tags) > 0 {
			sb.WriteString(fmt.Sprintf("- **Tags:** %s\n", strings.Join(d.Tags, ", ")))
			sb.WriteString("\n")
		}
		if len(d.Links.Precedes) > 0 {
			sb.WriteString(fmt.Sprintf("- **Precedes:** %s\n", strings.Join(d.Links.Precedes, ", ")))
		}
		if len(d.Links.Succeeds) > 0 {
			sb.WriteString(fmt.Sprintf("- **Succeeds:** %s\n", strings.Join(d.Links.Succeeds, ", ")))
		}
		if len(d.Links.Custom) > 0 {
			for tag, targets := range d.Links.Custom {
				sb.WriteString(fmt.Sprintf("- **%s:** %s\n", strings.Title(tag), strings.Join(targets, ", ")))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (p *ListDecisionsPresenter) renderSimple(decisions []domain.Decision) string {
	var sb strings.Builder
	for _, d := range decisions {
		sb.WriteString(fmt.Sprintf("%s [%s] - %s : %s\n", d.ID, d.Status, d.Title, d.Tags))
	}
	return sb.String()
}
