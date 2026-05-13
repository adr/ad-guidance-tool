package madr

import (
	"fmt"
	"strings"
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
