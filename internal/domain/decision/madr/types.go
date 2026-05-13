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
