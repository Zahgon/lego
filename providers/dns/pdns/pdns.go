package pdns

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/pdns/internal"
)

const (
	envNamespace = "PDNS_"

	EnvAPIKey = envNamespace + "API_KEY"
	EnvAPIURL = envNamespace + "API_URL"

	EnvTTL                = envNamespace + "TTL"
	EnvAPIVersion         = envNamespace + "API_VERSION"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
	EnvServerName         = envNamespace + "SERVER_NAME"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	APIKey             string
	Host               *url.URL
	ServerName         string
	APIVersion         int
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *internal.Client
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

func findTxtRecord(zone *internal.HostedZone, fqdn string) *internal.RRSet {
	_ = "STUB: not implemented"
	return nil
}
