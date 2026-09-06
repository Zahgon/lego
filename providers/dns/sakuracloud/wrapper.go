package sakuracloud

import (
	"context"
	"sync"

	"github.com/sacloud/sacloud-sdk-go/api/iaas"
)

var mu sync.Mutex

func (d *DNSProvider) addTXTRecord(ctx context.Context, fqdn, value string, ttl int) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) cleanupTXTRecord(ctx context.Context, fqdn, value string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck // Clearer without De Morgan's law.

func (d *DNSProvider) getHostedZone(ctx context.Context, domain string) (*iaas.DNS, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
