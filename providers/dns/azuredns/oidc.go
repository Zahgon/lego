package azuredns

import (
	"context"
)

func checkOIDCConfig(config *Config) error { _ = "STUB: not implemented"; return nil }

func getOIDCAssertion(config *Config) func(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return nil
}

func getOIDCToken(ctx context.Context, config *Config) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
