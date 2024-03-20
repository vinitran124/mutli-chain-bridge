package main

import (
	backend "bridge"
	"github.com/urfave/cli/v2"
	"os"
)

func versionCmd(*cli.Context) error {
	backend.PrintVersion(os.Stdout)
	return nil
}
