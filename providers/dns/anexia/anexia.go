package anexia

import (
	"context"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/anexia/internal"
)

const (
	envNamespace = "ANEXIA_"

	EnvToken  = envNamespace + "TOKEN"
	EnvAPIURL = envNamespace + "API_URL"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const defaultTTL = 300

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	Token  string
	APIURL string

	TTL                int
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
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

func (d *DNSProvider) findRecordID(ctx context.Context, zoneName, recordName, rdata string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func findRecordIdentifier(zone *internal.Zone, recordName, rdata string) string {
	_ = "STUB: not implemented"
	return ""
}

func extractRecordName(fqdn, authZone string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
