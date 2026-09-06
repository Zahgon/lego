package cmd

import (
	"crypto/x509"

	"github.com/go-acme/lego/v5/certificate"
	"github.com/go-acme/lego/v5/cmd/internal/configuration"
	"github.com/go-acme/lego/v5/cmd/internal/hook"
	"github.com/go-acme/lego/v5/cmd/internal/storage"
	"github.com/go-acme/lego/v5/lego"
	"github.com/go-acme/lego/v5/registration"
	"github.com/urfave/cli/v3"
)

type lzSetUp func() (*lego.Client, error)

func newClient(cmd *cli.Command, account registration.User) (*lego.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newClientConfig(cmd *cli.Command, account registration.User) *lego.Config {
	_ = "STUB: not implemented"
	return nil
}

func getUserAgentFromFlag(cmd *cli.Command) string { _ = "STUB: not implemented"; return "" }

func getUserAgent(cmd *cli.Command, ua string) string { _ = "STUB: not implemented"; return "" }

func newObtainRequest(cmd *cli.Command, domains []string) (certificate.ObtainRequest, error) {
	_ = "STUB: not implemented"
	return *new(certificate.ObtainRequest), nil
}

func newObtainForCSRRequest(cmd *cli.Command, csr *x509.CertificateRequest) certificate.ObtainForCSRRequest {
	_ = "STUB: not implemented"
	return *new(certificate.ObtainForCSRRequest)
}

func newSaveOptions(cmd *cli.Command) *storage.SaveOptions { _ = "STUB: not implemented"; return nil }

func newHookManager(cmd *cli.Command, certsStorage *storage.CertificatesStorage, account *storage.Account) *hook.Manager {
	_ = "STUB: not implemented"
	return nil
}

func parseAddress(cmd *cli.Command, flgName string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func loadConfiguration(cmd *cli.Command) (*configuration.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getConfigurationPath(cmd *cli.Command) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
