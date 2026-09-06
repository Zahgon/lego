package nicru

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/nicru/internal"
)

const (
	envNamespace = "NICRU_"

	EnvUsername  = envNamespace + "USER"
	EnvPassword  = envNamespace + "PASSWORD"
	EnvServiceID = envNamespace + "SERVICE_ID"
	EnvSecret    = envNamespace + "SECRET"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	TTL                int
	Username           string
	Password           string
	ServiceID          string
	Secret             string
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config

	lazyClient func() (*internal.Client, error)
}

func NewDNSProvider() (*DNSProvider, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) Present(ctx context.Context, domain, _, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) CleanUp(ctx context.Context, domain, _, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) Timeout() (timeout, interval time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}

func findZone(ctx context.Context, client *internal.Client, authZone string) (*internal.Zone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validate(config *Config) error { _ = "STUB: not implemented"; return nil }
