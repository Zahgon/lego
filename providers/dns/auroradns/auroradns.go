package auroradns

import (
	"context"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/nrdcg/auroradns"
)

const (
	envNamespace = "AURORA_"

	EnvAPIKey   = envNamespace + "API_KEY"
	EnvSecret   = envNamespace + "SECRET"
	EnvEndpoint = envNamespace + "ENDPOINT"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

const defaultBaseURL = "https://api.auroradns.eu"

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	BaseURL            string
	APIKey             string
	Secret             string
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *auroradns.Client

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

func (d *DNSProvider) getZoneInformationByName(name string) (auroradns.Zone, error) {
	_ = "STUB: not implemented"
	return *new(auroradns.Zone), nil
}
