package td

type pEntity interface{}
type void struct{}

// Empty struct to characterize that a unit is alive.
var alive void

// Group should associate to a pointer an empty struct iff the corresponding
// entity is alive. Dead entities must be removed from the set for garbage
// collection.
type group map[pEntity]void

// e should be a pointer towards a unit, projectile, etc. to be added to the
// group.
func (g group) add(e pEntity) {
	g[e] = alive
}

// Deletes a pointer to an existing entity from the group g. Noop if the entity
// is not part of the group.
func (g group) delete(e pEntity) {
	delete(g, e)
}
