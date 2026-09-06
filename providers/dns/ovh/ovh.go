package ovh

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/ovh/go-ovh/ovh"
)

const (
	envNamespace = "OVH_"

	EnvEndpoint = envNamespace + "ENDPOINT"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const (
	EnvApplicationKey    = envNamespace + "APPLICATION_KEY"
	EnvApplicationSecret = envNamespace + "APPLICATION_SECRET"
	EnvConsumerKey       = envNamespace + "CONSUMER_KEY"
)

const (
	EnvClientID     = envNamespace + "CLIENT_ID"
	EnvClientSecret = envNamespace + "CLIENT_SECRET"
)

const EnvAccessToken = envNamespace + "ACCESS_TOKEN"

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Record struct {
	ID        int64  `json:"id,omitempty"`
	FieldType string `json:"fieldType,omitempty"`
	SubDomain string `json:"subDomain,omitempty"`
	Target    string `json:"target,omitempty"`
	TTL       int    `json:"ttl,omitempty"`
	Zone      string `json:"zone,omitempty"`
}

type OAuth2Config struct {
	ClientID     string
	ClientSecret string
}

type Config struct {
	APIEndpoint string

	ApplicationKey    string
	ApplicationSecret string
	ConsumerKey       string

	OAuth2Config *OAuth2Config

	AccessToken string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

func (c *Config) hasAppKeyAuth() bool { _ = "STUB: not implemented"; return false }

type DNSProvider struct {
	config *Config
	client *ovh.Client

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

func newClient(config *Config) (*ovh.Client, error) { _ = "STUB: not implemented"; return nil, nil }
