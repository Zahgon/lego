package huaweicloud

import (
	"context"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/challenge/dns01"
	"github.com/go-acme/lego/v5/providers/dns/huaweicloud/internal"
)

const (
	envNamespace = "HUAWEICLOUD_"

	EnvAccessKeyID     = envNamespace + "ACCESS_KEY_ID"
	EnvSecretAccessKey = envNamespace + "SECRET_ACCESS_KEY"
	EnvRegion          = envNamespace + "REGION"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int32
	HTTPTimeout        time.Duration
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *internal.DnsClient

	recordIDs   map[string]string
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

func (d *DNSProvider) getOrCreateRecordSetID(domain, zoneID string, info dns01.ChallengeInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *DNSProvider) getZoneID(authZone string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
