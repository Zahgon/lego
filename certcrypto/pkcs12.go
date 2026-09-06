package certcrypto

import (
	"software.sslmate.com/src/go-pkcs12"
)

const (
	PKCS12LegacyDES  = "DES"
	PKCS12LegacyRC2  = "RC2"
	PKCS12Modern2023 = "SHA256"
	PKCS12Modern2026 = "PBMAC1"
)

func AllPKCS12Formats() []string { _ = "STUB: not implemented"; return nil }

func IsPKCS12Supported(format string) bool { _ = "STUB: not implemented"; return false }

func GetPKCS12Encoder(format string) (*pkcs12.Encoder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
