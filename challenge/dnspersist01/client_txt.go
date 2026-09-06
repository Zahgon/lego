package dnspersist01

import (
	"context"

	"github.com/miekg/dns"
)

const maxCNAMEFollows = 50

type TXTRecord struct {
	Value string
	TTL   uint32
}

type TXTResult struct {
	Records    []TXTRecord
	CNAMEChain []string
}

func (r TXTResult) String() string { _ = "STUB: not implemented"; return "" }

func (c *Client) LookupTXT(ctx context.Context, fqdn string) (TXTResult, error) {
	_ = "STUB: not implemented"
	return *new(TXTResult), nil
}

func (c *Client) lookupTXT(ctx context.Context, fqdn string, nameservers []string, recursive bool) (TXTResult, error) {
	_ = "STUB: not implemented"
	return *new(TXTResult), nil
}

func extractTXTRecords(msg *dns.Msg, name string) []TXTRecord {
	_ = "STUB: not implemented"
	return nil
}
