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
