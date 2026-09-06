package gandiv5

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/gandiv5/internal"
)

const (
	envNamespace = "GANDIV5_"

	EnvAPIKey              = envNamespace + "API_KEY"
	EnvPersonalAccessToken = envNamespace + "PERSONAL_ACCESS_TOKEN"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const minTTL = 300

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type inProgressInfo struct {
	fieldName string
	authZone  string
}

type Config struct {
	BaseURL             string
	APIKey              string
	PersonalAccessToken string
	PropagationTimeout  time.Duration
	PollingInterval     time.Duration
	TTL                 int
	HTTPClient          *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *internal.Client

	inProgressFQDNs map[string]inProgressInfo
	inProgressMu    sync.Mutex

	findZoneByFqdn func(ctx context.Context, fqdn string) (string, error)
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
