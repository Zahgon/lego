package vkcloud

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/vkcloud/internal"
)

const (
	envNamespace = "VK_CLOUD_"

	EnvDNSEndpoint = envNamespace + "DNS_ENDPOINT"

	EnvIdentityEndpoint = envNamespace + "IDENTITY_ENDPOINT"
	EnvDomainName       = envNamespace + "DOMAIN_NAME"

	EnvProjectID = envNamespace + "PROJECT_ID"
	EnvUsername  = envNamespace + "USERNAME"
	EnvPassword  = envNamespace + "PASSWORD"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

const (
	defaultIdentityEndpoint = "https://infra.mail.ru/identity/v3/"
	defaultDNSEndpoint      = "https://mcs.mail.ru/public-dns/v2/dns"
)

const defaultDomainName = "users"

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	ProjectID string
	Username  string
	Password  string

	DNSEndpoint string

	IdentityEndpoint string
	DomainName       string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	client *internal.Client
	config *Config
}

func NewDNSProvider() (*DNSProvider, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) Present(ctx context.Context, domain, _, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) CleanUp(ctx context.Context, domain, _, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) Timeout() (timeout, interval time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}

func (d *DNSProvider) upsertTXTRecord(zoneUUID, name, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) removeTXTRecord(zoneUUID, name, value string) error {
	_ = "STUB: not implemented"
	return nil
}
