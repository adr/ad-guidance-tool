package ade

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"sync"
)

// cliHandler is a slog.Handler that writes human-readable CLI output.
// It is the single output standard for all ADE diagnostic and result messages.
// Line format: LEVEL  message  [key=val ...]
// The level label is padded to 6 characters so all lines are consistently aligned.
type cliHandler struct {
	mu       sync.Mutex
	w        io.Writer
	level    slog.Level
	skipWarn bool
	attrs    []slog.Attr
}

func newCLIHandler(w io.Writer, level slog.Level, skipWarn bool) *cliHandler {
	return &cliHandler{w: w, level: level, skipWarn: skipWarn}
}

func (h *cliHandler) Enabled(_ context.Context, l slog.Level) bool {
	if l == slog.LevelWarn && h.skipWarn {
		return false
	}
	return l >= h.level
}

func (h *cliHandler) Handle(_ context.Context, r slog.Record) error {
	const levelWidth = 6
	prefix := fmt.Sprintf("%-*s", levelWidth, r.Level.String())
	indent := strings.Repeat(" ", levelWidth)

	var attrs []slog.Attr
	attrs = append(attrs, h.attrs...)
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a)
		return true
	})

	// Separate single-line and multi-line attribute values.
	var inline []string
	var extra []string
	for _, a := range attrs {
		val := a.Value.String()
		if strings.Contains(val, "\n") {
			extra = append(extra, indent+a.Key+":")
			for _, line := range strings.Split(strings.TrimRight(val, "\n"), "\n") {
				extra = append(extra, indent+"  "+line)
			}
		} else {
			inline = append(inline, fmt.Sprintf("%s=%s", a.Key, val))
		}
	}

	var sb strings.Builder
	sb.WriteString(prefix)
	sb.WriteString(r.Message)
	if len(inline) > 0 {
		sb.WriteString("  ")
		sb.WriteString(strings.Join(inline, " "))
	}
	for _, line := range extra {
		sb.WriteByte('\n')
		sb.WriteString(line)
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	fmt.Fprintln(h.w, sb.String())
	return nil
}

func (h *cliHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	cp := *h
	cp.attrs = make([]slog.Attr, len(h.attrs)+len(attrs))
	copy(cp.attrs, h.attrs)
	copy(cp.attrs[len(h.attrs):], attrs)
	return &cp
}

func (h *cliHandler) WithGroup(_ string) slog.Handler { return h }
