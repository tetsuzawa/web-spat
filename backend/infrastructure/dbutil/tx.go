package dbutil

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// TXHandler is handler for working with transaction.
// This is wrapper function for commit and rollback.
func TXHandler(db *sqlx.DB, f func(*sqlx.Tx) error) error {
	var err error
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("start transaction failed -> %w", err)
	}
	defer func() {
		if err := recover(); err != nil {
			rollBackErr := tx.Rollback()
			if rollBackErr != nil {
				log.Fatalf("rollback failed: %v", rollBackErr)
			}
			log.Print("Rollback operation")
			return
		}
	}()
	if err := f(tx); err != nil {
		return fmt.Errorf("transaction: operation failed -> %w", err)
	}
	return nil
}
