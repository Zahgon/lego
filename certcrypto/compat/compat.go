package compat

import (
	"github.com/go-acme/lego/v5/certcrypto"
)

const (
	EC256   = KeyTypeCompat(certcrypto.EC256)
	EC384   = KeyTypeCompat(certcrypto.EC384)
	RSA2048 = KeyTypeCompat(certcrypto.RSA2048)
	RSA3072 = KeyTypeCompat(certcrypto.RSA3072)
	RSA4096 = KeyTypeCompat(certcrypto.RSA4096)
	RSA8192 = KeyTypeCompat(certcrypto.RSA8192)
)

type KeyTypeCompat certcrypto.KeyType

func (k *KeyTypeCompat) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }
