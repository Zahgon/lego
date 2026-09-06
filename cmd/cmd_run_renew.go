package cmd

import (
	"context"
	"crypto/x509"
	"math"
	"time"

	"github.com/go-acme/lego/v5/cmd/internal/hook"
	"github.com/go-acme/lego/v5/cmd/internal/storage"
	"github.com/go-acme/lego/v5/lego"
	"github.com/urfave/cli/v3"
)

const noDays = -math.MaxInt

type renewProcessor struct {
	cmd *cli.Command

	lazyClient lzSetUp

	certsStorage *storage.CertificatesStorage
	hookManager  *hook.Manager
}

func (p *renewProcessor) renew(ctx context.Context, certID string, resource *storage.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *renewProcessor) renewForDomains(ctx context.Context, certID string, domains []string, changed bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *renewProcessor) renewForCSR(ctx context.Context, certID string, changed bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *renewProcessor) getARIInfo(ctx context.Context, certID string, cert *x509.Certificate) (*time.Time, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func getFlagRenewDays(cmd *cli.Command) int { _ = "STUB: not implemented"; return 0 }

func isInRenewalPeriod(cert *x509.Certificate, certID string, days int, now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func getDueDate(x509Cert *x509.Certificate, days int, now time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func getARIRenewalTime(ctx context.Context, willingToSleep time.Duration, cert *x509.Certificate, certID string, client *lego.Client) *time.Time {
	_ = "STUB: not implemented"
	return nil
}

func randomSleep(cmd *cli.Command) { _ = "STUB: not implemented"; return }

func merge(prevDomains, nextDomains []string) []string { _ = "STUB: not implemented"; return nil }

func sameDomainsCertificate(cert *x509.Certificate, csr *x509.CertificateRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func sameDomains(a, b []string) bool { _ = "STUB: not implemented"; return false }

func hasChanged(resource *storage.Certificate, cmd *cli.Command) bool {
	_ = "STUB: not implemented"
	return false
}
