package decision

import (
	"strings"
	"testing"
)

func TestCommented(t *testing.T) {
	presenter := NewCommentPresenter()

	output := captureOutput(func() {
		presenter.Commented("0001", "alice", "Great idea!")
	})

	expected := `Comment added by alice to decision 0001: "Great idea!"`

	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain: %q\nGot: %q", expected, output)
	}
}
