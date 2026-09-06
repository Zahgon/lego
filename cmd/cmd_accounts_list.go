package cmd

import (
	"context"

	"github.com/go-acme/lego/v5/cmd/internal/storage"
	"github.com/urfave/cli/v3"
)

type ListAccount struct {
	*storage.Account

	Path string `json:"path,omitempty"`
}

func createAccountsList() *cli.Command { _ = "STUB: not implemented"; return nil }

func listAccounts(_ context.Context, cmd *cli.Command) error { _ = "STUB: not implemented"; return nil }

func listAccountsText(basePath string) error { _ = "STUB: not implemented"; return nil }

func listAccountsJSON(basePath string) error { _ = "STUB: not implemented"; return nil }

func readAccounts(basePath string) ([]ListAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
