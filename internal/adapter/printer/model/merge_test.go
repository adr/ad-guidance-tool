package model

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMergeModelsPresenter_Merged(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	p := NewMergePresenter()
	err := p.Merged("modelA", "modelB", "target", 5)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = oldStdout

	expected := "Successfully merged 5 decisions from models modelA and modelB to new directory: target\n"
	if got := buf.String(); got != expected {
		t.Errorf("expected output %q, got %q", expected, got)
	}
}
