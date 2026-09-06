package env

import (
	"time"
)

func Get(names ...string) (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

func GetWithFallback(groups ...[]string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetOneWithFallback[T any](main string, defaultValue T, fn func(string) (T, error), names ...string) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func getOneWithFallback(main string, names ...string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func GetOrDefaultString(envVar, defaultValue string) string { _ = "STUB: not implemented"; return "" }

func GetOrDefaultBool(envVar string, defaultValue bool) bool {
	_ = "STUB: not implemented"
	return false
}

func GetOrDefaultInt(envVar string, defaultValue int) int { _ = "STUB: not implemented"; return 0 }

func GetOrDefaultSecond(envVar string, defaultValue time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func getOrDefault[T any](envVar string, defaultValue T, fn func(string) (T, error)) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func GetOrFile(envVar string) string { _ = "STUB: not implemented"; return "" }

func ParseSecond(s string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func ParseString(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ParsePairs(raw string) (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }
