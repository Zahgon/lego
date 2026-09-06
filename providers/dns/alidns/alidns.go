package alidns

import (
	"context"
	"time"

	alidns "github.com/go-acme/alidns-20150109/v5/client"
	"github.com/go-acme/lego/v5/challenge"
)

const (
	envNamespace = "ALICLOUD_"

	EnvRAMRole       = envNamespace + "RAM_ROLE"
	EnvAccessKey     = envNamespace + "ACCESS_KEY"
	EnvSecretKey     = envNamespace + "SECRET_KEY"
	EnvSecurityToken = envNamespace + "SECURITY_TOKEN"
	EnvRegionID      = envNamespace + "REGION_ID"
	EnvLine          = envNamespace + "LINE"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const defaultRegionID = "cn-hangzhou"

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	RAMRole            string
	APIKey             string
	SecretKey          string
	SecurityToken      string
	RegionID           string
	Line               string
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPTimeout        time.Duration
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *alidns.Client
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

func (d *DNSProvider) getHostedZone(ctx context.Context, domain string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *DNSProvider) newTxtRecord(zone, fqdn, value string) (*alidns.AddDomainRecordRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) findTxtRecords(ctx context.Context, fqdn string) ([]*alidns.DescribeDomainRecordsResponseBodyDomainRecordsRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractRecordName(fqdn, zone string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
