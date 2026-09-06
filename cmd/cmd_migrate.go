package cmd

import (
	"context"
	"os"

	"github.com/go-acme/lego/v5/cmd/internal/configuration"
	"github.com/urfave/cli/v3"
)

const callToAction = `#######
#
# lego is an independent, free, and open-source project, if you value it, consider supporting it! ❤️
#
# https://donate.ldez.dev
#
#######

`

func createMigrate() *cli.Command { _ = "STUB: not implemented"; return nil }

func migration(_ context.Context, cmd *cli.Command) error { _ = "STUB: not implemented"; return nil }

func createConfigurationFile(root string, cfg *configuration.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func createSuggestedConfiguration(file *os.File, cfg *configuration.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func suggestedConfigurationFallback(cfg *configuration.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}
