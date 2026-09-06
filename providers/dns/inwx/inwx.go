package inwx

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/nrdcg/goinwx"
)

const (
	envNamespace = "INWX_"

	EnvUsername     = envNamespace + "USERNAME"
	EnvPassword     = envNamespace + "PASSWORD"
	EnvSharedSecret = envNamespace + "SHARED_SECRET"
	EnvSandbox      = envNamespace + "SANDBOX"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	Username           string
	Password           string
	SharedSecret       string
	Sandbox            bool
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config         *Config
	client         *goinwx.Client
	previousUnlock time.Time
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

func (d *DNSProvider) twoFactorAuth(info *goinwx.LoginResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) computeSleep(now time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
