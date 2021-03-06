package txutil

import (
	"fmt"
	"log"
)

type ITXHandler interface {
	Begin() (ITX, error)
}

type ITX interface {
	Commit() error
	Rollback() error
}

type TransactionBody func(tx ITX) error

// TXHandler is handler for working with transaction.
// This is wrapper function for commit and rollback.
func TXHandler(txHandler ITXHandler, f TransactionBody) error {
	var err error
	tx, err := txHandler.Begin()
	if err != nil {
		return fmt.Errorf("start transaction failed -> %w", err)
	}
	defer func() {
		if err := recover(); err != nil {
			rollBackErr := tx.Rollback()
			if rollBackErr != nil {
				log.Fatalf("rollback failed: %v", rollBackErr)
			}
			log.Println("Rollback operation")
			return
		}
	}()
	if err := f(tx); err != nil {
		return fmt.Errorf("transaction: operation failed -> %w", err)
	}
	return tx.Commit()
}
