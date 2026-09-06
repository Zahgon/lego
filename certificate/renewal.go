package certificate

import (
	"context"
	"crypto/x509"
	"time"

	"github.com/go-acme/lego/v5/acme"
)

type RenewalInfo struct {
	*acme.ExtendedRenewalInfo
}

func (r *RenewalInfo) ShouldRenewAt(now time.Time, willingToSleep time.Duration) *time.Time {
	_ = "STUB: not implemented"
	return nil
}

func (c *Certifier) GetRenewalInfo(ctx context.Context, cert *x509.Certificate) (*RenewalInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
