package log

import (
	"fmt"
	"log/slog"
)

type LazyMessage struct {
	msg  string
	args []any
}

func LazySprintf(msg string, args ...any) LazyMessage {
	_ = "STUB: not implemented"
	return *new(LazyMessage)
}

func (l LazyMessage) String() string { _ = "STUB: not implemented"; return "" }

func Debugf(msg fmt.Stringer, args ...slog.Attr) { _ = "STUB: not implemented"; return }

func Infof(msg fmt.Stringer, args ...slog.Attr) { _ = "STUB: not implemented"; return }

func Warnf(msg fmt.Stringer, args ...slog.Attr) { _ = "STUB: not implemented"; return }

func Errorf(msg fmt.Stringer, args ...slog.Attr) { _ = "STUB: not implemented"; return }

func logLazy(level slog.Level, msg fmt.Stringer, args ...slog.Attr) {
	_ = "STUB: not implemented"
	return
}
