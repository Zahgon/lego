package azuredns

import (
	"context"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/go-acme/lego/v5/challenge"
)

const (
	envNamespace = "AZURE_"

	EnvEnvironment    = envNamespace + "ENVIRONMENT"
	EnvSubscriptionID = envNamespace + "SUBSCRIPTION_ID"
	EnvResourceGroup  = envNamespace + "RESOURCE_GROUP"
	EnvZoneName       = envNamespace + "ZONE_NAME"
	EnvPrivateZone    = envNamespace + "PRIVATE_ZONE"

	EnvTenantID     = envNamespace + "TENANT_ID"
	EnvClientID     = envNamespace + "CLIENT_ID"
	EnvClientSecret = envNamespace + "CLIENT_SECRET"

	EnvOIDCToken              = envNamespace + "OIDC_TOKEN"
	EnvOIDCTokenFilePath      = envNamespace + "OIDC_TOKEN_FILE_PATH"
	EnvOIDCRequestURL         = envNamespace + "OIDC_REQUEST_URL"
	EnvGitHubOIDCRequestURL   = "ACTIONS_ID_TOKEN_REQUEST_URL"
	altEnvArmOIDCRequestURL   = "ARM_OIDC_REQUEST_URL"
	EnvOIDCRequestToken       = envNamespace + "OIDC_REQUEST_TOKEN"
	EnvGitHubOIDCRequestToken = "ACTIONS_ID_TOKEN_REQUEST_TOKEN"
	altEnvArmOIDCRequestToken = "ARM_OIDC_REQUEST_TOKEN"

	EnvServiceConnectionID                  = envNamespace + "SERVICE_CONNECTION_ID"
	altEnvServiceConnectionID               = "SERVICE_CONNECTION_ID"
	altEnvArmAdoPipelineServiceConnectionID = "ARM_ADO_PIPELINE_SERVICE_CONNECTION_ID"
	altEnvArmOIDCAzureServiceConnectionID   = "ARM_OIDC_AZURE_SERVICE_CONNECTION_ID"
	EnvSystemAccessToken                    = envNamespace + "SYSTEM_ACCESS_TOKEN"
	altEnvSystemAccessToken                 = "SYSTEM_ACCESSTOKEN"

	EnvAuthMethod     = envNamespace + "AUTH_METHOD"
	EnvAuthMSITimeout = envNamespace + "AUTH_MSI_TIMEOUT"

	EnvServiceDiscoveryFilter = envNamespace + "SERVICEDISCOVERY_FILTER"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

type Config struct {
	ZoneName string

	SubscriptionID string
	ResourceGroup  string
	PrivateZone    bool

	Environment cloud.Configuration

	ClientID     string
	ClientSecret string
	TenantID     string

	OIDCToken         string
	OIDCTokenFilePath string
	OIDCRequestURL    string
	OIDCRequestToken  string

	ServiceConnectionID string
	SystemAccessToken   string

	AuthMethod     string
	AuthMSITimeout time.Duration

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
	HTTPClient         *http.Client

	ServiceDiscoveryFilter string
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

type DNSProvider struct {
	provider challenge.ProviderTimeout
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
