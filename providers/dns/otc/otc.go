package otc

import (
	"context"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/otc/internal"
)

const (
	envNamespace = "OTC_"

	EnvDomainName       = envNamespace + "DOMAIN_NAME"
	EnvUserName         = envNamespace + "USER_NAME"
	EnvPassword         = envNamespace + "PASSWORD"
	EnvProjectName      = envNamespace + "PROJECT_NAME"
	EnvIdentityEndpoint = envNamespace + "IDENTITY_ENDPOINT"
	EnvPrivateZone      = envNamespace + "PRIVATE_ZONE"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
	EnvSequenceInterval   = envNamespace + "SEQUENCE_INTERVAL"
)

const defaultIdentityEndpoint = "https://iam.eu-de.otc.t-systems.com:443/v3/auth/tokens"

const minTTL = 300

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	DomainName       string
	ProjectName      string
	UserName         string
	Password         string
	IdentityEndpoint string
	PrivateZone      bool

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	SequenceInterval   time.Duration
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
