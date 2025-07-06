package model

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestImportModelPresenter_Imported(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	p := NewImportPresenter()
	err := p.Imported("source-model", "target-model", 5)

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = oldStdout

	expectedOutput := "Successfully imported model source-model with 5 decisions to: target-model\n"
	if got := buf.String(); got != expectedOutput {
		t.Errorf("unexpected output:\nexpected: %q\ngot: %q", expectedOutput, got)
	}

	if err != nil {
		t.Errorf("expected nil error, got: %v", err)
	}
}
