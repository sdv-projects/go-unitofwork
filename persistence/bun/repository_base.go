package bun

import (
	uow "github.com/sdv-projects/go-unitofwork"
	bun "github.com/uptrace/bun"
)

type RepositoryBase struct {
	uow.RepositoryBase

	tr *Transation
}

func (r *RepositoryBase) GetTx() bun.Tx {
	return r.tr.GetTx()
}

func NewRepositoryBase(tr *Transation) RepositoryBase {
	return RepositoryBase{
		RepositoryBase: uow.NewRepositoryBase(),
		tr:             tr,
	}
}
