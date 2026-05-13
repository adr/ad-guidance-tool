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
