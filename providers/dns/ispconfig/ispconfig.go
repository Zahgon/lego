package ispconfig

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/ispconfig/internal"
)

const (
	envNamespace = "ISPCONFIG_"

	EnvServerURL = envNamespace + "SERVER_URL"
	EnvUsername  = envNamespace + "USERNAME"
	EnvPassword  = envNamespace + "PASSWORD"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
	EnvInsecureSkipVerify = envNamespace + "INSECURE_SKIP_VERIFY"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	ServerURL string
	Username  string
	Password  string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPClient         *http.Client
	InsecureSkipVerify bool
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *internal.Client

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

func (d *DNSProvider) findZone(ctx context.Context, sessionID, fqdn string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
