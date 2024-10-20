package uow

import (
	"context"
	"reflect"
)

// UnitOfWorkFactory is a factory interface to create UnitOfWork instance.
type UnitOfWorkFactory interface {
	// New creates new UnitOfWork instance.
	New(ctx context.Context) (UnitOfWork, error)
}

// UnitOfWork is a interface of the unit of work pattern.
// The interface provides functionf to get a repository by an entitty and commit changes.
type UnitOfWork interface {
	// GetRepository gets a repository by the entity type.
	GetRepository(entityType reflect.Type) (Repository, error)
	// Committing saves all changes to the storage in one transaction.
	Commit(ctx context.Context) error
}

// Repository is a generic repository interface
// that provides functions for registering changes.
type Repository interface {
	// Add inserts new entity to the unit of work envronment.
	Add(ctx context.Context, e any) error
	// Update updates entity.
	Update(ctx context.Context, e any) error
	// Delete removes entity.
	Delete(ctx context.Context, e any) error
}

// UoWRepository is a repository interface specific to internal needs.
// The interface is used to perform repository-specific commit logic.
type UoWRepository interface {
	// Flush executes repository-specific commit logic.
	Flush(ctx context.Context) error
}

// RepositoryFactory is a factory interface to create repository instance.
type RepositoryFactory interface {
	// New creates new repository instance.
	New(entityType reflect.Type) (Repository, error)
}
