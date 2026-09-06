package log

import (
	"log/slog"
	"time"
)

func ErrorAttr(err error) slog.Attr { _ = "STUB: not implemented"; return *new(slog.Attr) }

func DomainAttr(v string) slog.Attr { _ = "STUB: not implemented"; return *new(slog.Attr) }

func DomainsAttr(v []string) slog.Attr { _ = "STUB: not implemented"; return *new(slog.Attr) }

func CertNameAttr(v string) slog.Attr { _ = "STUB: not implemented"; return *new(slog.Attr) }

func DurationAttr(key string, v time.Duration) slog.Attr {
	_ = "STUB: not implemented"
	return *new(slog.Attr)
}
