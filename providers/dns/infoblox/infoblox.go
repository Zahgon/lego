package infoblox

import (
	"context"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	infoblox "github.com/infobloxopen/infoblox-go-client/v2"
)

const (
	envNamespace = "INFOBLOX_"

	EnvHost          = envNamespace + "HOST"
	EnvPort          = envNamespace + "PORT"
	EnvUsername      = envNamespace + "USERNAME"
	EnvPassword      = envNamespace + "PASSWORD"
	EnvDNSView       = envNamespace + "DNS_VIEW"
	EnvWApiVersion   = envNamespace + "WAPI_VERSION"
	EnvSSLVerify     = envNamespace + "SSL_VERIFY"
	EnvCACertificate = envNamespace + "CA_CERTIFICATE"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const defaultPoolConnections = 10

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	Host string

	Port string

	Username string

	Password string

	DNSView string

	WapiVersion string

	SSLVerify bool

	CACertificate string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPTimeout        int
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config          *Config
	transportConfig infoblox.TransportConfig
	ibConfig        infoblox.HostConfig
	ibAuth          infoblox.AuthConfig

	recordRefs   map[string]string
	recordRefsMu sync.Mutex
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
