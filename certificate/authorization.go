package certificate

import (
	"context"

	"github.com/go-acme/lego/v5/acme"
)

func (c *Certifier) getAuthorizations(ctx context.Context, order acme.ExtendedOrder) ([]acme.Authorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Certifier) deactivateAuthorizations(ctx context.Context, order acme.ExtendedOrder, force bool) {
	_ = "STUB: not implemented"
	return
}
