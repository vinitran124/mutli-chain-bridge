package main

import (
	"bridge/cmd/api"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

const (
	envPath = ".env"
)

func init() {
	if err := godotenv.Overload(strings.Split(envPath, ",")...); err != nil {
		fmt.Println("Load env error", err.Error())
	}
}

func main() {
	app := &cli.App{
		Name:  "Bridge",
		Usage: "Bridge",
		Action: func(*cli.Context) error {
			fmt.Println("use --help")
			return nil
		},
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			api.NewCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("[Main] Run CLI error:", err.Error())
		return
	}
}
