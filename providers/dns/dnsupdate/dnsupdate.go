package dnsupdate

import (
	"context"
	"time"

	"github.com/bodgit/tsig/gss"
	"github.com/go-acme/lego/v5/challenge"
)

const (
	envNamespace = "DNSUPDATE_"

	EnvNameserver = envNamespace + "NAMESERVER"
	EnvDNSTimeout = envNamespace + "DNS_TIMEOUT"

	envTimeout = envNamespace + "TIMEOUT"

	EnvZones = envNamespace + "ZONES"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvSequenceInterval   = envNamespace + "SEQUENCE_INTERVAL"
)

const (
	envTSIG = envNamespace + "TSIG_"

	EnvTSIGFile = envTSIG + "FILE"

	EnvTSIGKey       = envTSIG + "KEY"
	EnvTSIGSecret    = envTSIG + "SECRET"
	EnvTSIGAlgorithm = envTSIG + "ALGORITHM"
)

const (
	envSubTSIGGSS = "TSIG_GSS_"

	envTSIGGSS = envNamespace + envSubTSIGGSS

	EnvTSIGGSSRealm      = envTSIGGSS + "REALM"
	EnvTSIGGSSUsername   = envTSIGGSS + "USERNAME"
	EnvTSIGGSSPassword   = envTSIGGSS + "PASSWORD"
	EnvTSIGGSSKeytabFile = envTSIGGSS + "KEYTAB_FILE"
)

const (
	actionRemove = "REMOVE"
	actionInsert = "INSERT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	Nameserver string
	DNSTimeout time.Duration

	Zones []string

	TSIGFile string

	TSIGAlgorithm string
	TSIGKey       string
	TSIGSecret    string

	TSIGGSSRealm      string
	TSIGGSSUsername   string
	TSIGGSSPassword   string
	TSIGGSSKeytabFile string

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	SequenceInterval   time.Duration
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

func (d *DNSProvider) Sequential() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (d *DNSProvider) Present(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) changeRecord(ctx context.Context, action, fqdn, value string, ttl int) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) negotiate(client *gss.Client) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *DNSProvider) findZone(ctx context.Context, fqdn string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func setupTSIG(config *Config) error { _ = "STUB: not implemented"; return nil }

func validateTSIGGSS(config *Config) error { _ = "STUB: not implemented"; return nil }

func prepareTSIG(config *Config) error { _ = "STUB: not implemented"; return nil }

func parseNameserver(ns string) (string, error) { _ = "STUB: not implemented"; return "", nil }
