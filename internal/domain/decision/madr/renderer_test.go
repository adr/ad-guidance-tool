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

func TestRenderFile_FrontmatterAndBody(t *testing.T) {
	d := Decision{Status: "proposed", Date: "2026-05-13", Tags: []string{"infra"}}
	body := "# T\n\n## Context and Problem Statement\n\nx\n"
	out, err := RenderFile(d, body)
	assert.NoError(t, err)
	assert.Contains(t, out, "---\n")
	assert.Contains(t, out, "status: proposed")
	assert.Contains(t, out, "tags:")
	assert.Contains(t, out, "- infra")
	assert.Contains(t, out, "# T")
}

func TestRenderFile_NoFrontmatterWhenAllEmpty(t *testing.T) {
	d := Decision{}
	body := "# T\n"
	out, err := RenderFile(d, body)
	assert.NoError(t, err)
	assert.False(t, strings.HasPrefix(out, "---\n"), "expected body-only output, got: %q", out)
}

func TestRenderFile_AppendsCommentsSection(t *testing.T) {
	d := Decision{
		Status: "accepted",
		Comments: []Comment{
			{Author: "danielle", Date: "2026-05-13 14:22:01", Text: "First."},
		},
	}
	body := "# T\n\n## Context and Problem Statement\n\nx\n"
	out, err := RenderFile(d, body)
	assert.NoError(t, err)
	assert.Contains(t, out, "## Comments")
	assert.Contains(t, out, "@danielle:")
	assert.Contains(t, out, "First.")
}

func TestRenderFile_StripsExistingCommentsSectionBeforeAppending(t *testing.T) {
	d := Decision{
		Comments: []Comment{{Author: "current", Date: "2026-05-15 00:00:00", Text: "new"}},
	}
	body := "# T\n\n## Context and Problem Statement\n\nx\n\n## Comments\n\n* stale\n"
	out, err := RenderFile(d, body)
	assert.NoError(t, err)
	assert.NotContains(t, out, "* stale")
	assert.Contains(t, out, "@current:")
}
