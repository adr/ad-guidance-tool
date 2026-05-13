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

func TestParseFrontmatter_Full(t *testing.T) {
	yml := `status: "accepted"
date: 2026-05-13
decision-makers:
  - "danielle"
tags:
  - infrastructure
links:
  related-to:
    - "0004"
comments:
  - author: "danielle"
    date: "2026-05-13 14:22:01"
    text: "Initial."
`
	fm, err := ParseFrontmatter(yml)
	assert.NoError(t, err)
	assert.Equal(t, "accepted", fm.Status)
	assert.Equal(t, []string{"danielle"}, fm.DecisionMakers)
	assert.Equal(t, []string{"infrastructure"}, fm.Tags)
	assert.Equal(t, []string{"0004"}, fm.Links["related-to"])
	assert.Len(t, fm.Comments, 1)
	assert.Equal(t, "Initial.", fm.Comments[0].Text)
}

func TestParseFrontmatter_Empty(t *testing.T) {
	fm, err := ParseFrontmatter("")
	assert.NoError(t, err)
	assert.Equal(t, Frontmatter{}, fm)
}

func TestParseFilename_Valid(t *testing.T) {
	id, slug, err := ParseFilename("0042-use-kafka.md")
	assert.NoError(t, err)
	assert.Equal(t, "0042", id)
	assert.Equal(t, "use-kafka", slug)
}

func TestParseFilename_WithSubdirectory(t *testing.T) {
	id, slug, err := ParseFilename("infra/0042-use-kafka.md")
	assert.NoError(t, err)
	assert.Equal(t, "0042", id)
	assert.Equal(t, "use-kafka", slug)
}

func TestParseFilename_Invalid(t *testing.T) {
	_, _, err := ParseFilename("AD0042-use-kafka.md")
	assert.Error(t, err)
	_, _, err = ParseFilename("0042.md")
	assert.Error(t, err)
}

func TestIsLegacyADG_DetectsFilenamePrefix(t *testing.T) {
	assert.True(t, IsLegacyADG("AD0001-foo.md", []byte("# T")))
}

func TestIsLegacyADG_DetectsBodyAnchor(t *testing.T) {
	assert.True(t, IsLegacyADG("0001-foo.md", []byte(`# T
## <a name="question"></a> Question
`)))
}

func TestIsLegacyADG_DetectsLegacyStatus(t *testing.T) {
	assert.True(t, IsLegacyADG("0001-foo.md", []byte("---\nstatus: open\n---\n")))
}

func TestIsLegacyADG_PureMADRPasses(t *testing.T) {
	assert.False(t, IsLegacyADG("0001-foo.md", []byte("---\nstatus: accepted\n---\n# T\n")))
}
