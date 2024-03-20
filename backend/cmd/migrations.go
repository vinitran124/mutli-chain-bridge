package main

import (
	"bridge/app/utils"
	"bridge/config"
	"embed"
	"fmt"
	"github.com/pressly/goose/v3"
	"github.com/urfave/cli/v2"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func beforeMigration(c *cli.Context) error {
	goose.SetBaseFS(embedMigrations)
	utils.SetContextSQL()
	return nil
}

func startMigration(c *cli.Context) error {
	migrateAction := c.String(config.FlagMigrateAction)
	switch migrateAction {
	case config.FlagUpAction:
		return goose.Up(utils.GetContextSQL().Client.DB, "migrations")
	case config.FlagDownAction:
		return goose.Down(utils.GetContextSQL().Client.DB, "migrations")
	default:
		return fmt.Errorf(`migration: invalid magration flags. "up" or "down" only`)
	}
}
