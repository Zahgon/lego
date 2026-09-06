package iijdpf

import (
	"context"
)

func (d *DNSProvider) addTxtRecord(ctx context.Context, zoneID, fqdn, rdata string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) deleteTxtRecord(ctx context.Context, zoneID, fqdn, rdata string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) commit(ctx context.Context, zoneID string) error {
	_ = "STUB: not implemented"
	return nil
}
