package url

// A Store persists a URL to path mapping.
type Store interface {
	// Get a URL, given a short path.
	Get(path string) (string, error)

	// Put a short path to URL mapping. Returns the next index that should be
	// used to generate a path. Stores must return an error that satisfies
	// IsTaken() *and* the next available index if the requested index has
	// already been used by another writer.
	Put(path, url string, index uint64) (uint64, error)

	// GetIndex returns the next index that should be used to generate a path.
	GetIndex() (uint64, error)
}

// IsTaken determines whether an error indicates an index was already taken.
// It does this by walking down the stack of errors built by pkg/errors and
// returning true for the first error that implements the following interface:
//
// type notfounder interface {
//   Taken()
// }
func IsTaken(err error) bool {
	if err == nil {
		return false
	}
	for {
		if _, ok := err.(interface {
			Taken()
		}); ok {
			return true
		}
		if c, ok := err.(interface {
			Cause() error
		}); ok {
			err = c.Cause()
			continue
		}
		return false
	}
}
