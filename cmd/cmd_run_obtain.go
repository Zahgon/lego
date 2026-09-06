package cmd

import (
	"context"

	"github.com/go-acme/lego/v5/cmd/internal/hook"
	"github.com/go-acme/lego/v5/cmd/internal/storage"
	"github.com/go-acme/lego/v5/lego"
	"github.com/urfave/cli/v3"
)

func obtain(ctx context.Context, cmd *cli.Command, certID string, lazyClient lzSetUp, certsStorage *storage.CertificatesStorage, hookManager *hook.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func obtainForDomains(ctx context.Context, cmd *cli.Command, client *lego.Client, certID string, certsStorage *storage.CertificatesStorage, hookManager *hook.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func obtainForCSR(ctx context.Context, cmd *cli.Command, client *lego.Client, certID string, certsStorage *storage.CertificatesStorage, hookManager *hook.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
