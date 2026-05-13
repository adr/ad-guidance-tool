package madr

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitFile_WithFrontmatter(t *testing.T) {
	in := "---\nstatus: proposed\n---\n\n# Title\n\nbody\n"
	fm, body, err := SplitFile([]byte(in))
	assert.NoError(t, err)
	assert.Equal(t, "status: proposed\n", fm)
	assert.True(t, strings.HasPrefix(body, "# Title"))
}

func TestSplitFile_NoFrontmatter(t *testing.T) {
	in := "# Title\n\nbody\n"
	fm, body, err := SplitFile([]byte(in))
	assert.NoError(t, err)
	assert.Equal(t, "", fm)
	assert.True(t, strings.HasPrefix(body, "# Title"))
}

func TestSplitFile_FrontmatterMissingCloser(t *testing.T) {
	in := "---\nstatus: proposed\n\n# Title\n"
	_, _, err := SplitFile([]byte(in))
	assert.Error(t, err)
}

func TestParseBody_FindsCanonicalSections(t *testing.T) {
	body := `# Title

## Context and Problem Statement

Some context.

## Considered Options

* A
* B

## Decision Outcome

Chosen option: "A", because reasons.
`
	parsed, err := ParseBody(body)
	assert.NoError(t, err)
	assert.Equal(t, "Title", parsed.Title)
	assert.Contains(t, parsed.Sections, "context")
	assert.Contains(t, parsed.Sections, "options")
	assert.Contains(t, parsed.Sections, "outcome")
	assert.Equal(t, []string{"A", "B"}, parsed.Options)
}

func TestParseBody_CaseInsensitiveHeaders(t *testing.T) {
	body := `# T

## context and problem statement

x

## CONSIDERED OPTIONS

* A
`
	parsed, err := ParseBody(body)
	assert.NoError(t, err)
	assert.Contains(t, parsed.Sections, "context")
	assert.Contains(t, parsed.Sections, "options")
}

func TestParseBody_PreservesUnknownH2(t *testing.T) {
	body := `# T

## Context and Problem Statement

x

## Risks

* something
`
	parsed, err := ParseBody(body)
	assert.NoError(t, err)
	assert.Contains(t, parsed.CustomSections, "Risks")
}

func TestParseBody_ChosenOption(t *testing.T) {
	body := `# T

## Considered Options

* A
* B

## Decision Outcome

Chosen option: "B", because B is better.
`
	parsed, err := ParseBody(body)
	assert.NoError(t, err)
	assert.Equal(t, "B", parsed.ChosenOption)
	assert.Equal(t, "B is better", parsed.OutcomeRationale)
}
