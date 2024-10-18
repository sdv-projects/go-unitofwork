package uow

// Specification interface.
// Implement the interface for the domain entity in the application project.
type Specification[Entity any] interface {
	// IsSatisfiedBy check if t is satisfied by the specification.
	IsSatisfiedBy(t *Entity) bool
}
