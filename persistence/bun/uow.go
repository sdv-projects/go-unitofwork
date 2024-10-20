package bun

import (
	"context"

	uow "github.com/sdv-projects/go-unitofwork"
	bun "github.com/uptrace/bun"
)

type UnitOfWork struct {
	uow.UoWBase

	tr *Transation
}

func (u *UnitOfWork) Commit(ctx context.Context) error {
	err := u.tr.Begin(ctx)
	if err != nil {
		return err
	}
	defer u.tr.GetTx().Rollback()

	if err := u.Flush(ctx); err != nil {
		return err
	}

	return u.tr.GetTx().Commit()
}

func NewUnitOfWork(factory uow.RepositoryFactory, t *Transation) *UnitOfWork {
	return &UnitOfWork{
		UoWBase: uow.NewUoWBase(factory),
		tr:      t,
	}
}

type UnitOfWorkFactory struct {
	db       *bun.DB
	delegate CreateRepositoryFactoryFn
}

func (f *UnitOfWorkFactory) New(ctx context.Context) (uow.UnitOfWork, error) {
	tr := NewTransation(f.db)
	repoFactory := f.delegate(tr)

	return &UnitOfWork{
		UoWBase: uow.NewUoWBase(repoFactory),
		tr:      tr,
	}, nil
}

func NewUnitOfWorkFactory(db *bun.DB, delegate CreateRepositoryFactoryFn) *UnitOfWorkFactory {
	return &UnitOfWorkFactory{
		db:       db,
		delegate: delegate,
	}
}
