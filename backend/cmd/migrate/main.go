package main

import (
	"bridge/app/utils"
	"embed"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

const (
	envPath = ".env"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func init() {
	if err := godotenv.Overload(strings.Split(envPath, ",")...); err != nil {
		fmt.Println("Load env error", err.Error())
	}
}
func main() {
	goose.SetBaseFS(embedMigrations)
	utils.SetContextSQL()
	app := &cli.App{
		Name:  "migrate",
		Usage: "database migration",
		Commands: []*cli.Command{
			commandUp(),
			commandDown(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func commandUp() *cli.Command {
	return &cli.Command{
		Name: "up",
		Action: func(c *cli.Context) error {
			return goose.Up(SQLRepository().DB, "migrations")
		},
	}
}

func commandDown() *cli.Command {
	return &cli.Command{
		Name: "down",
		Action: func(c *cli.Context) error {
			return goose.Down(SQLRepository().DB, "migrations")
		},
	}
}
