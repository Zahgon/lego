package mittwald

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/mittwald/internal"
)

const (
	envNamespace = "MITTWALD_"

	EnvToken = envNamespace + "TOKEN"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvSequenceInterval   = envNamespace + "SEQUENCE_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const minTTL = 300

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	Token              string
	TTL                int
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	SequenceInterval   time.Duration
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *internal.Client

	zoneIDs   map[string]string
	zoneIDsMu sync.Mutex
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

func (d *DNSProvider) Sequential() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (d *DNSProvider) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) getOrCreateZone(ctx context.Context, fqdn string) (*internal.DNSZone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findDomain(domains []internal.Domain, fqdn string) (internal.Domain, error) {
	_ = "STUB: not implemented"
	return *new(internal.Domain), nil
}

func findZone(zones []internal.DNSZone, fqdn string) (internal.DNSZone, error) {
	_ = "STUB: not implemented"
	return *new(internal.DNSZone), nil
}
