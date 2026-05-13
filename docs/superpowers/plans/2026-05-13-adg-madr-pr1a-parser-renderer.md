# PR 1a — MADR Parser, Renderer, Types (Additive Subpackage)

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Land all the MADR-shaped types, body and frontmatter parser, renderer, fixture set, and round-trip property test as a new `internal/domain/decision/madr/` subpackage. **Purely additive**: no existing code is modified or deleted. After this PR, `go build ./...` and `go test ./...` both pass; the new package is used by no one yet.

**Architecture:** New isolated subpackage. The `madr.Decision` struct is the future canonical type but lives in `madr/` until PR 1b moves it up and retires the legacy types. Parser is pure markdown (no HTML anchors). Renderer regenerates the `## Comments` body section from `madr.Comment` frontmatter list. Round-trip property `parse(f) → render → f' ⟹ diff(f, f') == ∅` is the load-bearing invariant.

**Tech Stack:** Go 1.22, `gopkg.in/yaml.v3`, `github.com/stretchr/testify/assert`.

**Reference spec:** `docs/superpowers/specs/2026-05-13-adg-madr-fork-design.md` — file format / data model section, and the round-trip property statement.

---

## Pre-flight ratifications

The spec leaves two design calls open. This PR series commits to:

1. **`index.yaml` dropped** entirely in PR 1b. PR 1a doesn't touch it.
2. **`adg revise` stays independent** of supersession. Not exercised in PR 1a.

If either should flip, raise it before Task 1.

---

## File structure

**New files (all created; no modifications, no deletions):**
- `testdata/fixtures/madr/full.md`
- `testdata/fixtures/madr/minimal.md`
- `testdata/fixtures/madr/bare.md`
- `testdata/fixtures/madr/bare-minimal.md`
- `testdata/fixtures/madr/canonical.md`
- `testdata/fixtures/madr/custom-h2.md`
- `testdata/fixtures/madr/h3-subsections.md`
- `testdata/fixtures/madr/no-frontmatter.md`
- `testdata/fixtures/madr/all-extensions.md`
- `internal/domain/decision/madr/types.go`
- `internal/domain/decision/madr/types_test.go`
- `internal/domain/decision/madr/parser.go`
- `internal/domain/decision/madr/parser_test.go`
- `internal/domain/decision/madr/renderer.go`
- `internal/domain/decision/madr/renderer_test.go`
- `internal/domain/decision/madr/roundtrip_test.go`

No file in `cmd/`, `internal/adapter/`, `internal/application/`, or `internal/infrastructure/` is touched in this PR.

---

## Task 1: MADR fixture set

**Files:** the 9 fixtures listed above under `testdata/fixtures/madr/`.

- [ ] **Step 1: Create directory and fetch upstream MADR 4.0 templates**

```bash
mkdir -p testdata/fixtures/madr
curl -sL https://raw.githubusercontent.com/adr/madr/4.0.0/template/adr-template.md > testdata/fixtures/madr/full.md
curl -sL https://raw.githubusercontent.com/adr/madr/4.0.0/template/adr-template-minimal.md > testdata/fixtures/madr/minimal.md
curl -sL https://raw.githubusercontent.com/adr/madr/4.0.0/template/adr-template-bare.md > testdata/fixtures/madr/bare.md
curl -sL https://raw.githubusercontent.com/adr/madr/4.0.0/template/adr-template-bare-minimal.md > testdata/fixtures/madr/bare-minimal.md
```

- [ ] **Step 2: Write `canonical.md`** — the fork's default template:

```markdown
# {short title, representative of solved problem and found solution}

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
```

- [ ] **Step 3: Write `custom-h2.md`**:

```markdown
---
status: "proposed"
date: 2026-05-13
---

# Migrate from MySQL to PostgreSQL

## Context and Problem Statement

We need a database with stronger JSONB support.

## Considered Options

* Stay on MySQL with workarounds
* Migrate to PostgreSQL

## Decision Outcome

Chosen option: "Migrate to PostgreSQL", because JSONB support is first-class.

## Risks

* Migration tooling immaturity
* 6-month dual-write window

## Open Questions

* Do we colocate read replicas?
```

- [ ] **Step 4: Write `h3-subsections.md`**:

```markdown
# Adopt structured logging

## Context and Problem Statement

Free-form logs are hard to grep.

## Considered Options

* Keep free-form
* Adopt structured logging (zap)

## Decision Outcome

Chosen option: "Adopt structured logging (zap)", because grep-friendly logs cost less in incidents.

### Consequences

* Good, because incident MTTR drops.
* Bad, because every site needs to be migrated.

### Confirmation

Lint rule rejects `fmt.Println` in production code.
```

- [ ] **Step 5: Write `no-frontmatter.md`**:

```markdown
# An ADR without YAML frontmatter

## Context and Problem Statement

Sometimes you want to skip the metadata.

## Considered Options

* Add frontmatter
* Skip frontmatter

## Decision Outcome

Chosen option: "Skip frontmatter", because the minimal template permits it.
```

- [ ] **Step 6: Write `all-extensions.md`**:

```markdown
---
status: "accepted"
date: 2026-05-13
decision-makers:
  - "danielle"
consulted:
  - "rsmith"
informed: []
tags:
  - infrastructure
  - migration
links:
  related-to:
    - "0004"
supersedes:
  - "0017"
comments:
  - author: "danielle"
    date: "2026-05-13 14:22:01"
    text: "Initial decision; revisit after Q3."
  - author: "rsmith"
    date: "2026-05-14 09:00:00"
    text: "Confirmed in prod load test."
---

# Use Kafka for the event bus

## Context and Problem Statement

We need an event bus that scales to 100k msg/sec.

## Considered Options

* Use Kafka
* Use NATS
* Roll our own

## Decision Outcome

Chosen option: "Use Kafka", because the operations team already runs it.

### Consequences

* Good, because zero new operational burden.
* Bad, because Kafka's footprint is heavyweight for our scale.

## Comments

* **2026-05-13 14:22:01 — @danielle:** Initial decision; revisit after Q3.
* **2026-05-14 09:00:00 — @rsmith:** Confirmed in prod load test.
```

- [ ] **Step 7: Commit**

```bash
git add testdata/fixtures/madr/
git commit -m "test: add MADR 4.0 fixture set for parser/renderer tests"
```

---

## Task 2: MADR types (Decision, Comment, Frontmatter)

**Files:**
- Create: `internal/domain/decision/madr/types.go`
- Create: `internal/domain/decision/madr/types_test.go`

- [ ] **Step 1: Write failing test**

```go
package madr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestFrontmatter_YAMLRoundtrip(t *testing.T) {
	fm := Frontmatter{
		Status:         "accepted",
		Date:           "2026-05-13",
		DecisionMakers: []string{"danielle"},
		Tags:           []string{"infra"},
		Links:          map[string][]string{"related-to": {"0004"}},
		Supersedes:     []string{"0017"},
		Comments: []Comment{
			{Author: "danielle", Date: "2026-05-13 14:22:01", Text: "First."},
		},
	}

	out, err := yaml.Marshal(fm)
	assert.NoError(t, err)

	var got Frontmatter
	assert.NoError(t, yaml.Unmarshal(out, &got))
	assert.Equal(t, fm, got)
}

func TestFrontmatter_LegacyOutcome_Omitempty(t *testing.T) {
	fm := Frontmatter{Status: "proposed"}
	out, err := yaml.Marshal(fm)
	assert.NoError(t, err)
	assert.NotContains(t, string(out), "legacy-outcome")
}

func TestDecision_ToFromFrontmatter(t *testing.T) {
	d := Decision{
		ID:     "0042",
		Slug:   "use-kafka",
		Title:  "Use Kafka",
		Status: "accepted",
		Tags:   []string{"infra"},
	}
	fm := d.Frontmatter()
	assert.Equal(t, "accepted", fm.Status)
	assert.Equal(t, []string{"infra"}, fm.Tags)

	round := DecisionFromFrontmatter(fm)
	round.ID, round.Slug, round.Title = d.ID, d.Slug, d.Title
	assert.Equal(t, d, round)
}
```

- [ ] **Step 2: Run, expect compilation failure** (`madr` package doesn't exist):

```bash
go test ./internal/domain/decision/madr/
```

- [ ] **Step 3: Create `types.go`**

```go
// Package madr defines the MADR 4.0–native types, parser, and renderer used by
// the fork's file format. This package is self-contained; integration with the
// rest of the codebase happens in PR 1b.
package madr

// Decision is the in-memory representation of an Architectural Decision Record.
//
// ID, Slug, and Title are derived (filename + H1) and not stored in frontmatter.
// All other fields are persisted to frontmatter via Frontmatter().
type Decision struct {
	// Identity (from filename)
	ID   string
	Slug string

	// Title (from H1)
	Title string

	// MADR frontmatter
	Status         string
	Date           string
	DecisionMakers []string
	Consulted      []string
	Informed       []string

	// ADG extensions
	Tags          []string
	Links         map[string][]string
	Supersedes    []string
	Comments      []Comment
	LegacyOutcome bool
}

// Comment is one entry in the ADG-extension `comments` frontmatter list.
// Date is a timestamp with time-of-day (YYYY-MM-DD HH:MM:SS) to preserve
// ordering within a single day; the ADR-level Decision.Date is day-precision.
type Comment struct {
	Author string `yaml:"author"`
	Date   string `yaml:"date"`
	Text   string `yaml:"text"`
}

// Frontmatter is the YAML-serializable shape persisted at the top of an ADR file.
// Every field is omitempty so the fork respects MADR's "frontmatter is optional"
// principle for minimal ADRs.
type Frontmatter struct {
	Status         string              `yaml:"status,omitempty"`
	Date           string              `yaml:"date,omitempty"`
	DecisionMakers []string            `yaml:"decision-makers,omitempty"`
	Consulted      []string            `yaml:"consulted,omitempty"`
	Informed       []string            `yaml:"informed,omitempty"`
	Tags           []string            `yaml:"tags,omitempty"`
	Links          map[string][]string `yaml:"links,omitempty"`
	Supersedes     []string            `yaml:"supersedes,omitempty"`
	Comments       []Comment           `yaml:"comments,omitempty"`
	LegacyOutcome  bool                `yaml:"legacy-outcome,omitempty"`
}

func (d Decision) Frontmatter() Frontmatter {
	return Frontmatter{
		Status:         d.Status,
		Date:           d.Date,
		DecisionMakers: d.DecisionMakers,
		Consulted:      d.Consulted,
		Informed:       d.Informed,
		Tags:           d.Tags,
		Links:          d.Links,
		Supersedes:     d.Supersedes,
		Comments:       d.Comments,
		LegacyOutcome:  d.LegacyOutcome,
	}
}

func DecisionFromFrontmatter(fm Frontmatter) Decision {
	return Decision{
		Status:         fm.Status,
		Date:           fm.Date,
		DecisionMakers: fm.DecisionMakers,
		Consulted:      fm.Consulted,
		Informed:       fm.Informed,
		Tags:           fm.Tags,
		Links:          fm.Links,
		Supersedes:     fm.Supersedes,
		Comments:       fm.Comments,
		LegacyOutcome:  fm.LegacyOutcome,
	}
}
```

- [ ] **Step 4: Run, expect pass**

```bash
go test ./internal/domain/decision/madr/ -v
```

- [ ] **Step 5: Commit**

```bash
git add internal/domain/decision/madr/types.go internal/domain/decision/madr/types_test.go
git commit -m "feat(madr): Decision, Comment, Frontmatter types"
```

---

## Task 3: Parser — SplitFile (frontmatter / body split)

**Files:**
- Create: `internal/domain/decision/madr/parser.go`
- Create: `internal/domain/decision/madr/parser_test.go`

- [ ] **Step 1: Write failing tests**

```go
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
```

- [ ] **Step 2: Run, expect failure** (`SplitFile` undefined).

- [ ] **Step 3: Create `parser.go`**

```go
package madr

import (
	"bytes"
	"fmt"
)

// SplitFile separates the optional YAML frontmatter (between `---` fences at the
// top of the file) from the markdown body. Returns frontmatter text without the
// fences (may be empty), body text, or an error if the frontmatter is opened but
// never closed.
func SplitFile(content []byte) (frontmatter, body string, err error) {
	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))

	if !bytes.HasPrefix(content, []byte("---\n")) {
		return "", string(content), nil
	}

	rest := content[len("---\n"):]
	closeIdx := bytes.Index(rest, []byte("\n---\n"))
	if closeIdx == -1 {
		if bytes.HasSuffix(rest, []byte("\n---")) {
			closeIdx = len(rest) - len("\n---")
			return string(rest[:closeIdx]), "", nil
		}
		return "", "", fmt.Errorf("frontmatter opened with `---` but never closed")
	}

	fm := string(rest[:closeIdx+1])
	bodyStart := closeIdx + len("\n---\n")
	return fm, string(rest[bodyStart:]), nil
}
```

- [ ] **Step 4: Run, expect pass**

```bash
go test ./internal/domain/decision/madr/ -run TestSplitFile -v
```

- [ ] **Step 5: Commit**

```bash
git add internal/domain/decision/madr/parser.go internal/domain/decision/madr/parser_test.go
git commit -m "feat(madr): SplitFile separates YAML frontmatter from markdown body"
```

---

## Task 4: Parser — body sections, options, chosen option

**Files:**
- Modify: `internal/domain/decision/madr/parser.go`
- Modify: `internal/domain/decision/madr/parser_test.go`

- [ ] **Step 1: Add failing tests**

```go
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
```

- [ ] **Step 2: Run, expect failure** (`ParseBody`, `ParsedBody` undefined).

- [ ] **Step 3: Append to `parser.go`**

```go
import (
	// ... existing ...
	"regexp"
	"strings"
)

// ParsedBody is the result of ParseBody — everything we extract from a body.
type ParsedBody struct {
	Title            string
	Sections         map[string]string // canonical key (lowercase) -> raw section text including the H2 line
	Options          []string          // bullet items under Considered Options, in order
	ChosenOption     string            // text from `Chosen option: "..."`
	OutcomeRationale string            // text after `because ` and before the trailing `.`
	CustomSections   map[string]string // unrecognized H2 header text -> raw section text
}

// canonicalSections maps lowercased H2 header text to a short key.
// We use exact equality (case-insensitive) on the header text — NOT
// contains-style matching — so a header like "Considered Trade-offs"
// is treated as a custom section, not as "Considered Options".
var canonicalSections = map[string]string{
	"context and problem statement": "context",
	"decision drivers":              "drivers",
	"considered options":            "options",
	"decision outcome":              "outcome",
	"pros and cons of the options":  "pros-cons",
	"more information":              "more",
	"comments":                      "comments",
}

var (
	h1Re      = regexp.MustCompile(`(?m)^# +(.+)$`)
	h2Re      = regexp.MustCompile(`(?m)^## +(.+?)\s*$`)
	bulletRe  = regexp.MustCompile(`(?m)^\s*\*\s+(.+)$`)
	chosenRe  = regexp.MustCompile(`(?m)^Chosen option:\s*"([^"]*)"(?:\s*,\s*because\s+(.+?))?\.?\s*$`)
)

func ParseBody(body string) (*ParsedBody, error) {
	pb := &ParsedBody{
		Sections:       map[string]string{},
		CustomSections: map[string]string{},
	}

	if m := h1Re.FindStringSubmatch(body); m != nil {
		pb.Title = strings.TrimSpace(m[1])
	}

	h2Indexes := h2Re.FindAllStringSubmatchIndex(body, -1)
	for i, idx := range h2Indexes {
		start := idx[0]
		end := len(body)
		if i+1 < len(h2Indexes) {
			end = h2Indexes[i+1][0]
		}
		section := body[start:end]
		headerText := strings.TrimSpace(body[idx[2]:idx[3]])
		key, isCanonical := canonicalSections[strings.ToLower(headerText)]
		if isCanonical {
			pb.Sections[key] = section
		} else {
			pb.CustomSections[headerText] = section
		}
	}

	if opts, ok := pb.Sections["options"]; ok {
		for _, m := range bulletRe.FindAllStringSubmatch(opts, -1) {
			pb.Options = append(pb.Options, strings.TrimSpace(m[1]))
		}
	}

	if outcome, ok := pb.Sections["outcome"]; ok {
		if m := chosenRe.FindStringSubmatch(outcome); m != nil {
			pb.ChosenOption = m[1]
			pb.OutcomeRationale = strings.TrimSpace(m[2])
		}
	}

	return pb, nil
}
```

- [ ] **Step 4: Run, expect pass**

```bash
go test ./internal/domain/decision/madr/ -run TestParseBody -v
```

- [ ] **Step 5: Commit**

```bash
git add internal/domain/decision/madr/parser.go internal/domain/decision/madr/parser_test.go
git commit -m "feat(madr): ParseBody extracts sections, options, and chosen option"
```

---

## Task 5: Parser — frontmatter, filename, legacy detection

**Files:**
- Modify: `internal/domain/decision/madr/parser.go`
- Modify: `internal/domain/decision/madr/parser_test.go`

- [ ] **Step 1: Add failing tests**

```go
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
```

- [ ] **Step 2: Run, expect failure**.

- [ ] **Step 3: Append to `parser.go`**

```go
import (
	// ... existing ...
	"path/filepath"
	"gopkg.in/yaml.v3"
)

func ParseFrontmatter(text string) (Frontmatter, error) {
	if strings.TrimSpace(text) == "" {
		return Frontmatter{}, nil
	}
	var fm Frontmatter
	if err := yaml.Unmarshal([]byte(text), &fm); err != nil {
		return Frontmatter{}, fmt.Errorf("invalid frontmatter YAML: %w", err)
	}
	return fm, nil
}

var filenameRe = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)

func ParseFilename(path string) (id, slug string, err error) {
	base := filepath.Base(path)
	m := filenameRe.FindStringSubmatch(base)
	if m == nil {
		return "", "", fmt.Errorf("filename %q does not match NNNN-slug.md", base)
	}
	return m[1], m[2], nil
}

var (
	legacyADFilenameRe = regexp.MustCompile(`^AD\d{4}-.*\.md$`)
	legacyAnchorRe     = regexp.MustCompile(`<a name="(question|options|criteria|outcome|comments|comment-\d+|option-\d+)"></a>`)
	legacyStatusRe     = regexp.MustCompile(`(?m)^status:\s*(open|decided)\s*$`)
	legacyADRIDRe      = regexp.MustCompile(`(?m)^adr_id:\s*`)
)

// IsLegacyADG returns true if the file appears to use the pre-MADR ADG format.
// Read-side commands will use this to refuse legacy files and steer users to
// `adg migrate` (PR 4).
func IsLegacyADG(path string, content []byte) bool {
	if legacyADFilenameRe.MatchString(filepath.Base(path)) {
		return true
	}
	if legacyAnchorRe.Match(content) {
		return true
	}
	if legacyStatusRe.Match(content) {
		return true
	}
	if legacyADRIDRe.Match(content) {
		return true
	}
	return false
}
```

- [ ] **Step 4: Run, expect pass**

```bash
go test ./internal/domain/decision/madr/ -v
```

- [ ] **Step 5: Commit**

```bash
git add internal/domain/decision/madr/parser.go internal/domain/decision/madr/parser_test.go
git commit -m "feat(madr): ParseFrontmatter, ParseFilename, IsLegacyADG"
```

---

## Task 6: Renderer — canonical template for new ADRs

**Files:**
- Create: `internal/domain/decision/madr/renderer.go`
- Create: `internal/domain/decision/madr/renderer_test.go`

- [ ] **Step 1: Write failing test**

```go
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
```

- [ ] **Step 2: Run, expect failure**.

- [ ] **Step 3: Create `renderer.go`**

```go
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
// freshly-created ADR.
func RenderNewBody(title string) string {
	return fmt.Sprintf(canonicalTemplate, title)
}

// renderCommentsSection produces the trailing ## Comments H2 from a Comment list.
// Returns "" if the list is empty.
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
```

- [ ] **Step 4: Run, expect pass**

```bash
go test ./internal/domain/decision/madr/ -run TestRenderNewBody -v
```

- [ ] **Step 5: Commit**

```bash
git add internal/domain/decision/madr/renderer.go internal/domain/decision/madr/renderer_test.go
git commit -m "feat(madr): RenderNewBody emits canonical MADR template"
```

---

## Task 7: Renderer — full file (frontmatter + body + comments)

**Files:**
- Modify: `internal/domain/decision/madr/renderer.go`
- Modify: `internal/domain/decision/madr/renderer_test.go`

- [ ] **Step 1: Add failing tests**

```go
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
```

- [ ] **Step 2: Run, expect failure**.

- [ ] **Step 3: Append to `renderer.go`**

```go
import (
	// ... existing ...
	"bytes"
	"gopkg.in/yaml.v3"
)

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
```

- [ ] **Step 4: Run, expect pass**

```bash
go test ./internal/domain/decision/madr/ -run TestRenderFile -v
```

- [ ] **Step 5: Commit**

```bash
git add internal/domain/decision/madr/renderer.go internal/domain/decision/madr/renderer_test.go
git commit -m "feat(madr): RenderFile assembles frontmatter + body + regenerated comments"
```

---

## Task 8: Round-trip property test over fixture set

**Files:**
- Create: `internal/domain/decision/madr/roundtrip_test.go`

- [ ] **Step 1: Write the property test**

```go
package madr

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRoundTrip_AllFixtures is the load-bearing test:
//   parse(f) -> render -> f' such that diff(f, f') is empty.
// Any failure indicates a parser/renderer drift.
func TestRoundTrip_AllFixtures(t *testing.T) {
	fixtures, err := filepath.Glob("../../../../testdata/fixtures/madr/*.md")
	assert.NoError(t, err)
	assert.NotEmpty(t, fixtures, "no fixtures found at expected glob")

	for _, path := range fixtures {
		t.Run(filepath.Base(path), func(t *testing.T) {
			raw, err := os.ReadFile(path)
			assert.NoError(t, err)

			fmText, body, err := SplitFile(raw)
			assert.NoError(t, err)
			fm, err := ParseFrontmatter(fmText)
			assert.NoError(t, err)
			d := DecisionFromFrontmatter(fm)

			out, err := RenderFile(d, body)
			assert.NoError(t, err)

			assert.Equal(t,
				strings.TrimRight(string(raw), "\n"),
				strings.TrimRight(out, "\n"),
				"round-trip drift in %s", path,
			)
		})
	}
}
```

- [ ] **Step 2: Run; some fixtures will likely fail**

```bash
go test ./internal/domain/decision/madr/ -run TestRoundTrip -v
```

Likely failures and how to address them:
- **YAML key order drift.** `yaml.v3` marshals struct fields in declaration order; our `Frontmatter` struct's field order should match what the fixtures use. If a fixture has a different order, the round-trip will fail. Either reorder the struct fields, or normalize the fixture's order to match.
- **Bool default drift.** `informed: []` in a fixture vs. omitempty handling. Either populate or use omitempty consistently.
- **Quoted vs. unquoted scalars.** `status: "accepted"` (quoted) round-trips through yaml.v3 as `status: accepted` (unquoted) for plain strings. Adjust the fixture to use unquoted form OR add a custom marshaler. Recommendation: adjust fixtures to use unquoted form for plain strings; reserve quotes for values that need them.
- **Trailing newlines.** The test trims trailing newlines so this isn't an assertion issue, but the renderer's output should still be tidy (exactly one trailing `\n`).
- **`## Comments` section formatting.** If a fixture has `## Comments` content, the renderer regenerates it from frontmatter — must match exactly. Tune the renderer's `* **{date} — @{author}:** {text}\n` template to match fixture format.

For each failure: diff `raw` vs. `out`, fix renderer or fixture, re-run. Keep iterating until every fixture passes.

- [ ] **Step 3: Once all pass, commit**

```bash
git add internal/domain/decision/madr/roundtrip_test.go internal/domain/decision/madr/renderer.go testdata/fixtures/madr/
git commit -m "test(madr): round-trip property over fixture set; tune renderer to pass all"
```

---

## Task 9: Final build + test sweep + open PR

- [ ] **Step 1: Full clean build**

```bash
go build ./...
```

Expected: clean — no errors anywhere. PR 1a is purely additive; existing packages were not touched.

- [ ] **Step 2: Full test suite**

```bash
go test ./... 2>&1 | tail -20
```

Expected: all tests pass (existing tests untouched; new `madr` package tests pass).

- [ ] **Step 3: Push and open PR**

```bash
git checkout -b feat/pr1a-madr-parser-renderer
git push -u origin feat/pr1a-madr-parser-renderer
gh pr create --title "feat: MADR 4.0 parser, renderer, types (additive subpackage)" --body "$(cat <<'EOF'
## Summary

* Adds `internal/domain/decision/madr/` subpackage with MADR-shaped types (`Decision`, `Comment`, `Frontmatter`), markdown body parser, frontmatter parser, filename parser, legacy-format detector, and renderer.
* Adds fixture set under `testdata/fixtures/madr/` (4 upstream MADR templates + 5 synthetic fixtures).
* Adds the load-bearing round-trip property test: `parse(f) → render → f' ⟹ diff(f, f') == ∅` across all fixtures.

Purely additive — no existing files are modified or deleted. `go build ./...` and `go test ./...` both pass.

This is the first of four sub-PRs (1a–d) that together replace ADG's HTML-anchor file format with MADR 4.0. PR 1b switches the repository to the new types; 1c ports adapters/cmd; 1d rewrites `models/clean` and updates the README.

See `docs/superpowers/specs/2026-05-13-adg-madr-fork-design.md` for the full design and `docs/superpowers/plans/2026-05-13-adg-madr-pr1a-parser-renderer.md` for this PR's task breakdown.

## Test plan

- [ ] `go test ./internal/domain/decision/madr/...` passes
- [ ] Round-trip property test passes for all fixtures
- [ ] `go test ./...` (full suite) passes — existing tests unaffected

🤖 Generated with [Claude Code](https://claude.com/claude-code)
EOF
)"
```

---

## Self-Review

1. **Spec coverage (PR 1a slice):** MADR types ✓ (Task 2), body parser ✓ (Tasks 4-5), frontmatter parser ✓ (Task 5), filename + legacy detection ✓ (Task 5), canonical template renderer ✓ (Task 6), full file renderer with comments regeneration ✓ (Task 7), round-trip property test ✓ (Task 8), fixture set ✓ (Task 1). Repository, service, validator, adapters, cmd, models/clean, README are all explicit PR 1b/1c/1d scope per the spec.

2. **Placeholder scan:** Task 8 step 2 contains a "likely failures and how to address them" prose rather than concrete code. This is intentional — round-trip tuning is empirical (depends on yaml.v3 ordering quirks at runtime) and the engineer adapts based on actual diffs. Acceptable departure from "complete code in every step" given the iterative nature of that specific step.

3. **Type consistency:** `Decision`, `Comment`, `Frontmatter`, `ParsedBody`, `SplitFile`, `ParseBody`, `ParseFrontmatter`, `ParseFilename`, `IsLegacyADG`, `RenderNewBody`, `RenderFile` — used consistently across all tasks. The package qualifier `madr.X` will appear in PR 1b consumers.

4. **Scope check:** 9 tasks, ~30-45 minutes each. PR 1a stays additive. No premature integration work.

5. **Ambiguity check:** The two pre-flight ratifications are documented; nothing else is ambiguous in this slice.
