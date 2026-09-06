package transip

import (
	"context"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	transipdomain "github.com/transip/gotransip/v6/domain"
)

const (
	envNamespace = "TRANSIP_"

	EnvAccountName    = envNamespace + "ACCOUNT_NAME"
	EnvPrivateKeyPath = envNamespace + "PRIVATE_KEY_PATH"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	AccountName        string
	PrivateKeyPath     string
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int64
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config     *Config
	repository transipdomain.Repository
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
