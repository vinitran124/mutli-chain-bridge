package database

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"os"
)

const (
	envDBDataSourcePassword = "POSTGRES_PASSWORD"
	envDBDataSourceUser     = "POSTGRES_USER"
	envDBDataSourceHost     = "DB_HOST"
	envDBDataSourcePort     = "POSTGRES_PORT"
	envDBDataSourceDB       = "POSTGRES_DB"
)

type SQLRepository struct {
	Client *bun.DB
}

func (r *SQLRepository) Init() error {
	password := os.Getenv(envDBDataSourcePassword)
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv(envDBDataSourceUser),
		os.Getenv(envDBDataSourcePassword),
		os.Getenv(envDBDataSourceHost),
		os.Getenv(envDBDataSourcePort),
		os.Getenv(envDBDataSourceDB),
	)

	log.Println(dsn)

	var db *bun.DB
	db = bun.NewDB(sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(dsn),
		pgdriver.WithPassword(password),
	)), pgdialect.New())

	r.Client = db
	return nil
}

func (r *SQLRepository) Close() error {
	return r.Client.Close()
}

func NewRepository() (*SQLRepository, error) {
	r := new(SQLRepository)
	if err := r.Init(); err != nil {
		return nil, err
	}

	return r, nil
}
