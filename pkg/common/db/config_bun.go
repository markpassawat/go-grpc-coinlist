package db

import (
	"context"
	"database/sql"
	"fmt"

	"bytes"

	// "github.com/Forward-Protocol/APH-event-service/cmd/hello_world/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Config struct {
	HostDB     string `required:"true"`
	PortDB     string `required:"true"`
	UserDB     string `required:"true"`
	PasswordDB string `required:"true"`
	NameDB     string `required:"true"`
	PoolSize   int    `split_words:"true"`
	AppName    string `split_words:"true"`
	SSLEnable  string `split_words:"true" envconfig:"SSLENABLE"`
}

func get_ctx() context.Context {
	ctx := context.Background()
	return ctx
}

func StringConcat(cfg *Config) string {
	var bufString bytes.Buffer
	bufString.WriteString("postgres://")
	bufString.WriteString(cfg.UserDB)
	bufString.WriteString(":")
	bufString.WriteString(cfg.PasswordDB)
	bufString.WriteString("@")
	bufString.WriteString(cfg.HostDB)
	bufString.WriteString(":")
	bufString.WriteString(cfg.PortDB)
	bufString.WriteString("/")
	bufString.WriteString(cfg.NameDB)
	bufString.WriteString("?sslmode=")
	bufString.WriteString(cfg.SSLEnable)
	// fmt.Println("string concat : ", bufString)
	return bufString.String()
}

func connect(Path string) *bun.DB {

	// Open a PostgreSQL database.
	dsn := Path
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return db
}

func (cfg *Config) MustConnectDB() (*bun.DB, context.Context) {
	path := StringConcat(cfg)
	fmt.Println("path : ", path)
	db := connect(path)
	ctx := get_ctx()
	return db, ctx
}

func ConnectDatabase() (db *bun.DB) {

	// dsn := "postgres://fmlypwnvvuddlt:16dcd95388b05793ac810566390c0c0bda498b4c897de5242740ca44df2cc68c@ec2-52-21-136-176.compute-1.amazonaws.com:5432/d1dkb33aaj0sof?sslmode=require"
	dsn := "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db = bun.NewDB(sqldb, pgdialect.New())

	return 
}