package engine

// Package defines a package name, it's dependencies on other packages, what it resolves, and the platforms it is targeted for.
type Package struct {
	Name         string
	Dependencies []string
	Resolves     []string
	Platforms    []string
}
