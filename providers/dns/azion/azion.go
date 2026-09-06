package azion

import (
	"context"
	"net/http"
	"time"

	"github.com/aziontech/azionapi-go-sdk/idns"
	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/challenge/dns01"
)

const (
	envNamespace = "AZION_"

	EnvPersonalToken = envNamespace + "PERSONAL_TOKEN"
	EnvPageSize      = envNamespace + "PAGE_SIZE"

	EnvTTL                = envNamespace + "TTL"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	PersonalToken string
	PageSize      int

	PollingInterval    time.Duration
	PropagationTimeout time.Duration
	TTL                int
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *idns.APIClient
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

func (d *DNSProvider) findZone(ctx context.Context, fqdn string) (*idns.Zone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) findExistingTXTRecord(ctx context.Context, zoneID int32, recordName string) (*idns.RecordGet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authContext(ctx context.Context, key string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func extractSubDomain(info dns01.ChallengeInfo, zone *idns.Zone) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
