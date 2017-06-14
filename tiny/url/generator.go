package url

// A Generator generates short paths.
type Generator interface {
	// Generate the next available short path given the next available integer.
	Generate(next uint64) (string, error)
}
