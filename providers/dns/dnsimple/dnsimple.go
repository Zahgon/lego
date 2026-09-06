package dnsimple

import (
	"context"
	"time"

	"github.com/dnsimple/dnsimple-go/v9/dnsimple"
	"github.com/go-acme/lego/v5/challenge"
)

const (
	envNamespace = "DNSIMPLE_"

	EnvOAuthToken = envNamespace + "OAUTH_TOKEN"
	EnvBaseURL    = envNamespace + "BASE_URL"
	EnvDebug      = envNamespace + "DEBUG"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	Debug              bool
	AccessToken        string
	BaseURL            string
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *dnsimple.Client
}

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

func (d *DNSProvider) getHostedZone(ctx context.Context, domain string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *DNSProvider) findTxtRecords(ctx context.Context, fqdn string) ([]dnsimple.ZoneRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTxtRecord(zoneName, fqdn, value string, ttl int) (dnsimple.ZoneRecordAttributes, error) {
	_ = "STUB: not implemented"
	return *new(dnsimple.ZoneRecordAttributes), nil
}

func (d *DNSProvider) getAccountID(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
