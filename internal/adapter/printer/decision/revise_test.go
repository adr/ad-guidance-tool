package decision

import (
	"strings"
	"testing"
)

func TestRevised(t *testing.T) {
	presenter := NewRevisePresenter()

	output := captureOutput(func() {
		presenter.Revised("0001", "0002")
	})

	expected := "Successfully revised decision 0001 → new decision 0002"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain: %q, but got: %q", expected, output)
	}
}
