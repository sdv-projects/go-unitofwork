package bun

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type Transation struct {
	db *bun.DB
	tx bun.Tx
}

func (t *Transation) GetTx() bun.Tx {
	return t.tx
}

func (t *Transation) Begin(ctx context.Context) (err error) {
	t.tx, err = t.db.BeginTx(ctx, &sql.TxOptions{})

	return
}

func NewTransation(db *bun.DB) *Transation {
	return &Transation{
		db: db,
	}
}
