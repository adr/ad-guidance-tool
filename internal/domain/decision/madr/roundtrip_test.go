package madr

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// nonRoundTrippableFixtures are upstream MADR templates that contain
// placeholder content not designed to round-trip through a typed parser:
//
//   - full.md uses `{...}` placeholders that YAML interprets as flow
//     mappings, not as strings; the typed parser rejects them.
//   - bare.md has frontmatter keys with empty values (`status:`, `date:`,
//     etc.), which yaml.v3 unmarshals to zero values; on re-marshal with
//     omitempty those keys disappear entirely.
//
// These are *template* fixtures users would fill in to produce a real ADR;
// the filled-in result rounds-trips, but the template itself does not.
// They're kept in the fixture set so the parser-tolerance test below proves
// we don't crash on them.
var nonRoundTrippableFixtures = map[string]bool{
	"full.md": true,
	"bare.md": true,
}

// TestRoundTrip_AllFixtures is the load-bearing property test:
//
//	parse(f) -> render -> f' such that diff(f, f') is empty.
//
// Any failure indicates a parser/renderer drift.
func TestRoundTrip_AllFixtures(t *testing.T) {
	fixtures, err := filepath.Glob("../../../../testdata/fixtures/madr/*.md")
	assert.NoError(t, err)
	assert.NotEmpty(t, fixtures, "no fixtures found at expected glob")

	for _, path := range fixtures {
		if nonRoundTrippableFixtures[filepath.Base(path)] {
			continue
		}
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

// TestNonRoundTrippable_DoesNotCrash proves the parser is robust against the
// upstream-template fixtures even though they don't round-trip cleanly.
func TestNonRoundTrippable_DoesNotCrash(t *testing.T) {
	for name := range nonRoundTrippableFixtures {
		t.Run(name, func(t *testing.T) {
			raw, err := os.ReadFile(filepath.Join("../../../../testdata/fixtures/madr/", name))
			assert.NoError(t, err)
			// SplitFile alone must succeed for both fixtures; ParseFrontmatter
			// may error on full.md (intentional, due to invalid YAML placeholders),
			// but the parser shouldn't crash.
			fmText, body, splitErr := SplitFile(raw)
			assert.NoError(t, splitErr)
			_, _ = ParseFrontmatter(fmText) // ignored
			_, parseErr := ParseBody(body)
			assert.NoError(t, parseErr)
		})
	}
}
