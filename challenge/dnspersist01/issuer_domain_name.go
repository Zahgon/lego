package dnspersist01

import (
	"github.com/go-acme/lego/v5/acme"
	"golang.org/x/net/idna"
)

//nolint:gochecknoglobals // test seam for injecting IDNA conversion failures/variants.
var issuerDomainNameToASCII = idna.Lookup.ToASCII

func validateIssuerDomainNames(chlng acme.Challenge) error { _ = "STUB: not implemented"; return nil }

func validateIssuerDomainName(name string) error { _ = "STUB: not implemented"; return nil }

func isLDHLabel(label string) bool { _ = "STUB: not implemented"; return false }

func isLowerAlphaNum(c byte) bool { _ = "STUB: not implemented"; return false }

func normalizeUserSuppliedIssuerDomainName(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
