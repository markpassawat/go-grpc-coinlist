package db

import (
	bun "github.com/uptrace/bun"
)

// TestConnection will use 1+1 as test to check connectivity
func TestConnection(db *bun.DB) error {
	_, err := db.Exec("select 1+1")
	if err != nil {
		return err
	}
	return nil
}
