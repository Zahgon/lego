package liquidweb

import (
	"context"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	lw "github.com/liquidweb/liquidweb-go/client"
)

const (
	envNamespace    = "LIQUID_WEB_"
	altEnvNamespace = "LWAPI_"

	EnvURL      = envNamespace + "URL"
	EnvUsername = envNamespace + "USERNAME"
	EnvPassword = envNamespace + "PASSWORD"
	EnvZone     = envNamespace + "ZONE"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const defaultBaseURL = "https://api.liquidweb.com"

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	BaseURL            string
	Username           string
	Password           string
	Zone               string
	TTL                int
	PollingInterval    time.Duration
	PropagationTimeout time.Duration
	HTTPTimeout        time.Duration
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *lw.API

	recordIDs   map[string]int
	recordIDsMu sync.Mutex
}

func NewDNSProvider() (*DNSProvider, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) Timeout() (time.Duration, time.Duration) {
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

func (d *DNSProvider) findZone(domain string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func altEnvName(v string) string { _ = "STUB: not implemented"; return "" }
