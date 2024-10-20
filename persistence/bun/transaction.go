package bun

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

// Transaction wraps the Bun transaction
// for use between UnitOfWork and repositories.
type Transation struct {
	db *bun.DB
	tx bun.Tx
}

// GetTx gets current openned Bun transaction.
func (t *Transation) GetTx() bun.Tx {
	return t.tx
}

// Begin starts Bun transaction.
func (t *Transation) Begin(ctx context.Context) (err error) {
	t.tx, err = t.db.BeginTx(ctx, &sql.TxOptions{})

	return
}

func NewTransation(db *bun.DB) *Transation {
	return &Transation{
		db: db,
	}
}
