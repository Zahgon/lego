package aliesa

import (
	"context"
	"sync"
	"time"

	esa "github.com/go-acme/esa-20240910/v3/client"
	"github.com/go-acme/lego/v5/challenge"
)

const (
	envNamespace = "ALIESA_"

	EnvRAMRole       = envNamespace + "RAM_ROLE"
	EnvAccessKey     = envNamespace + "ACCESS_KEY"
	EnvSecretKey     = envNamespace + "SECRET_KEY"
	EnvSecurityToken = envNamespace + "SECURITY_TOKEN"
	EnvRegionID      = envNamespace + "REGION_ID"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const defaultRegionID = "cn-hangzhou"

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	RAMRole       string
	APIKey        string
	SecretKey     string
	SecurityToken string
	RegionID      string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPTimeout        time.Duration
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *esa.Client

	recordIDs   map[string]int64
	recordIDsMu sync.Mutex
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

func (d *DNSProvider) getSiteID(ctx context.Context, fqdn string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
