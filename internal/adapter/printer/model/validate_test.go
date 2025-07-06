package model

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = old

	return buf.String()
}

func TestModelValidatePresenter_ModelValidated_AllValid(t *testing.T) {
	presenter := NewModelValidatePresenter()

	output := captureOutput(func() {
		presenter.ModelValidated("test-model", nil, nil)
	})

	expected := "test-model model metadata is valid and index is up to date\n" +
		"test-model model file content is valid with correct anchors\n"

	if output != expected {
		t.Errorf("unexpected output:\nexpected:\n%s\ngot:\n%s", expected, output)
	}
}

func TestModelValidatePresenter_ModelValidated_InvalidIndex(t *testing.T) {
	presenter := NewModelValidatePresenter()

	output := captureOutput(func() {
		presenter.ModelValidated("test-model", errors.New("index missing keys"), nil)
	})

	expected := "test-model model metadata is invalid: index missing keys\n"

	if output != expected {
		t.Errorf("unexpected output:\nexpected:\n%s\ngot:\n%s", expected, output)
	}
}

func TestModelValidatePresenter_ModelValidated_InvalidData(t *testing.T) {
	presenter := NewModelValidatePresenter()

	output := captureOutput(func() {
		presenter.ModelValidated("test-model", nil, errors.New("missing anchor in decision content"))
	})

	expected := "test-model model metadata is valid and index is up to date\n" +
		"test-model model file content is invalid: missing anchor in decision content\n"

	if output != expected {
		t.Errorf("unexpected output:\nexpected:\n%s\ngot:\n%s", expected, output)
	}
}
