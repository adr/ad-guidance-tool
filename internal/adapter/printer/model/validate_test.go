package model

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"adg/internal/application/outputport"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = old

	return buf.String()
}

func TestModelValidatePresenter_ModelValidated_NoIssues(t *testing.T) {
	presenter := NewModelValidatePresenter()

	output := captureOutput(func() {
		presenter.ModelValidated("test-model", nil)
	})

	expected := "test-model model is valid\n"
	if output != expected {
		t.Errorf("unexpected output:\nexpected:\n%s\ngot:\n%s", expected, output)
	}
}

func TestModelValidatePresenter_ModelValidated_WithIssues(t *testing.T) {
	presenter := NewModelValidatePresenter()

	issues := []outputport.ValidationIssue{
		{ID: "0001", Message: "filename does not match NNNN-slug.md"},
		{ID: "0002", Message: "comment 1 has empty text"},
	}

	output := captureOutput(func() {
		presenter.ModelValidated("test-model", issues)
	})

	for _, expected := range []string{
		"test-model model has 2 validation issue(s):",
		"ID 0001: filename does not match NNNN-slug.md",
		"ID 0002: comment 1 has empty text",
	} {
		if !strings.Contains(output, expected) {
			t.Errorf("expected %q in output, got:\n%s", expected, output)
		}
	}
}
