package acmedns

import (
	"context"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/nrdcg/goacmedns"
)

const (
	envNamespace = "ACME_DNS_"

	EnvAPIBase = envNamespace + "API_BASE"

	EnvAllowList = envNamespace + "ALLOWLIST"

	EnvStoragePath = envNamespace + "STORAGE_PATH"

	EnvStorageBaseURL = envNamespace + "STORAGE_BASE_URL"
)

var _ challenge.Provider = (*DNSProvider)(nil)

type Config struct {
	APIBase        string
	AllowList      []string
	StoragePath    string
	StorageBaseURL string
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type acmeDNSClient interface {
	UpdateTXTRecord(ctx context.Context, account goacmedns.Account, value string) error

	RegisterAccount(ctx context.Context, allowFrom []string) (goacmedns.Account, error)
}

type DNSProvider struct {
	config  *Config
	client  acmeDNSClient
	storage goacmedns.Storage
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

func (d *DNSProvider) CleanUp(_ context.Context, _, _, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) register(ctx context.Context, domain, fqdn string) (goacmedns.Account, error) {
	_ = "STUB: not implemented"
	return *new(goacmedns.Account), nil
}

func getStorage(config *Config) (goacmedns.Storage, error) {
	_ = "STUB: not implemented"
	return *new(goacmedns.Storage), nil
}

type ErrCNAMERequired struct {
	Domain string

	FQDN string

	Target string
}

func (e ErrCNAMERequired) Error() string { _ = "STUB: not implemented"; return "" }
