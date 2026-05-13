package madr

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
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
	bodyBytes := rest[bodyStart:]
	// Strip one optional leading blank line between frontmatter close and body.
	// The renderer always emits this blank line; consuming it here makes the
	// "body" string canonical regardless of whether frontmatter was present.
	if len(bodyBytes) > 0 && bodyBytes[0] == '\n' {
		bodyBytes = bodyBytes[1:]
	}
	return fm, string(bodyBytes), nil
}

// ParsedBody is the result of ParseBody — everything we extract from a body.
type ParsedBody struct {
	Title            string
	Sections         map[string]string // canonical key -> raw section text (incl. H2 line)
	Options          []string          // bullet items under Considered Options, in order
	ChosenOption     string            // text from `Chosen option: "..."`
	OutcomeRationale string            // text after `because ` and before the trailing `.`
	CustomSections   map[string]string // unrecognized H2 header text -> raw section text
}

// canonicalSections maps lowercased H2 header text to a short key.
// We use exact equality (case-insensitive) on header text — NOT contains-style —
// so a header like "Considered Trade-offs" is treated as custom, not as options.
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
	h1Re     = regexp.MustCompile(`(?m)^# +(.+)$`)
	h2Re     = regexp.MustCompile(`(?m)^## +(.+?)\s*$`)
	bulletRe = regexp.MustCompile(`(?m)^\s*\*\s+(.+)$`)
	chosenRe = regexp.MustCompile(`(?m)^Chosen option:\s*"([^"]*)"(?:\s*,\s*because\s+(.+?))?\.?\s*$`)
)

// ParseBody extracts the H1 title, recognized canonical sections, options
// bullets, and Decision Outcome's chosen option from a MADR-shaped body.
// Unknown H2 headers are preserved in CustomSections so the renderer can
// reproduce them verbatim.
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
