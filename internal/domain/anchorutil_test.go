package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnchorForSection(t *testing.T) {
	assert.Equal(t, `<a name="question"></a>`, AnchorForSection(AnchorSectionQuestion))
	assert.Equal(t, `<a name="options"></a>`, AnchorForSection(AnchorSectionOptions))
	assert.Equal(t, `<a name="criteria"></a>`, AnchorForSection(AnchorSectionCriteria))
	assert.Equal(t, `<a name="outcome"></a>`, AnchorForSection(AnchorSectionOutcome))
	assert.Equal(t, `<a name="comments"></a>`, AnchorForSection(AnchorSectionComments))
}

func TestAnchorLinkToSection(t *testing.T) {
	assert.Equal(t, `[Intro](#question)`, AnchorLinkToSection(AnchorSectionQuestion, "Intro"))
	assert.Equal(t, `[Choices](#options)`, AnchorLinkToSection(AnchorSectionOptions, "Choices"))
}

func TestAnchorForOption(t *testing.T) {
	assert.Equal(t, `<a name="option-1"></a>`, AnchorForOption(1))
	assert.Equal(t, `<a name="option-5"></a>`, AnchorForOption(5))
}

func TestAnchorLinkToOption(t *testing.T) {
	assert.Equal(t, `[Option 1](#option-1)`, AnchorLinkToOption(1))
	assert.Equal(t, `[Option 3](#option-3)`, AnchorLinkToOption(3))
}

func TestAnchorForComment(t *testing.T) {
	expected := `<a name="comment-2"></a>2. (2024-05-01) Alice: Looks good`
	result := AnchorForComment(2, "Alice", "2024-05-01", "Looks good")
	assert.Equal(t, expected, result)
}

func TestAnchorLinkToComment(t *testing.T) {
	assert.Equal(t, `[Comment 2](#comment-2)`, AnchorLinkToComment(2))
}
