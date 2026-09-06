package certcrypto

import (
	"crypto"
	"crypto/elliptic"
	"crypto/x509"
	"math/big"
)

const (
	EC256   = KeyType("EC256")
	EC384   = KeyType("EC384")
	RSA2048 = KeyType("RSA2048")
	RSA3072 = KeyType("RSA3072")
	RSA4096 = KeyType("RSA4096")
	RSA8192 = KeyType("RSA8192")
)

type KeyType string

func (k KeyType) String() string { _ = "STUB: not implemented"; return "" }

func ToKeyType(keyType string) (KeyType, error) {
	_ = "STUB: not implemented"
	return *new(KeyType), nil
}

func AllKeyTypes() []KeyType { _ = "STUB: not implemented"; return nil }

func IsSupported(keyType KeyType) bool { _ = "STUB: not implemented"; return false }

func GetPrivateKeyType(signer crypto.Signer) (KeyType, error) {
	_ = "STUB: not implemented"
	return *new(KeyType), nil
}

func GetCertificateKeyType(cert *x509.Certificate) (KeyType, error) {
	_ = "STUB: not implemented"
	return *new(KeyType), nil
}

func GetCSRKeyType(csr *x509.CertificateRequest) (KeyType, error) {
	_ = "STUB: not implemented"
	return *new(KeyType), nil
}

func GetKeyType(key any) (KeyType, error) { _ = "STUB: not implemented"; return *new(KeyType), nil }

func getRSAKeyType(n *big.Int) (KeyType, error) {
	_ = "STUB: not implemented"
	return *new(KeyType), nil
}

func getECDSAKeyType(curve elliptic.Curve) (KeyType, error) {
	_ = "STUB: not implemented"
	return *new(KeyType), nil
}
