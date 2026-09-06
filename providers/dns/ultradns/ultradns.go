package ultradns

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
)

const (
	envNamespace = "ULTRADNS_"

	EnvUsername = envNamespace + "USERNAME"
	EnvPassword = envNamespace + "PASSWORD"
	EnvEndpoint = envNamespace + "ENDPOINT"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

const defaultEndpoint = "https://api.ultradns.com/"

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type DNSProvider struct {
	config *Config
	client *client.Client
}

type Config struct {
	Username string
	Password string
	Endpoint string

	TTL                int
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

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
