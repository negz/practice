package shortener

import (
	"time"

	"github.com/pkg/errors"

	"github.com/negz/practice/tiny/url"
)

type retries struct {
	count      uint
	initial    time.Duration
	multiplier float32
}

// A Shortener gets and creates short URLs!
type shortener struct {
	g url.Generator
	s url.Store
	i uint64
	r *retries
}

// An Option for creating a Shortener.
type Option func(*shortener) error

// WithRetries overrides the default retry logic used when attempting to create
// a URL that has already been consumed.
func WithRetries(count uint, initial time.Duration, multiplier float32) Option {
	return func(s *shortener) error {
		s.r = &retries{count, initial, multiplier}
		return nil
	}
}

// New returns a new shortener!
func New(g url.Generator, s url.Store, o ...Option) (url.Shortener, error) {
	// Get our starting index from the store.
	i, err := s.GetIndex()
	if err != nil {
		return nil, errors.Wrap(err, "cannot determine initial URL index")
	}
	sh := &shortener{g, s, i, &retries{5, 200 * time.Millisecond, 1.5}}
	for _, so := range o {
		if err := so(sh); err != nil {
			return nil, errors.Wrap(err, "cannot apply shortener option")
		}
	}
	return sh, nil
}

func (s *shortener) Get(path string) (string, error) {
	path, err := s.s.Get(path)
	if err != nil {
		return "", errors.Wrapf(err, "cannot get path %s", path)
	}
	return path, nil
}

func (s *shortener) Create(u string) (string, error) {
	// TODO(negz): Push this check and set logic down into the Store?
	sleep := s.r.initial
	for tries := s.r.count; tries > 0; tries-- {
		path, err := s.g.Generate(s.i)
		if err != nil {
			return "", errors.Wrap(err, "cannot generate URL path")
		}

		i, err := s.s.Put(path, u, s.i)
		if err == nil {
			s.i = i
			return path, nil
		}

		if !url.IsTaken(err) {
			return "", errors.Wrap(err, "cannot store URL")
		}

		// URL was taken. Backoff and retry.
		s.i = i
		time.Sleep(sleep)
		sleep = time.Duration(float32(sleep) * s.r.multiplier)
	}
	return "", errors.New("could not create unique URL path")
}
