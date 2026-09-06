package registration

import (
	"crypto"

	"github.com/go-acme/lego/v5/acme"
)

type User interface {
	GetEmail() string
	GetRegistration() *acme.ExtendedAccount
	GetPrivateKey() crypto.Signer
}
