package b62generator

import (
	"github.com/negz/practice/tiny/url"
	"github.com/pkg/errors"
)

// TODO(negz): Make this (and thus the byte buffer size) configurable.
const maxNext = 3521614606207
const digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type b62 struct {
	b [7]byte
}

// New returns a URL Generator that generates 7 character base62 URL paths.
func New() (url.Generator, error) {
	return &b62{}, nil
}

func (g *b62) Generate(next uint64) (string, error) {
	if next > maxNext {
		return "", errors.Errorf("next integer %d would overflow 7 byte URL path", next)
	}
	// Create a new 7 byte buffer.
	b := g.b

	// Convert next to a base62 string.
	base := uint64(len(digits))
	i := len(b)
	for next > 0 {
		i--
		b[i] = digits[next%base]
		next /= base
	}

	return string(b[i:]), nil
}
