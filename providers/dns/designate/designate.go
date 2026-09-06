package designate

import (
	"context"
	"sync"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
)

const (
	envNamespace = "DESIGNATE_"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"

	EnvZoneName = envNamespace + "ZONE_NAME"

	envNamespaceClient = "OS_"

	EnvAuthURL       = envNamespaceClient + "AUTH_URL"
	EnvUsername      = envNamespaceClient + "USERNAME"
	EnvPassword      = envNamespaceClient + "PASSWORD"
	EnvUserID        = envNamespaceClient + "USER_ID"
	EnvAppCredID     = envNamespaceClient + "APPLICATION_CREDENTIAL_ID"
	EnvAppCredName   = envNamespaceClient + "APPLICATION_CREDENTIAL_NAME"
	EnvAppCredSecret = envNamespaceClient + "APPLICATION_CREDENTIAL_SECRET"
	EnvTenantName    = envNamespaceClient + "TENANT_NAME"
	EnvRegionName    = envNamespaceClient + "REGION_NAME"
	EnvProjectID     = envNamespaceClient + "PROJECT_ID"
	EnvCloud         = envNamespaceClient + "CLOUD"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	ZoneName           string
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int

	opts       gophercloud.AuthOptions
	regionName string
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *gophercloud.ServiceClient

	dnsEntriesMu sync.Mutex
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

func (d *DNSProvider) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) createRecord(zoneID, fqdn, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) updateRecord(record *recordsets.RecordSet, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) getZoneID(wanted string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *DNSProvider) getRecord(zoneID, wanted string) (*recordsets.RecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) getZoneName(ctx context.Context, fqdn string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
