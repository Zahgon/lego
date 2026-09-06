package log

import (
	"log/slog"
	"time"
)

type FormattableDuration time.Duration

func (f FormattableDuration) String() string { _ = "STUB: not implemented"; return "" }

func (f FormattableDuration) LogValue() slog.Value {
	_ = "STUB: not implemented"
	return *new(slog.Value)
}
