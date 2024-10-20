package uow

import "errors"

var (
	// ErrRepositoryNotRegistred      = errors.New("repository not registred")
	ErrUoWRepositoryNotImplemented = errors.New("the repository don't implement UoWRepository")
)
