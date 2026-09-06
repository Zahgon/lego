package oraclecloud

import (
	"context"
	"net/http"
	"time"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/nrdcg/oci-go-sdk/common/v1065"
	"github.com/nrdcg/oci-go-sdk/dns/v1065"
)

const (
	envNamespace = "OCI_"

	EnvAuthType = envNamespace + "AUTH_TYPE"

	EnvCompartmentOCID = envNamespace + "COMPARTMENT_OCID"
	EnvRegion          = envNamespace + "REGION"

	EnvProfile    = envNamespace + "PROFILE"
	EnvConfigFile = envNamespace + "CONFIG_FILE"

	envPrivKey           = envNamespace + "PRIVKEY"
	EnvPrivKeyFile       = envPrivKey + "_FILE"
	EnvPrivKeyPass       = envPrivKey + "_PASS"
	EnvTenancyOCID       = envNamespace + "TENANCY_OCID"
	EnvUserOCID          = envNamespace + "USER_OCID"
	EnvPubKeyFingerprint = envNamespace + "PUBKEY_FINGERPRINT"

	altEnvPrivateKey         = envNamespace + "PRIVATE_KEY"
	altEnvPrivateKeyPath     = altEnvPrivateKey + "_PATH"
	altEnvPrivateKeyPassword = altEnvPrivateKey + "_PASSWORD"
	altEnvFingerprint        = envNamespace + "FINGERPRINT"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

const (
	altEnvTFVarNamespace          = "TF_VAR_"
	altEnvTFVarRegion             = altEnvTFVarNamespace + "region"
	altEnvTFVarFingerprint        = altEnvTFVarNamespace + "fingerprint"
	altEnvTFVarUserOCID           = altEnvTFVarNamespace + "user_ocid"
	altEnvTFVarTenancyOCID        = altEnvTFVarNamespace + "tenancy_ocid"
	altEnvTFVarPrivateKeyPath     = altEnvTFVarNamespace + "private_key_path"
	altEnvTFVarPrivateKeyPassword = altEnvTFVarNamespace + "private_key_password"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	CompartmentID     string
	OCIConfigProvider common.ConfigurationProvider

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPClient         *http.Client
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	client *dns.DnsClient
	config *Config
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
