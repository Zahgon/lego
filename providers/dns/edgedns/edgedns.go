package edgedns

import (
	"context"
	"time"

	edgegriddns "github.com/akamai/AkamaiOPEN-edgegrid-golang/v13/pkg/dns"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v13/pkg/edgegrid"
	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/challenge/dns01"
)

const (
	envNamespace = "AKAMAI_"

	EnvEdgeRc           = envNamespace + "EDGERC"
	EnvEdgeRcSection    = envNamespace + "EDGERC_SECTION"
	EnvAccountSwitchKey = envNamespace + "ACCOUNT_SWITCH_KEY"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

const (
	EnvHost         = envNamespace + "HOST"
	EnvClientToken  = envNamespace + "CLIENT_TOKEN"
	EnvClientSecret = envNamespace + "CLIENT_SECRET"
	EnvAccessToken  = envNamespace + "ACCESS_TOKEN"
)

const (
	defaultPropagationTimeout = 3 * time.Minute
	defaultPollInterval       = 15 * time.Second
)

const maxBody = 131072

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	*edgegrid.Config

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
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

func getZone(ctx context.Context, domain string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func containsValue(values []string, value string) bool { _ = "STUB: not implemented"; return false }

func isNotFound(err error) bool { _ = "STUB: not implemented"; return false }

func filterRData(existingRec *edgegriddns.GetRecordResponse, info dns01.ChallengeInfo) []string {
	_ = "STUB: not implemented"
	return nil
}
