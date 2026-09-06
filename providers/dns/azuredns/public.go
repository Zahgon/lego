package azuredns

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/go-acme/lego/v5/challenge"
)

var _ challenge.ProviderTimeout = (*DNSProviderPublic)(nil)

type DNSProviderPublic struct {
	config                *Config
	credentials           azcore.TokenCredential
	serviceDiscoveryZones map[string]ServiceDiscoveryZone
}

func NewDNSProviderPublic(config *Config, credentials azcore.TokenCredential) (*DNSProviderPublic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProviderPublic) Timeout() (timeout, interval time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}

func (d *DNSProviderPublic) Present(ctx context.Context, domain, _, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProviderPublic) CleanUp(ctx context.Context, domain, _, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProviderPublic) getHostedZone(ctx context.Context, fqdn string) (ServiceDiscoveryZone, error) {
	_ = "STUB: not implemented"
	return *new(ServiceDiscoveryZone), nil
}

type publicZoneClient struct {
	zone         ServiceDiscoveryZone
	recordClient *armdns.RecordSetsClient
}

func newPublicZoneClient(zone ServiceDiscoveryZone, credential azcore.TokenCredential, environment cloud.Configuration) (*publicZoneClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c publicZoneClient) Get(ctx context.Context, subDomain string) (armdns.RecordSetsClientGetResponse, error) {
	_ = "STUB: not implemented"
	return *new(armdns.RecordSetsClientGetResponse), nil
}

func (c publicZoneClient) CreateOrUpdate(ctx context.Context, subDomain string, rec armdns.RecordSet) (armdns.RecordSetsClientCreateOrUpdateResponse, error) {
	_ = "STUB: not implemented"
	return *new(armdns.RecordSetsClientCreateOrUpdateResponse), nil
}

func (c publicZoneClient) Delete(ctx context.Context, subDomain string) (armdns.RecordSetsClientDeleteResponse, error) {
	_ = "STUB: not implemented"
	return *new(armdns.RecordSetsClientDeleteResponse), nil
}

func publicUniqueRecords(recordSet armdns.RecordSet, value string) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}
