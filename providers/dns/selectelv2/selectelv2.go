package selectelv2

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	selectelapi "github.com/selectel/domains-go/pkg/v2"
)

const (
	envNamespace = "SELECTELV2_"

	EnvBaseURL        = envNamespace + "BASE_URL"
	EnvUsernameOS     = envNamespace + "USERNAME"
	EnvPasswordOS     = envNamespace + "PASSWORD"
	EnvDomainName     = envNamespace + "ACCOUNT_ID"
	EnvProjectID      = envNamespace + "PROJECT_ID"
	EnvAuthRegion     = envNamespace + "AUTH_REGION"
	EnvAuthURL        = envNamespace + "AUTH_URL"
	EnvUserDomainName = envNamespace + "USER_DOMAIN_NAME"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const (
	defaultBaseURL    = "https://api.selectel.ru/domains/v2"
	defaultAuthRegion = "ru-1"
	defaultAuthURL    = "https://cloud.api.selcloud.ru/identity/v3/"
)

const (
	defaultTTL                = 60
	defaultPropagationTimeout = 120 * time.Second
	defaultPollingInterval    = 5 * time.Second
	defaultHTTPTimeout        = 30 * time.Second
)

const tokenHeader = "X-Auth-Token"

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

var errNotFound = errors.New("rrset not found")

type Config struct {
	BaseURL        string
	Username       string
	Password       string
	DomainName     string
	ProjectID      string
	AuthURL        string
	AuthRegion     string
	UserDomainName string

	TTL                int
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	baseClient selectelapi.DNSClient[selectelapi.Zone, selectelapi.RRSet]
	config     *Config
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

func (d *DNSProvider) Present(ctx context.Context, domain, _, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) CleanUp(ctx context.Context, domain, _, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) authorize(ctx context.Context) (*clientWrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func obtainOpenstackToken(ctx context.Context, config *Config) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type clientWrapper struct {
	selectelapi.DNSClient[selectelapi.Zone, selectelapi.RRSet]
}

func (w *clientWrapper) getZone(ctx context.Context, name string) (*selectelapi.Zone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *clientWrapper) getRRset(ctx context.Context, name, zoneID string) (*selectelapi.RRSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
