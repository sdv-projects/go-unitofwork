package uow

import (
	"context"
	"reflect"
)

type UoWBase struct {
	factory RepositoryFactory
	rMap    map[reflect.Type]Repository
	rList   []UoWRepository
}

func (u *UoWBase) GetRepository(entityType reflect.Type) (Repository, error) {
	if repo, ok := u.rMap[entityType]; ok {
		return repo, nil
	}

	repo, err := u.factory.New(entityType)
	if err != nil {
		return nil, err
	}

	uowRepo, ok := repo.(UoWRepository)
	if !ok {
		return nil, ErrUoWRepositoryNotImplemented
	}

	u.rMap[entityType] = repo
	u.rList = append(u.rList, uowRepo)

	return repo, nil
}

func (u *UoWBase) Flush(ctx context.Context) error {
	for _, r := range u.rList {
		err := r.Flush(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewUoWBase(factory RepositoryFactory) UoWBase {
	return UoWBase{
		rMap:    make(map[reflect.Type]Repository),
		rList:   make([]UoWRepository, 0),
		factory: factory,
	}
}
