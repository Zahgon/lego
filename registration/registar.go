package registration

import (
	"context"
	"crypto"

	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/acme/api"
)

const mailTo = "mailto:"

type RegisterOptions struct {
	TermsOfServiceAgreed bool
}

type RegisterEABOptions struct {
	TermsOfServiceAgreed bool
	Kid                  string
	HmacEncoded          string
}

type Registrar struct {
	core *api.Core
	user User
}

func NewRegistrar(core *api.Core, user User) *Registrar { _ = "STUB: not implemented"; return nil }

func (r *Registrar) Register(ctx context.Context, options RegisterOptions) (*acme.ExtendedAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registrar) RegisterWithExternalAccountBinding(ctx context.Context, options RegisterEABOptions) (*acme.ExtendedAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registrar) QueryRegistration(ctx context.Context) (*acme.ExtendedAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registrar) UpdateRegistration(ctx context.Context, options RegisterOptions) (*acme.ExtendedAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registrar) DeleteRegistration(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registrar) ResolveAccountByKey(ctx context.Context) (*acme.ExtendedAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registrar) KeyRollover(ctx context.Context, newKey crypto.Signer) error {
	_ = "STUB: not implemented"
	return nil
}

func RegisterWithZeroSSL(ctx context.Context, r *Registrar, email string) (*acme.ExtendedAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
