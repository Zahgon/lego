package cmd

import (
	"time"

	"github.com/urfave/cli/v3"
)

func createArchives() *cli.Command { _ = "STUB: not implemented"; return nil }

func parseArchiveDate(filename string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}
