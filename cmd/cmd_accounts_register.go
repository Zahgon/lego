package cmd

import (
	"context"

	"github.com/go-acme/lego/v5/acme"
	"github.com/go-acme/lego/v5/cmd/internal/storage"
	"github.com/go-acme/lego/v5/lego"
	"github.com/urfave/cli/v3"
)

func createAccountRegister() *cli.Command { _ = "STUB: not implemented"; return nil }

func register(ctx context.Context, cmd *cli.Command) error { _ = "STUB: not implemented"; return nil }

func handleRegistration(ctx context.Context, cmd *cli.Command, lazyClient lzSetUp, accountsStorage *storage.AccountsStorage, account *storage.Account, allowRegister bool) error {
	_ = "STUB: not implemented"
	return nil
}

func registerAccount(ctx context.Context, cmd *cli.Command, client *lego.Client) (*acme.ExtendedAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleTOS(cmd *cli.Command, client *lego.Client) bool { _ = "STUB: not implemented"; return false }

func updateAccountOrigin(account *storage.Account) { _ = "STUB: not implemented"; return }
