package cmd

import (
	"log/slog"

	"github.com/go-acme/lego/v5/cmd/internal/configuration"
	"github.com/urfave/cli/v3"
)

const rfc3339NanoNatural = "2006-01-02T15:04:05.000000000Z07:00"

func setUpLogger(cmd *cli.Command, logCfg *configuration.Log) { _ = "STUB: not implemented"; return }

func getLogLeveler(lvl string) slog.Leveler { _ = "STUB: not implemented"; return *new(slog.Leveler) }
