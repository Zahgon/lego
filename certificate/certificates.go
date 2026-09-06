package certificate

import (
	"context"
	"crypto"
	"crypto/x509"
	"time"

	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/acme/api"
	"github.com/go-acme/lego/v5/certcrypto"
	"golang.org/x/crypto/ocsp"
)

const (
	DefaultOverallRequestLimit = 18
)

const maxBodySize = 1024 * 1024

type Resource struct {
	ID      string   `json:"id"`
	Domains []string `json:"domains"`

	KeyType certcrypto.KeyType `json:"keyType"`

	PreferredChain string `json:"preferredChain,omitempty"`
	Profile        string `json:"profile,omitempty"`

	CertURL       string `json:"certUrl"`
	CertStableURL string `json:"certStableUrl"`

	PrivateKey        []byte `json:"-"`
	Certificate       []byte `json:"-"`
	IssuerCertificate []byte `json:"-"`
	CSR               []byte `json:"-"`
}

type ObtainRequest struct {
	Domains        []string
	MustStaple     bool
	EmailAddresses []string

	PrivateKey crypto.Signer
	KeyType    certcrypto.KeyType

	NotBefore        time.Time
	NotAfter         time.Time
	Bundle           bool
	PreferredChain   string
	EnableCommonName bool

	Profile string

	AlwaysDeactivateAuthorizations bool

	ReplacesCertID string
}

func (r ObtainRequest) EffectiveKeyType() (certcrypto.KeyType, error) {
	_ = "STUB: not implemented"
	return *new(certcrypto.KeyType), nil
}

type ObtainForCSRRequest struct {
	CSR *x509.CertificateRequest

	PrivateKey crypto.Signer

	NotBefore        time.Time
	NotAfter         time.Time
	Bundle           bool
	PreferredChain   string
	EnableCommonName bool

	Profile string

	AlwaysDeactivateAuthorizations bool

	ReplacesCertID string
}

func (r ObtainForCSRRequest) EffectiveKeyType() (certcrypto.KeyType, error) {
	_ = "STUB: not implemented"
	return *new(certcrypto.KeyType), nil
}

type resolver interface {
	Solve(ctx context.Context, authorizations []acme.Authorization) error
}

type CertifierOptions struct {
	Timeout             time.Duration
	OverallRequestLimit int
}

type Certifier struct {
	core                *api.Core
	resolver            resolver
	options             CertifierOptions
	overallRequestLimit int
}

func NewCertifier(core *api.Core, resolver resolver, options CertifierOptions) *Certifier {
	_ = "STUB: not implemented"
	return nil
}

func (c *Certifier) Obtain(ctx context.Context, request ObtainRequest) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Certifier) ObtainForCSR(ctx context.Context, request ObtainForCSRRequest) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Certifier) getForOrder(ctx context.Context, domains []string, order acme.ExtendedOrder, request ObtainRequest) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Certifier) getForCSR(ctx context.Context, certRes *Resource, order acme.ExtendedOrder, csr []byte, bundle bool, preferredChain string) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Certifier) checkResponse(ctx context.Context, certRes *Resource, order acme.ExtendedOrder, bundle bool, preferredChain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Certifier) Revoke(ctx context.Context, cert []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Certifier) RevokeWithReason(ctx context.Context, cert []byte, reason *uint) error {
	_ = "STUB: not implemented"
	return nil
}

type RenewOptions struct {
	NotBefore time.Time
	NotAfter  time.Time

	Bundle           bool
	PreferredChain   string
	EnableCommonName bool

	Profile string

	AlwaysDeactivateAuthorizations bool

	MustStaple     bool
	EmailAddresses []string
}

func (c *Certifier) Renew(ctx context.Context, certRes Resource, options *RenewOptions) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Certifier) GetOCSP(ctx context.Context, bundle []byte) ([]byte, *ocsp.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Certifier) Get(ctx context.Context, url string, bundle bool) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hasPreferredChain(issuer []byte, preferredChain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func checkOrderStatus(order acme.ExtendedOrder) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getObtainRequestPrivateKey(request ObtainRequest) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}
