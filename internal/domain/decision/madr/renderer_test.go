package madr

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderNewBody_CanonicalTemplate(t *testing.T) {
	body := RenderNewBody("Use Kafka")
	assert.True(t, strings.HasPrefix(body, "# Use Kafka\n"))
	assert.Contains(t, body, "## Context and Problem Statement")
	assert.Contains(t, body, "## Decision Drivers")
	assert.Contains(t, body, "## Considered Options")
	assert.Contains(t, body, "## Decision Outcome")
	assert.Contains(t, body, "### Consequences")
}
