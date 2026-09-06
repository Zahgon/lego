package cmd

import (
	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/lego"
	"github.com/urfave/cli/v3"
)

func setupChallenges(cmd *cli.Command, client *lego.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func setupHTTPProvider(cmd *cli.Command, client *lego.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func createHTTPProvider(cmd *cli.Command) (challenge.Provider, error) {
	_ = "STUB: not implemented"
	return *new(challenge.Provider), nil
}

func setupTLSProvider(cmd *cli.Command, client *lego.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func setupDNS(cmd *cli.Command, client *lego.Client) error { _ = "STUB: not implemented"; return nil }

func setupDNSPersist(cmd *cli.Command, client *lego.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func getNetworkStack(cmd *cli.Command) challenge.NetworkStack {
	_ = "STUB: not implemented"
	return *new(challenge.NetworkStack)
}
