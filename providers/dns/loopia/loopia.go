package loopia

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/loopia/internal"
)

const (
	envNamespace = "LOOPIA_"

	EnvAPIUser     = envNamespace + "API_USER"
	EnvAPIPassword = envNamespace + "API_PASSWORD"
	EnvAPIURL      = envNamespace + "API_URL"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const minTTL = 300

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type dnsClient interface {
	AddTXTRecord(ctx context.Context, domain, subdomain string, ttl int, value string) error
	RemoveTXTRecord(ctx context.Context, domain, subdomain string, recordID int) error
	GetTXTRecords(ctx context.Context, domain, subdomain string) ([]internal.RecordObj, error)
	RemoveSubdomain(ctx context.Context, domain, subdomain string) error
}

type Config struct {
	BaseURL            string
	APIUser            string
	APIPassword        string
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client dnsClient

	inProgressInfo map[string]int
	inProgressMu   sync.Mutex

	findZoneByFqdn func(ctx context.Context, fqdn string) (string, error)
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

func (d *DNSProvider) splitDomain(ctx context.Context, fqdn string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
