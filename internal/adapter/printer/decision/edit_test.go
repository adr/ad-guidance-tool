package decision

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdited(t *testing.T) {
	presenter := NewEditPresenter()
	decisionID := "0077"

	output := captureOutput(func() {
		presenter.Edited(decisionID)
	})

	expected := fmt.Sprintf("Decision %s updated successfully.\n", decisionID)
	assert.Equal(t, expected, output)
	assert.True(t, strings.Contains(output, decisionID))
}
