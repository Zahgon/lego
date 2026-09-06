package cloudflare

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
)

const (
	envNamespace = "CLOUDFLARE_"

	EnvEmail  = envNamespace + "EMAIL"
	EnvAPIKey = envNamespace + "API_KEY"

	EnvDNSAPIToken  = envNamespace + "DNS_API_TOKEN"
	EnvZoneAPIToken = envNamespace + "ZONE_API_TOKEN"

	EnvBaseURL = envNamespace + "BASE_URL"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const (
	altEnvNamespace = "CF_"

	altEnvEmail = altEnvNamespace + "API_EMAIL"
)

const (
	minTTL = 120
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	AuthEmail string
	AuthKey   string

	AuthToken string
	ZoneToken string

	BaseURL string

	TTL                int
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	client *metaClient
	config *Config

	recordIDs   map[string]string
	recordIDsMu sync.Mutex
}

func NewDNSProvider() (*DNSProvider, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:errorlint

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

func altEnvName(v string) string { _ = "STUB: not implemented"; return "" }
