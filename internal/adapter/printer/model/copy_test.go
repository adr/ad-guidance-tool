package model

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestCopyModelPresenter_Copied(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	presenter := NewCopyPresenter()
	presenter.Copied("source-model", "target-model", 3)

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = oldStdout

	expected := "Successfully copied 3 decisions from model source-model to new model target-model\n"
	actual := buf.String()
	if actual != expected {
		t.Errorf("expected output %q, got %q", expected, actual)
	}
}
