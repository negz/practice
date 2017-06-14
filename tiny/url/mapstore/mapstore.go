package mapstore

import (
	"sync"

	"github.com/negz/practice/tiny/url"
	"github.com/pkg/errors"
)

type notFound struct {
	error
}

func (err *notFound) NotFound() {}

type taken struct {
	error
}

func (err *taken) Taken() {}

type store struct {
	i   uint64
	m   map[string]string
	mtx *sync.RWMutex
}

// New returns a new map store.
func New() (url.Store, error) {
	return &store{i: 0, m: make(map[string]string), mtx: &sync.RWMutex{}}, nil
}

func (s *store) GetIndex() (uint64, error) {
	return s.i, nil
}

func (s *store) Get(path string) (string, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	url, ok := s.m[path]
	if !ok {
		return "", errors.Wrapf(&notFound{}, "cannot find URL for path %s", path)
	}
	return url, nil
}

func (s *store) Put(path, url string, index uint64) (uint64, error) {
	if index != s.i {
		return s.i, errors.Wrap(&taken{}, "requested index already consumed")
	}
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.m[path] = url
	s.i++
	return s.i, nil
}
