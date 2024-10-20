package bun

import (
	uow "github.com/sdv-projects/go-unitofwork"
)

// CreateRepositoryFactoryFn is a delegate function
// that is used to create a RepositoryFactory instance in the UnitOfWork implementation.
type CreateRepositoryFactoryFn func(t *Transation) uow.RepositoryFactory
