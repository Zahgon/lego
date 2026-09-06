package iijdpf

import (
	"context"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	dpfapi "github.com/mimuret/golang-iij-dpf/pkg/api"
)

const (
	envNamespace = "IIJ_DPF_"

	EnvAPIToken    = envNamespace + "API_TOKEN"
	EnvServiceCode = envNamespace + "DPM_SERVICE_CODE"

	EnvAPIEndpoint        = envNamespace + "API_ENDPOINT"
	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	Token       string
	ServiceCode string

	Endpoint           string
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	client dpfapi.ClientInterface
	config *Config
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
