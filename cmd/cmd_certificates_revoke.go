package cmd

import (
	"context"

	"github.com/go-acme/lego/v5/cmd/internal/storage"
	"github.com/go-acme/lego/v5/lego"
	"github.com/urfave/cli/v3"
)

func createRevoke() *cli.Command { _ = "STUB: not implemented"; return nil }

func revokeFromConfig(ctx context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func revoke(ctx context.Context, cmd *cli.Command) error { _ = "STUB: not implemented"; return nil }

func revokeCertificate(ctx context.Context, client *lego.Client, store *storage.Storage, certID string, reason uint, keep bool) error {
	_ = "STUB: not implemented"
	return nil
}
