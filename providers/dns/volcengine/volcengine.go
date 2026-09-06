package volcengine

import (
	"context"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	volc "github.com/volcengine/volc-sdk-golang/service/dns"
)

const (
	envNamespace = "VOLC_"

	EnvAccessKey = envNamespace + "ACCESSKEY"
	EnvSecretKey = envNamespace + "SECRETKEY"

	EnvRegion = envNamespace + "REGION"
	EnvHost   = envNamespace + "HOST"
	EnvScheme = envNamespace + "SCHEME"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const defaultTTL = 600

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	AccessKey string
	SecretKey string

	Region string
	Host   string
	Scheme string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPTimeout        time.Duration
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	client *volc.Client
	config *Config

	recordIDs   map[string]*string
	recordIDsMu sync.Mutex
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

func (d *DNSProvider) getZone(ctx context.Context, fqdn string) (volc.TopZoneResponse, error) {
	_ = "STUB: not implemented"
	return *new(volc.TopZoneResponse), nil
}

func newClient(config *Config) *volc.Client { _ = "STUB: not implemented"; return nil }
