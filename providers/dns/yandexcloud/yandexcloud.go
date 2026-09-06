package yandexcloud

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	ycdnsproto "github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
	ycdns "github.com/yandex-cloud/go-sdk/services/dns/v1"
	"github.com/yandex-cloud/go-sdk/v2/credentials"
)

const (
	envNamespace = "YANDEX_CLOUD_"

	EnvIamToken = envNamespace + "IAM_TOKEN"
	EnvFolderID = envNamespace + "FOLDER_ID"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	IamToken string
	FolderID string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	client ycdns.DnsZoneClient
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

func (d *DNSProvider) getZones(ctx context.Context) ([]*ycdnsproto.DnsZone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) upsertRecordSetData(ctx context.Context, zoneID, name, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) removeRecordSetData(ctx context.Context, zoneID, name, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeCredentials(accountB64 string) (credentials.Credentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.Credentials), nil
}

func appendRecordSetData(record *ycdnsproto.RecordSet, value string) bool {
	_ = "STUB: not implemented"
	return false
}
