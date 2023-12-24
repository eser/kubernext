package log

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
)

type DebugHandlerOptions struct {
	SlogOptions slog.HandlerOptions
}

type DebugHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *DebugHandler) Handle(ctx context.Context, record slog.Record) error {
	var level string

	switch record.Level {
	case slog.LevelError:
		level = "ERROR: "
	case slog.LevelWarn:
		level = "WARN: "
	case slog.LevelInfo:
		level = "INFO: "
	case slog.LevelDebug:
		level = "DEBUG: "
	}

	fields := make(map[string]interface{}, record.NumAttrs())
	record.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	bytes, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return fmt.Errorf("DebugHandler error: %w", err)
	}

	timeStr := record.Time.Format("[15:05:05.000]")
	msg := record.Message
	details := string(bytes)

	h.l.Println(timeStr, level, msg, details)

	return nil
}

func NewDebugHandler(
	out io.Writer,
	opts DebugHandlerOptions,
) *DebugHandler {
	h := &DebugHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOptions),
		l:       log.New(out, "", 0),
	}

	return h
}
