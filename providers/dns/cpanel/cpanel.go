package cpanel

import (
	"context"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/cpanel/internal/shared"
)

const (
	envNamespace = "CPANEL_"

	EnvMode     = envNamespace + "MODE"
	EnvUsername = envNamespace + "USERNAME"
	EnvToken    = envNamespace + "TOKEN"
	EnvBaseURL  = envNamespace + "BASE_URL"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type apiClient interface {
	FetchZoneInformation(ctx context.Context, domain string) ([]shared.ZoneRecord, error)
	AddRecord(ctx context.Context, serial uint32, domain string, record shared.Record) (*shared.ZoneSerial, error)
	EditRecord(ctx context.Context, serial uint32, domain string, record shared.Record) (*shared.ZoneSerial, error)
	DeleteRecord(ctx context.Context, serial uint32, domain string, lineIndex int) (*shared.ZoneSerial, error)
}

type Config struct {
	Mode               string
	Username           string
	Token              string
	BaseURL            string
	TTL                int
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client apiClient
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

func (d *DNSProvider) Present(ctx context.Context, domain, _, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) CleanUp(ctx context.Context, domain, _, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func getZoneSerial(zoneFqdn string, zoneInfo []shared.ZoneRecord) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func createClient(config *Config) (apiClient, error) {
	_ = "STUB: not implemented"
	return *new(apiClient), nil
}
