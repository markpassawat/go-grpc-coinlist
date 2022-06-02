package db

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// DB defines database interface.
type DB interface {
	orm.DB
	Begin() (*pg.Tx, error)
	RunInTransaction(ctx context.Context, fn func(*pg.Tx) error) error
}

// TestConnection will use 1+1 as test to check connectivity
func TestConnection(db *pg.DB) error {
	queryResult := 0
	if _, err := db.QueryOne(pg.Scan(&queryResult), `SELECT 1+1`); err != nil || queryResult != 2 {
		return err
	}
	return nil
}
