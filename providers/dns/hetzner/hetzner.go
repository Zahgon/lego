package hetzner

import (
	"context"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/hetzner/internal/hetznerv1"
	"github.com/go-acme/lego/v5/providers/dns/hetzner/internal/legacy"
)

const (
	EnvAPIKey   = legacy.EnvAPIKey
	EnvAPIToken = hetznerv1.EnvAPIToken

	EnvTTL                = hetznerv1.EnvTTL
	EnvPropagationTimeout = hetznerv1.EnvPropagationTimeout
	EnvPollingInterval    = hetznerv1.EnvPollingInterval
	EnvHTTPTimeout        = hetznerv1.EnvHTTPTimeout
)

const minTTL = 60

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	APIKey string

	APIToken string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	provider challenge.ProviderTimeout
}

func NewDNSProvider() (*DNSProvider, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) Timeout() (timeout, interval time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}

func (d *DNSProvider) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}
