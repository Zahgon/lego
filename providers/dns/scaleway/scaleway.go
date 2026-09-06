package scaleway

import (
	"context"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	scwdomain "github.com/scaleway/scaleway-sdk-go/api/domain/v2beta1"
)

const (
	envNamespace = "SCALEWAY_"

	EnvAPIToken  = envNamespace + "API_TOKEN"
	EnvProjectID = envNamespace + "PROJECT_ID"

	altEnvNamespace = "SCW_"

	EnvAccessKey = altEnvNamespace + "ACCESS_KEY"
	EnvSecretKey = altEnvNamespace + "SECRET_KEY"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const (
	minTTL                    = 60
	defaultPollingInterval    = 10 * time.Second
	defaultPropagationTimeout = 120 * time.Second
)

const dumpAccessKey = "SCWXXXXXXXXXXXXXXXXX"

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	ProjectID string
	SecretKey string
	AccessKey string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *scwdomain.API
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

func altEnvName(v string) string { _ = "STUB: not implemented"; return "" }
