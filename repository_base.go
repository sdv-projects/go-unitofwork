package uow

import (
	"context"
	"sync"
)

type RepositoryBase struct {
	added   []any
	deleted []any
	updated []any

	mu sync.Mutex
}

func (r *RepositoryBase) Add(ctx context.Context, e any) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.added = append(r.added, e)

	return nil
}

func (r *RepositoryBase) Update(ctx context.Context, e any) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.updated = append(r.updated, e)

	return nil
}

func (r *RepositoryBase) Delete(ctx context.Context, e any) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.deleted = append(r.deleted, e)

	return nil
}

func (r *RepositoryBase) GetAdded() []any {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.added
}

func (r *RepositoryBase) GetDeleted() []any {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.deleted
}

func (r *RepositoryBase) GetUpdated() []any {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.updated
}

func NewRepositoryBase() RepositoryBase {
	return RepositoryBase{
		added:   make([]any, 0),
		deleted: make([]any, 0),
		updated: make([]any, 0),
		mu:      sync.Mutex{},
	}
}
