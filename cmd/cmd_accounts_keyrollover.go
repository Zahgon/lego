package cmd

import (
	"context"
	"crypto"

	"github.com/go-acme/lego/v5/certcrypto"
	"github.com/urfave/cli/v3"
)

func createAccountKeyRollover() *cli.Command { _ = "STUB: not implemented"; return nil }

func accountKeyRollover(ctx context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func getPrivateKey(cmd *cli.Command, keyType certcrypto.KeyType) (crypto.Signer, certcrypto.KeyType, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), *new(certcrypto.KeyType), nil
}
