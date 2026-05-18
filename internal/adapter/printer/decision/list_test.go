package decision

import (
	"adg/internal/domain/decision"
	"strings"
	"testing"
)

func TestListed_JSON(t *testing.T) {
	presenter := NewListPresenter()
	decisions := []decision.Decision{
		{ID: "0002", Title: "Decision B", Status: "proposed"},
		{ID: "0001", Title: "Decision A", Status: "accepted"},
	}

	output := captureOutput(func() {
		presenter.Listed(decisions, "json")
	})

	if !strings.Contains(output, `"Title": "Decision A"`) || !strings.Contains(output, `"Title": "Decision B"`) {
		t.Errorf("JSON output missing expected content:\n%s", output)
	}
}

func TestListed_YAML(t *testing.T) {
	presenter := NewListPresenter()
	decisions := []decision.Decision{
		{ID: "0002", Title: "Decision B", Status: "proposed"},
	}

	output := captureOutput(func() {
		presenter.Listed(decisions, "yaml")
	})

	if !strings.Contains(output, "title: Decision B") {
		t.Errorf("YAML output missing expected content:\n%s", output)
	}
}

func TestListed_Markdown(t *testing.T) {
	presenter := NewListPresenter()
	decisions := []decision.Decision{
		{
			ID: "0001", Title: "Decision A", Status: "proposed",
			Tags:       []string{"tag1"},
			Supersedes: []string{"0002"},
			Links: map[string][]string{
				"related": {"0004"},
			},
		},
	}

	output := captureOutput(func() {
		presenter.Listed(decisions, "md")
	})

	if !strings.Contains(output, "### 0001 - Decision A") || !strings.Contains(output, "- **Supersedes:** 0002") || !strings.Contains(output, "**Related:** 0004") {
		t.Errorf("Markdown output missing expected content:\n%s", output)
	}
}

func TestListed_Simple(t *testing.T) {
	presenter := NewListPresenter()
	decisions := []decision.Decision{
		{ID: "0001", Title: "Simple Decision", Status: "proposed", Tags: []string{"alpha", "beta"}},
	}

	output := captureOutput(func() {
		presenter.Listed(decisions, "simple")
	})

	if !strings.Contains(output, "0001 [proposed] - Simple Decision : [alpha beta]") {
		t.Errorf("Simple output format incorrect:\n%s", output)
	}
}

func TestListed_EmptyModel(t *testing.T) {
	presenter := NewListPresenter()
	output := captureOutput(func() {
		presenter.Listed([]decision.Decision{}, "json")
	})

	if !strings.Contains(output, "Model is empty") {
		t.Errorf("Empty model message not shown:\n%s", output)
	}
}
