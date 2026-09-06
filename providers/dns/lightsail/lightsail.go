package lightsail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/go-acme/lego/v5/challenge"
)

const (
	envNamespace = "LIGHTSAIL_"

	EnvRegion  = envNamespace + "REGION"
	EnvDNSZone = "DNS_ZONE"

	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

const maxRetries = 5

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	DNSZone            string
	Region             string
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	client *lightsail.Client
	config *Config
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
