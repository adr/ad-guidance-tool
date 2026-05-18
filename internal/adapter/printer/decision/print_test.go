package decision

import (
	"adg/internal/application/outputport"
	"strings"
	"testing"
)

const sampleBody = `# Decision Title

## Context and Problem Statement

What should we do?

## Considered Options

* Option A
* Option B

## Decision Outcome

Chosen option: "Option A", because reasons.

## Comments

### 2026-01-01 12:00:00 — Jane

Looks good.
`

func TestPrinted_NoFilter_PrintsFullBody(t *testing.T) {
	presenter := NewPrintPresenter()

	bodies := []outputport.DecisionBody{
		{ID: "0002", Body: sampleBody},
		{ID: "0001", Body: sampleBody},
	}

	output := captureOutput(func() {
		presenter.Printed(bodies, nil)
	})

	for _, expected := range []string{
		"===== Decision 0001 =====",
		"===== Decision 0002 =====",
		"## Context and Problem Statement",
		"## Considered Options",
		"## Decision Outcome",
	} {
		if !strings.Contains(output, expected) {
			t.Errorf("expected output to contain %q\n%s", expected, output)
		}
	}

	// IDs should be sorted ascending.
	if strings.Index(output, "Decision 0001") > strings.Index(output, "Decision 0002") {
		t.Errorf("expected 0001 to print before 0002:\n%s", output)
	}
}

func TestPrinted_FiltersBySection(t *testing.T) {
	presenter := NewPrintPresenter()

	bodies := []outputport.DecisionBody{
		{ID: "0001", Body: sampleBody},
	}

	output := captureOutput(func() {
		presenter.Printed(bodies, map[string]bool{"context": true, "outcome": true})
	})

	if !strings.Contains(output, "What should we do?") {
		t.Errorf("expected Context section, got:\n%s", output)
	}
	if !strings.Contains(output, "Chosen option:") {
		t.Errorf("expected Outcome section, got:\n%s", output)
	}
	if strings.Contains(output, "* Option A") {
		t.Errorf("Considered Options should not be printed when filtered out:\n%s", output)
	}
}
