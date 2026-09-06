package api

import (
	"context"
	"crypto/x509"
	"errors"

	"github.com/go-acme/lego/v5/acme"
)

var ErrNoARI = errors.New("renewalInfo[get/post]: server does not advertise a renewal info endpoint")

func (c *CertificateService) GetRenewalInfo(ctx context.Context, certID string) (*acme.ExtendedRenewalInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MakeARICertID(leaf *x509.Certificate) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
