package stackpath

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/challenge"
)

const (
	envNamespace = "STACKPATH_"

	EnvClientID     = envNamespace + "CLIENT_ID"
	EnvClientSecret = envNamespace + "CLIENT_SECRET"
	EnvStackID      = envNamespace + "STACK_ID"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	ClientID           string
	ClientSecret       string
	StackID            string
	TTL                int
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct{}

func NewDNSProvider() (*DNSProvider, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) Timeout() (timeout, interval time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}
