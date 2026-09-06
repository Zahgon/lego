package challenge

import (
	"github.com/go-acme/lego/v5/acme"
)

type Type string

const (
	HTTP01 = Type("http-01")

	DNS01 = Type("dns-01")

	DNSPersist01 = Type("dns-persist-01")

	TLSALPN01 = Type("tls-alpn-01")
)

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

func FindChallenge(chlgType Type, authz acme.Authorization) (acme.Challenge, error) {
	_ = "STUB: not implemented"
	return *new(acme.Challenge), nil
}

func GetTargetedDomain(authz acme.Authorization) string { _ = "STUB: not implemented"; return "" }
