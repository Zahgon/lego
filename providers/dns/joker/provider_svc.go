package joker

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/joker/internal/svc"
)

var _ challenge.ProviderTimeout = (*svcProvider)(nil)

type svcProvider struct {
	config *Config
	client *svc.Client
}

func newSvcProvider() (*svcProvider, error) { _ = "STUB: not implemented"; return nil, nil }

func newSvcProviderConfig(config *Config) (*svcProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *svcProvider) Timeout() (timeout, interval time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}

func (d *svcProvider) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *svcProvider) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *svcProvider) Sequential() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
