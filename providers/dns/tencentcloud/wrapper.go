package tencentcloud

import (
	"context"

	dnspod "github.com/go-acme/tencentclouddnspod/v20210323"
)

func (d *DNSProvider) getHostedZone(ctx context.Context, domain string) (*dnspod.DomainListItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) findTxtRecords(ctx context.Context, zone *dnspod.DomainListItem, fqdn string) ([]*dnspod.RecordListItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractRecordName(fqdn, zone string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
