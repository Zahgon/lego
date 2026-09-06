package dns01

import (
	"context"
)

func (c *Client) checkRecursiveNameserversPropagation(ctx context.Context, fqdn, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Client) checkAuthoritativeNameserversPropagation(ctx context.Context, fqdn, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Client) checkNameserversPropagationCustom(ctx context.Context, fqdn, value string, nameservers []string, addPort, recursive bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
