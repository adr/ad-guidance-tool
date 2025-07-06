package model

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestRebuildIndexPresenter_IndexRebuilt(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	p := NewRebuildIndexPresenter()
	p.IndexRebuilt("TestModel")

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = oldStdout

	expected := "Index successfully updated for model: TestModel\n"
	if got := buf.String(); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
