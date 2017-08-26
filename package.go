package engine

type Package struct {
	Name         string
	Dependencies []string
	Resolves     []string
	Platforms    []string
}
