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
	bodyBytes := rest[bodyStart:]
	// Strip one optional leading blank line between frontmatter close and body.
	// The renderer always emits this blank line; consuming it here makes the
	// "body" string canonical regardless of whether frontmatter was present.
	if len(bodyBytes) > 0 && bodyBytes[0] == '\n' {
		bodyBytes = bodyBytes[1:]
	}
	return fm, string(bodyBytes), nil
}
