package azuredns

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
	"github.com/go-acme/lego/v5/challenge"
)

var _ challenge.ProviderTimeout = (*DNSProviderPrivate)(nil)

type DNSProviderPrivate struct {
	config                *Config
	credentials           azcore.TokenCredential
	serviceDiscoveryZones map[string]ServiceDiscoveryZone
}

func NewDNSProviderPrivate(config *Config, credentials azcore.TokenCredential) (*DNSProviderPrivate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProviderPrivate) Timeout() (timeout, interval time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}

func (d *DNSProviderPrivate) Present(ctx context.Context, domain, _, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProviderPrivate) CleanUp(ctx context.Context, domain, _, keyAuth string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProviderPrivate) getHostedZone(ctx context.Context, fqdn string) (ServiceDiscoveryZone, error) {
	_ = "STUB: not implemented"
	return *new(ServiceDiscoveryZone), nil
}

type privateZoneClient struct {
	zone         ServiceDiscoveryZone
	recordClient *armprivatedns.RecordSetsClient
}

func newPrivateZoneClient(zone ServiceDiscoveryZone, credential azcore.TokenCredential, environment cloud.Configuration) (*privateZoneClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c privateZoneClient) Get(ctx context.Context, subDomain string) (armprivatedns.RecordSetsClientGetResponse, error) {
	_ = "STUB: not implemented"
	return *new(armprivatedns.RecordSetsClientGetResponse), nil
}

func (c privateZoneClient) CreateOrUpdate(ctx context.Context, subDomain string, rec armprivatedns.RecordSet) (armprivatedns.RecordSetsClientCreateOrUpdateResponse, error) {
	_ = "STUB: not implemented"
	return *new(armprivatedns.RecordSetsClientCreateOrUpdateResponse), nil
}

func (c privateZoneClient) Delete(ctx context.Context, subDomain string) (armprivatedns.RecordSetsClientDeleteResponse, error) {
	_ = "STUB: not implemented"
	return *new(armprivatedns.RecordSetsClientDeleteResponse), nil
}

func privateUniqueRecords(recordSet armprivatedns.RecordSet, value string) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}
