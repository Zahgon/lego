package vinyldns

import (
	"context"

	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func (d *DNSProvider) getRecordSet(ctx context.Context, fqdn string) (*vinyldns.RecordSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DNSProvider) createRecordSet(ctx context.Context, fqdn string, records []vinyldns.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) updateRecordSet(ctx context.Context, recordSet *vinyldns.RecordSet, newRecords []vinyldns.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) deleteRecordSet(ctx context.Context, existingRecord *vinyldns.RecordSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DNSProvider) waitForChanges(ctx context.Context, operation string, resp *vinyldns.RecordSetUpdateResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func splitDomain(ctx context.Context, fqdn string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
