package decision

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinked_WithReverseTag(t *testing.T) {
	presenter := NewLinkPresenter()
	sourceID := "001"
	targetID := "002"
	tag := "depends on"
	reverseTag := "supports"

	output := captureOutput(func() {
		presenter.Linked(sourceID, targetID, tag, reverseTag)
	})

	expected1 := fmt.Sprintf("Link added: %s →[%s]→ %s\n", sourceID, tag, targetID)
	expected2 := fmt.Sprintf("Reverse link added: %s →[%s]→ %s\n", targetID, reverseTag, sourceID)

	assert.Contains(t, output, expected1)
	assert.Contains(t, output, expected2)
}

func TestLinked_WithoutReverseTag(t *testing.T) {
	presenter := NewLinkPresenter()
	sourceID := "003"
	targetID := "004"
	tag := "blocks"
	reverseTag := ""

	output := captureOutput(func() {
		presenter.Linked(sourceID, targetID, tag, reverseTag)
	})

	expected := fmt.Sprintf("Link added: %s →[%s]→ %s\n", sourceID, tag, targetID)

	assert.Equal(t, expected, output)
	assert.NotContains(t, output, "Reverse link added")
}
