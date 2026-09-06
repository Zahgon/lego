package api

import (
	"context"

	"github.com/go-acme/lego/v5/acme"
)

type AuthorizationService service

func (c *AuthorizationService) Get(ctx context.Context, authzURL string) (acme.Authorization, error) {
	_ = "STUB: not implemented"
	return *new(acme.Authorization), nil
}

func (c *AuthorizationService) Deactivate(ctx context.Context, authzURL string) error {
	_ = "STUB: not implemented"
	return nil
}
