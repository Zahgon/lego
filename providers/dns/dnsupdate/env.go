package dnsupdate

import (
	"time"
)

const (
	altEnvRFC2136Namespace    = "RFC2136_"
	altEnvRFC3645SubNamespace = "RFC3645_"
)

func altEnvNames(v string) []string { _ = "STUB: not implemented"; return nil }

func getEnvString(name string) string { _ = "STUB: not implemented"; return "" }

func getEnvStringSlice(name string) []string { _ = "STUB: not implemented"; return nil }

func getOrDefaultString(name, defaultValue string) string { _ = "STUB: not implemented"; return "" }

func getOrDefaultSecond(name string, defaultValue time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func getOrDefaultInt(name string, defaultValue int) int { _ = "STUB: not implemented"; return 0 }

func getOneWithFallback[T any](main string, defaultValue T, fn func(string) (T, error)) T {
	_ = "STUB: not implemented"
	return *new(T)
}
