package bunny

import (
	"context"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/nrdcg/bunny-go"
)

const (
	envNamespace = "BUNNY_"

	EnvAPIKey = envNamespace + "API_KEY"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const minTTL = 60

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	APIKey string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *bunny.Client
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

func (d *DNSProvider) findZone(ctx context.Context, authZone string) (*bunny.DNSZone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findZone(zones *bunny.DNSZones, domain string) *bunny.DNSZone {
	_ = "STUB: not implemented"
	return nil
}

func possibleDomains(domain string) []string { _ = "STUB: not implemented"; return nil }
