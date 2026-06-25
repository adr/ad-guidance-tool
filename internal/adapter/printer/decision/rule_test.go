package decision

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRulePresenter_RuleGenerated(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	presenter := NewRulePresenter()
	presenter.RuleGenerated("0001", "/path/to/rule.rule")

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	assert.Contains(t, output, "Rule file for decision 0001 created at: /path/to/rule.rule")
}
