package mythicbeasts

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/mythicbeasts/internal"
)

const (
	envNamespace = "MYTHICBEASTS_"

	EnvUserName        = envNamespace + "USERNAME"
	EnvPassword        = envNamespace + "PASSWORD"
	EnvAPIEndpoint     = envNamespace + "API_ENDPOINT"
	EnvAuthAPIEndpoint = envNamespace + "AUTH_API_ENDPOINT"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	UserName           string
	Password           string
	HTTPClient         *http.Client
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	APIEndpoint        *url.URL
	AuthAPIEndpoint    *url.URL
	TTL                int
}

func NewDefaultConfig() (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

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
