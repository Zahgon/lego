package dns01

import (
	"context"

	"github.com/miekg/dns"
)

func (c *Client) resolveCNAME(ctx context.Context, fqdn string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Client) lookupCNAME(ctx context.Context, fqdn string) string {
	_ = "STUB: not implemented"
	return ""
}

func updateDomainWithCName(r *dns.Msg, fqdn string) string { _ = "STUB: not implemented"; return "" }
