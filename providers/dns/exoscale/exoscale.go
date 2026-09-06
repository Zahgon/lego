package exoscale

import (
	"context"
	"time"

	egoscale "github.com/exoscale/egoscale/v3"
	"github.com/go-acme/lego/v5/challenge"
)

const (
	envNamespace = "EXOSCALE_"

	EnvAPISecret = envNamespace + "API_SECRET"
	EnvAPIKey    = envNamespace + "API_KEY"
	EnvEndpoint  = envNamespace + "ENDPOINT"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	APIKey             string
	APISecret          string
	Endpoint           string
	HTTPTimeout        time.Duration
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int64
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *egoscale.Client
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

func (d *DNSProvider) findExistingZone(ctx context.Context, zoneName string) (*egoscale.DNSDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) findExistingRecordID(ctx context.Context, zoneID egoscale.UUID, recordName, value string) (egoscale.UUID, error) {
	_ = "STUB: not implemented"
	return *new(egoscale.UUID), nil
}

func (d *DNSProvider) findZoneAndRecordName(ctx context.Context, fqdn string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
