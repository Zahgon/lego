package bluecatv2

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/bluecatv2/internal"
)

const (
	envNamespace = "BLUECATV2_"

	EnvServerURL  = envNamespace + "SERVER_URL"
	EnvUsername   = envNamespace + "USERNAME"
	EnvPassword   = envNamespace + "PASSWORD"
	EnvConfigName = envNamespace + "CONFIG_NAME"
	EnvViewName   = envNamespace + "VIEW_NAME"
	EnvSkipDeploy = envNamespace + "SKIP_DEPLOY"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	ServerURL  string
	Username   string
	Password   string
	ConfigName string
	ViewName   string
	SkipDeploy bool

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *internal.Client

	zoneIDs     map[string]int64
	recordIDs   map[string]int64
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

func (d *DNSProvider) findZone(ctx context.Context, fqdn string) (*internal.ZoneResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
