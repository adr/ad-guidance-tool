package madr

import (
	"bytes"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

const canonicalTemplate = `# %s

## Context and Problem Statement

{...}

## Decision Drivers

* {driver 1}

## Considered Options

* {option 1}
* {option 2}

## Decision Outcome

Chosen option: "{option title}", because {justification}.

### Consequences

* Good, because {...}
* Bad, because {...}
`

// RenderNewBody emits the canonical minimal+Decision-Drivers template for a
// freshly-created ADR. The title is interpolated into the H1.
func RenderNewBody(title string) string {
	return fmt.Sprintf(canonicalTemplate, title)
}

// renderCommentsSection produces the trailing ## Comments H2 from a Comment
// list. Returns "" if the list is empty (the section is omitted entirely).
func renderCommentsSection(comments []Comment) string {
	if len(comments) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("## Comments\n\n")
	for _, c := range comments {
		b.WriteString(fmt.Sprintf("* **%s — @%s:** %s\n", c.Date, c.Author, c.Text))
	}
	return b.String()
}

// RenderFile assembles the on-disk bytes for an ADR: optional YAML frontmatter
// between `---` fences, then the body. Any existing `## Comments` section in
// the body is stripped, and a fresh one is rendered from d.Comments at the end.
//
// If d has no populated frontmatter fields, no frontmatter block is emitted —
// MADR's minimal template is frontmatter-free, and this respects that case.
func RenderFile(d Decision, body string) (string, error) {
	stripped := stripCommentsSection(body)
	stripped = strings.TrimRight(stripped, "\n") + "\n"

	commentsSection := renderCommentsSection(d.Comments)

	fm := d.Frontmatter()
	fmBytes, err := yaml.Marshal(fm)
	if err != nil {
		return "", fmt.Errorf("failed to marshal frontmatter: %w", err)
	}

	// yaml.Marshal of a fully-zero struct produces "{}\n". Trim and detect.
	hasFrontmatter := len(bytes.TrimSpace(fmBytes)) > 2

	var out bytes.Buffer
	if hasFrontmatter {
		out.WriteString("---\n")
		out.Write(fmBytes)
		out.WriteString("---\n\n")
	}
	out.WriteString(stripped)
	if commentsSection != "" {
		out.WriteString("\n")
		out.WriteString(commentsSection)
	}
	return out.String(), nil
}

// stripCommentsSection removes the `## Comments` H2 and its contents from a body.
// Anything after the next H2 (or EOF) is preserved.
func stripCommentsSection(body string) string {
	lines := strings.Split(body, "\n")
	var out []string
	skipping := false
	for _, line := range lines {
		if !skipping && strings.HasPrefix(strings.ToLower(strings.TrimSpace(line)), "## comments") {
			skipping = true
			continue
		}
		if skipping && strings.HasPrefix(strings.TrimSpace(line), "## ") {
			skipping = false
		}
		if !skipping {
			out = append(out, line)
		}
	}
	return strings.Join(out, "\n")
}
