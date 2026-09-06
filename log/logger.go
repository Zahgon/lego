package log

import (
	"log/slog"
	"os"
	"sync/atomic"
)

var defaultLogger atomic.Pointer[slog.Logger]

func init() {
	defaultLogger.Store(slog.New(slog.NewTextHandler(os.Stdout, nil)))
}

func Default() *slog.Logger { _ = "STUB: not implemented"; return nil }

func SetDefault(l *slog.Logger) { _ = "STUB: not implemented"; return }

func Fatal(msg string, args ...slog.Attr) { _ = "STUB: not implemented"; return }

func Debug(msg string, args ...slog.Attr) { _ = "STUB: not implemented"; return }

func Info(msg string, args ...slog.Attr) { _ = "STUB: not implemented"; return }

func Warn(msg string, args ...slog.Attr) { _ = "STUB: not implemented"; return }

func Error(msg string, args ...slog.Attr) { _ = "STUB: not implemented"; return }
