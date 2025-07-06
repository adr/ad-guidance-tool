package decision

import (
	"adg/internal/domain/decision"
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestAdded_OnlySuccesses(t *testing.T) {
	presenter := NewAddPresenter()
	successes := []*decision.Decision{
		{ID: "0001", Title: "First"},
		{ID: "0002", Title: "Second"},
	}

	output := captureOutput(func() {
		presenter.Added(successes, map[string]error{})
	})

	if !strings.Contains(output, "Decision First (0001) added successfully.") {
		t.Errorf("Expected success message for 'First', got: %s", output)
	}
	if !strings.Contains(output, "Decision Second (0002) added successfully.") {
		t.Errorf("Expected success message for 'Second', got: %s", output)
	}
}

func TestAdded_OnlyFailures(t *testing.T) {
	presenter := NewAddPresenter()
	failures := map[string]error{
		"Invalid A": fmt.Errorf("something went wrong"),
		"Invalid B": fmt.Errorf("duplicate title"),
	}

	output := captureOutput(func() {
		presenter.Added(nil, failures)
	})

	if !strings.Contains(output, `Failed to add decision "Invalid A": something went wrong`) {
		t.Errorf("Missing error for 'Invalid A'")
	}
	if !strings.Contains(output, `Failed to add decision "Invalid B": duplicate title`) {
		t.Errorf("Missing error for 'Invalid B'")
	}
}

func TestAdded_Mixed(t *testing.T) {
	presenter := NewAddPresenter()
	successes := []*decision.Decision{
		{ID: "0003", Title: "Valid"},
	}
	failures := map[string]error{
		"Broken": fmt.Errorf("bad format"),
	}

	output := captureOutput(func() {
		presenter.Added(successes, failures)
	})

	if !strings.Contains(output, "Decision Valid (0003) added successfully.") {
		t.Errorf("Expected success message")
	}
	if !strings.Contains(output, `Failed to add decision "Broken": bad format`) {
		t.Errorf("Expected error message")
	}
}

// helper
func captureOutput(f func()) string {
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout
	buf.ReadFrom(r)
	return buf.String()
}
