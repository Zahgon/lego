package certcrypto

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"net"
	"time"

	"golang.org/x/crypto/ocsp"
)

const (
	OCSPGood = ocsp.Good

	OCSPRevoked = ocsp.Revoked

	OCSPUnknown = ocsp.Unknown

	OCSPServerFailed = ocsp.ServerFailed
)

var (
	tlsFeatureExtensionOID = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 1, 24}
	ocspMustStapleFeature  = []byte{0x30, 0x03, 0x02, 0x01, 0x05}
)

type DERCertificateBytes []byte

func ParsePEMBundle(bundle []byte) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParsePEMPrivateKey(key []byte) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

func GeneratePrivateKey(keyType KeyType) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

type CSROptions struct {
	Domain         string
	SAN            []string
	MustStaple     bool
	EmailAddresses []string
}

func CreateCSR(privateKey crypto.Signer, opts CSROptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PEMEncode(data any) []byte { _ = "STUB: not implemented"; return nil }

func PEMBlock(data any) *pem.Block { _ = "STUB: not implemented"; return nil }

func pemDecode(data []byte) (*pem.Block, error) { _ = "STUB: not implemented"; return nil, nil }

func PemDecodeTox509CSR(data []byte) (*x509.CertificateRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParsePEMCertificate(cert []byte) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetCertificateMainDomain(cert *x509.Certificate) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetCSRMainDomain(cert *x509.CertificateRequest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getMainDomain(subject pkix.Name, dnsNames []string, ips []net.IP) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ExtractDomains(cert *x509.Certificate) []string { _ = "STUB: not implemented"; return nil }

func ExtractDomainsCSR(csr *x509.CertificateRequest) []string {
	_ = "STUB: not implemented"
	return nil
}

func GeneratePemCert(privateKey *rsa.PrivateKey, domain string, extensions []pkix.Extension) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateDerCert(privateKey *rsa.PrivateKey, expiration time.Time, domain string, extensions []pkix.Extension) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
