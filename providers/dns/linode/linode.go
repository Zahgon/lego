package linode

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/linode/linodego"
)

const (
	envNamespace = "LINODE_"

	EnvToken = envNamespace + "TOKEN"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const (
	minTTL             = 300
	dnsUpdateFreqMins  = 15
	dnsUpdateFudgeSecs = 120
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	Token              string
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPTimeout        time.Duration
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type hostedZoneInfo struct {
	domainID     int
	resourceName string
}

type DNSProvider struct {
	config *Config
	client *linodego.Client
}

func NewDNSProvider() (*DNSProvider, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) Timeout() (time.Duration, time.Duration) {
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

func (d *DNSProvider) getHostedZoneInfo(ctx context.Context, fqdn string) (*hostedZoneInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
