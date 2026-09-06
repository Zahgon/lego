package joker

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/joker/internal/dmapi"
)

var _ challenge.ProviderTimeout = (*dmapiProvider)(nil)

type dmapiProvider struct {
	config *Config
	client *dmapi.Client
}

func newDmapiProvider() (*dmapiProvider, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:errorlint // false-positive

func newDmapiProviderConfig(config *Config) (*dmapiProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dmapiProvider) Timeout() (timeout, interval time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}

func (d *dmapiProvider) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *dmapiProvider) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func formatResponseError(response *dmapi.Response, err error) error {
	_ = "STUB: not implemented"
	return nil
}
