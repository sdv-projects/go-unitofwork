package bun

import (
	uow "github.com/sdv-projects/go-unitofwork"
)

type CreateRepositoryFactoryFn func(t *Transation) uow.RepositoryFactory
