package cmd

import (
	"context"
	"io"

	"github.com/urfave/cli/v3"
)

const flgCode = "code"

func createDNSHelp() *cli.Command { _ = "STUB: not implemented"; return nil }

func createDNSHelpFlags() []cli.Flag { _ = "STUB: not implemented"; return nil }

func dnsHelp(_ context.Context, cmd *cli.Command) error { _ = "STUB: not implemented"; return nil }

type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) writeln(a ...any) { _ = "STUB: not implemented"; return }

func (ew *errWriter) writef(format string, a ...any) { _ = "STUB: not implemented"; return }
