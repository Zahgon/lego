package api

import (
	"context"
	"net/http"

	"github.com/go-acme/lego/v5/acme"
)

const maxBodySize = 1024 * 1024

type CertificateService service

func (c *CertificateService) Get(ctx context.Context, certURL string, bundle bool) (*acme.RawCertificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CertificateService) GetAll(ctx context.Context, certURL string, bundle bool) (map[string]*acme.RawCertificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CertificateService) Revoke(ctx context.Context, req acme.RevokeCertMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CertificateService) get(ctx context.Context, certURL string, bundle bool) (*acme.RawCertificate, http.Header, error) {
	_ = "STUB: not implemented"
	return nil, *new(http.Header), nil
}

func (c *CertificateService) getCertificateChain(cert []byte, bundle bool) *acme.RawCertificate {
	_ = "STUB: not implemented"
	return nil
}
