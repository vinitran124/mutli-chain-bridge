package main

import (
	"bridge/config"
	"embed"
	"fmt"
	"github.com/pressly/goose/v3"
	"github.com/urfave/cli/v2"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func beforeMigration(c *cli.Context) error {
	cfg, err := config.Load(c)
	if err != nil {
		return err
	}
	goose.SetBaseFS(embedMigrations)
	SetContextSQL(cfg.Database)
	return nil
}

func startMigration(c *cli.Context) error {
	migrateAction := c.String(config.FlagMigrateAction)
	switch migrateAction {
	case config.FlagUpAction:
		return goose.Up(GetContextSQL().DB, "migrations")
	case config.FlagDownAction:
		return goose.Down(GetContextSQL().DB, "migrations")
	default:
		return fmt.Errorf(`migration: invalid magration flags. "up" or "down" only`)
	}
}
