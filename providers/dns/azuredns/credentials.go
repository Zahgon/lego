package azuredns

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

const (
	authMethodEnv      = "env"
	authMethodWLI      = "wli"
	authMethodMSI      = "msi"
	authMethodCLI      = "cli"
	authMethodOIDC     = "oidc"
	authMethodPipeline = "pipeline"
)

//nolint:gocyclo // The complexity is related to the number of possible configurations.
func getCredentials(config *Config) (azcore.TokenCredential, error) {
	_ = "STUB: not implemented"
	return *new(azcore.TokenCredential), nil
}

type timeoutTokenCredential struct {
	cred    azcore.TokenCredential
	timeout time.Duration
}

func (w *timeoutTokenCredential) GetToken(ctx context.Context, opts policy.TokenRequestOptions) (azcore.AccessToken, error) {
	_ = "STUB: not implemented"
	return *new(azcore.AccessToken), nil
}

func getZoneName(ctx context.Context, config *Config, fqdn string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func checkPipelineConfig(config *Config) error { _ = "STUB: not implemented"; return nil }
