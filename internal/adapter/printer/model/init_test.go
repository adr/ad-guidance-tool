package model

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestInitModelPresenter_Initialized(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	p := NewInitPresenter()
	p.Initialized("test/model")

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = oldStdout

	expected := "Successfully created model directory: test/model\n"
	if got := buf.String(); got != expected {
		t.Errorf("expected output %q, got %q", expected, got)
	}
}
