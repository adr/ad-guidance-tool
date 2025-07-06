package decision

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecided(t *testing.T) {
	presenter := NewDecidePresenter()
	decisionID := "0010"

	output := captureOutput(func() {
		presenter.Decided(decisionID)
	})

	expected := fmt.Sprintf("Decision %s has been marked as decided.\n", decisionID)
	assert.Equal(t, expected, output)
	assert.True(t, strings.Contains(output, decisionID))
}
