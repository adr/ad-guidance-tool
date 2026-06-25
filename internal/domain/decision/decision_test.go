package decision

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestLinks_MarshalYAML(t *testing.T) {
	links := Links{
		Precedes: []string{"A", "B"},
		Succeeds: []string{"C"},
		Custom: map[string][]string{
			"related": {"X", "Y"},
		},
	}

	out, err := links.MarshalYAML()
	assert.NoError(t, err)

	asMap, ok := out.(map[string]any)
	assert.True(t, ok)

	assert.ElementsMatch(t, []string{"A", "B"}, asMap["precedes"])
	assert.ElementsMatch(t, []string{"C"}, asMap["succeeds"])
	assert.ElementsMatch(t, []string{"X", "Y"}, asMap["related"])
}

func TestLinks_UnmarshalYAML(t *testing.T) {
	yml := `
precedes:
  - A
  - B
succeeds:
  - C
custom1:
  - X
custom2:
  - Y
`

	var links Links
	err := yaml.Unmarshal([]byte(yml), &links)
	assert.NoError(t, err)

	assert.Equal(t, []string{"A", "B"}, links.Precedes)
	assert.Equal(t, []string{"C"}, links.Succeeds)
	assert.Equal(t, map[string][]string{
		"custom1": {"X"},
		"custom2": {"Y"},
	}, links.Custom)
}

func TestLinks_UnmarshalYAML_InvalidFormat(t *testing.T) {
	invalidYAML := `
- not
- a
- map
`

	var links Links
	err := yaml.Unmarshal([]byte(invalidYAML), &links)
	assert.Error(t, err)
}

func TestAsStringSlice_WithInterfaceSlice(t *testing.T) {
	input := []any{"a", 1, true}
	expected := []string{"a", "1", "true"}
	result := asStringSlice(input)
	assert.Equal(t, expected, result)
}

func TestAsStringSlice_WithStringSlice(t *testing.T) {
	input := []string{"x", "y"}
	assert.Equal(t, input, asStringSlice(input))
}

func TestAsStringSlice_WithNil(t *testing.T) {
	assert.Empty(t, asStringSlice(nil))
}

func TestAsStringSlice_WithUnsupportedType(t *testing.T) {
	assert.Empty(t, asStringSlice(42))
}

func TestDecision_MarshalYAML_OmitsEmptyFields(t *testing.T) {
	decision := Decision{
		ID:       "0001",
		Title:    "Test Decision",
		Status:   "accepted",
		Tags:     []string{},
		Links:    Links{},
		Comments: []Comment{},
	}

	data, err := yaml.Marshal(&decision)
	assert.NoError(t, err)

	yamlStr := string(data)
	assert.NotContains(t, yamlStr, "tags:")
	assert.NotContains(t, yamlStr, "links:")
	assert.NotContains(t, yamlStr, "comments:")
	assert.Contains(t, yamlStr, "adr_id: \"0001\"")
	assert.Contains(t, yamlStr, "title: Test Decision")
	assert.Contains(t, yamlStr, "status: accepted")
}

func TestDecision_MarshalYAML_IncludesNonEmptyFields(t *testing.T) {
	decision := Decision{
		ID:     "0002",
		Title:  "Another Decision",
		Status: "proposed",
		Tags:   []string{"urgent", "architecture"},
		Links: Links{
			Precedes: []string{"0001"},
		},
		Comments: []Comment{
			{Author: "Alice", Date: "2024-01-01", Comment: "Good idea"},
		},
	}

	data, err := yaml.Marshal(&decision)
	assert.NoError(t, err)

	yamlStr := string(data)
	assert.Contains(t, yamlStr, "tags:")
	assert.Contains(t, yamlStr, "links:")
	assert.Contains(t, yamlStr, "comments:")
	assert.Contains(t, yamlStr, "urgent")
	assert.Contains(t, yamlStr, "architecture")
	assert.Contains(t, yamlStr, "precedes:")
	assert.Contains(t, yamlStr, "Alice")
}

func TestLinks_MarshalYAML_OmitsEmptyArrays(t *testing.T) {
	links := Links{
		Precedes: []string{},
		Succeeds: []string{},
		Custom: map[string][]string{
			"empty":    {},
			"nonempty": {"X"},
		},
	}

	out, err := links.MarshalYAML()
	assert.NoError(t, err)

	asMap, ok := out.(map[string]any)
	assert.True(t, ok)

	assert.NotContains(t, asMap, "precedes")
	assert.NotContains(t, asMap, "succeeds")
	assert.NotContains(t, asMap, "empty")
	assert.Contains(t, asMap, "nonempty")
	assert.ElementsMatch(t, []string{"X"}, asMap["nonempty"])
}
