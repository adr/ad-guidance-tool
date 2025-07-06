package decision

import (
	"strings"
	"testing"
)

func TestTagged(t *testing.T) {
	presenter := NewTagPresenter()

	output := captureOutput(func() {
		presenter.Tagged("0001", []string{"critical", "UI"})
	})

	expected := "Tags [critical, UI] added to decision 0001"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain: %q, but got: %q", expected, output)
	}
}
