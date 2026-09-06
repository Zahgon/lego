package constellix

import (
	"context"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/constellix/internal"
)

const (
	envNamespace = "CONSTELLIX_"

	EnvAPIKey    = envNamespace + "API_KEY"
	EnvSecretKey = envNamespace + "SECRET_KEY"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	APIKey             string
	SecretKey          string
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

func (d *DNSProvider) createRecord(ctx context.Context, dom internal.Domain, fqdn, recordName, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) appendRecordValue(ctx context.Context, dom internal.Domain, recordID int64, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) removeRecordValue(ctx context.Context, dom internal.Domain, record *internal.Record, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func containsValue(record *internal.Record, value string) bool {
	_ = "STUB: not implemented"
	return false
}

func backoff(minimum, maximum time.Duration, attemptNum int, resp *http.Response) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
