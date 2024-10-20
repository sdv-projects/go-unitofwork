package bun

import (
	uow "github.com/sdv-projects/go-unitofwork"
	bun "github.com/uptrace/bun"
)

// RepositoryBase is a base struct with Bun-specific artifacts.
// It should be embeded into the entity repository implementation
// in your application.
type RepositoryBase struct {
	uow.RepositoryBase

	tr *Transation
}

// GetTx gets current openned Bun transaction.
// The function must be used in Flush(...) implementation.
func (r *RepositoryBase) GetTx() bun.Tx {
	return r.tr.GetTx()
}

func NewRepositoryBase(tr *Transation) RepositoryBase {
	return RepositoryBase{
		RepositoryBase: uow.NewRepositoryBase(),
		tr:             tr,
	}
}
