package azuredns

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

type ServiceDiscoveryZone struct {
	Name           string
	SubscriptionID string
	ResourceGroup  string
}

const (
	ResourceGraphTypePublicDNSZone  = "microsoft.network/dnszones"
	ResourceGraphTypePrivateDNSZone = "microsoft.network/privatednszones"
)

const ResourceGraphQueryOptionsTop int32 = 1000

func discoverDNSZones(ctx context.Context, config *Config, credentials azcore.TokenCredential) (map[string]ServiceDiscoveryZone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createGraphQuery(config *Config) string { _ = "STUB: not implemented"; return "" }
