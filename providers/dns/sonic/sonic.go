package sonic

import (
	"context"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/sonic/internal"
)

const (
	envNamespace = "SONIC_"

	EnvUserID = envNamespace + "USER_ID"
	EnvAPIKey = envNamespace + "API_KEY"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvSequenceInterval   = envNamespace + "SEQUENCE_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	UserID             string
	APIKey             string
	HTTPClient         *http.Client
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	SequenceInterval   time.Duration
	TTL                int
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

func (d *DNSProvider) Sequential() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
