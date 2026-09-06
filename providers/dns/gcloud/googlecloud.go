package gcloud

import (
	"context"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"golang.org/x/oauth2"
	gdns "google.golang.org/api/dns/v1"
)

const (
	envNamespace = "GCE_"

	EnvServiceAccount            = envNamespace + "SERVICE_ACCOUNT"
	EnvProject                   = envNamespace + "PROJECT"
	EnvZoneID                    = envNamespace + "ZONE_ID"
	EnvAllowPrivateZone          = envNamespace + "ALLOW_PRIVATE_ZONE"
	EnvDebug                     = envNamespace + "DEBUG"
	EnvImpersonateServiceAccount = envNamespace + "IMPERSONATE_SERVICE_ACCOUNT"
	EnvAccessToken               = envNamespace + "ACCESS_TOKEN"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

const changeStatusDone = "done"

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	Debug                     bool
	Project                   string
	ZoneID                    string
	AllowPrivateZone          bool
	ImpersonateServiceAccount string
	AccessToken               string
	PropagationTimeout        time.Duration
	PollingInterval           time.Duration
	TTL                       int
	HTTPClient                *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *gdns.Service
}

func NewDNSProvider() (*DNSProvider, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDNSProviderCredentials(project string) (*DNSProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDNSProviderServiceAccountKey(saKey []byte) (*DNSProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDNSProviderServiceAccount(saFile string) (*DNSProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) applyChanges(ctx context.Context, zone string, change *gdns.Change) error {
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

func (d *DNSProvider) getHostedZone(ctx context.Context, domain string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *DNSProvider) lookupHostedZoneID(ctx context.Context, domain string) (string, []*gdns.ManagedZone, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (d *DNSProvider) findTxtRecords(zone, fqdn string) ([]*gdns.ResourceRecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newClientFromCredentials(ctx context.Context, config *Config) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newClientFromServiceAccountKey(ctx context.Context, config *Config, saKey []byte) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newImpersonateClient(ctx context.Context, impersonateServiceAccount string, ts oauth2.TokenSource) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTokenSource(ctx context.Context, config *Config) (oauth2.TokenSource, error) {
	_ = "STUB: not implemented"
	return *new(oauth2.TokenSource), nil
}

func mustUnquote(raw string) string { _ = "STUB: not implemented"; return "" }

func autodetectProjectID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }
