package dnspersist01

import (
	"context"
)

func (c *Client) checkRecursiveNameserversPropagation(ctx context.Context, fqdn string, matcher RecordMatcher) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Client) checkAuthoritativeNameserversPropagation(ctx context.Context, fqdn string, matcher RecordMatcher) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Client) checkNameserversPropagationCustom(ctx context.Context, fqdn string, nameservers []string, matcher RecordMatcher, addPort, recursive bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
