package jdcloud

import (
	"context"
	"sync"
	"time"

	jdcclient "github.com/go-acme/jdcloud-sdk-go/services/domainservice/client"
	domainservice "github.com/go-acme/jdcloud-sdk-go/services/domainservice/models"
	"github.com/go-acme/lego/v5/challenge"
)

const (
	envNamespace = "JDCLOUD_"

	EnvAccessKeyID     = envNamespace + "ACCESS_KEY_ID"
	EnvAccessKeySecret = envNamespace + "ACCESS_KEY_SECRET"
	EnvRegionID        = envNamespace + "REGION_ID"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	AccessKeyID     string
	AccessKeySecret string
	RegionID        string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPTimeout        time.Duration
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	config *Config
	client *jdcclient.DomainserviceClient

	recordIDs   map[string]int
	domainIDs   map[string]int
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

func (d *DNSProvider) findZone(zone string) (*domainservice.DomainInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
