package api

import (
	"context"
	"crypto"

	"github.com/go-acme/lego/v5/acme"
)

type AccountService service

func (a *AccountService) New(ctx context.Context, req acme.Account) (acme.ExtendedAccount, error) {
	_ = "STUB: not implemented"
	return *new(acme.ExtendedAccount), nil
}

func (a *AccountService) NewEAB(ctx context.Context, req acme.Account, kid, hmacEncoded string) (acme.ExtendedAccount, error) {
	_ = "STUB: not implemented"
	return *new(acme.ExtendedAccount), nil
}

func (a *AccountService) Get(ctx context.Context, accountURL string) (acme.Account, error) {
	_ = "STUB: not implemented"
	return *new(acme.Account), nil
}

func (a *AccountService) Update(ctx context.Context, accountURL string, req acme.Account) (acme.Account, error) {
	_ = "STUB: not implemented"
	return *new(acme.Account), nil
}

func (a *AccountService) Deactivate(ctx context.Context, accountURL string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AccountService) KeyChange(ctx context.Context, newKey crypto.Signer) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeEABHmac(hmacEncoded string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
