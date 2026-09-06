package gigahostno

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/gigahostno/internal"
)

const (
	envNamespace = "GIGAHOSTNO_"

	EnvUsername = envNamespace + "USERNAME"
	EnvPassword = envNamespace + "PASSWORD"
	EnvSecret   = envNamespace + "SECRET"

	EnvAPIKey = envNamespace + "API_KEY"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Identifier interface {
	Authenticate(ctx context.Context) (*internal.Token, error)
}

type Config struct {
	Username string
	Password string
	Secret   string

	APIkey string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config

	identifier Identifier
	client     *internal.Client

	tokenMu sync.Mutex
	token   *internal.Token
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

func (d *DNSProvider) authenticate(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) findZone(ctx context.Context, fqdn string) (*internal.Zone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newIdentifier(config *Config) (Identifier, error) {
	_ = "STUB: not implemented"
	return *new(Identifier), nil
}
