package lb2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Weight determines how much traffic a backend will recieve.
type Weight uint

// A Backend represents a backend who can handle our traffic.
type Backend struct {
	weight  Weight
	handled uint
}

// Handle handles a request.
func (b *Backend) Handle() {
	b.handled++
}

func (b *Backend) String() string {
	return fmt.Sprintf("Backend{weight: %d, handled: %d}", b.weight, b.handled)
}

// A LoadBalancer balances incoming requests across various backends.
type LoadBalancer struct {
	w map[Weight]*Backend
	m *sync.RWMutex
}

// NewLoadBalancer creates a new LoadBalancer from zero or more backends.
func NewLoadBalancer(backends []*Backend) *LoadBalancer {
	rand.Seed(time.Now().UnixNano())
	lb := &LoadBalancer{w: make(map[Weight]*Backend), m: &sync.RWMutex{}}
	for _, b := range backends {
		lb.AddBackend(b)
	}
	return lb
}

// AddBackend registers a new backend with the LoadBalancer.
func (lb *LoadBalancer) AddBackend(backend *Backend) {
	lb.m.Lock()
	defer lb.m.Unlock()
	if backend.weight < 1 {
		return
	}
	total := Weight(len(lb.w))
	for w := total; w < total+backend.weight; w++ {
		lb.w[w] = backend
	}
}

// RemoveBackend deregisters a backend from the LoadBalancer.
func (lb *LoadBalancer) RemoveBackend(backend *Backend) {
	lb.m.Lock()
	defer lb.m.Unlock()

	// Let's assume we usually want to keep at least two backends.
	keep := make([]*Backend, 0, 2)
	for _, b := range lb.w {
		if b == backend {
			// Don't keep the backend we want to remove.
			continue
		}
		keep = append(keep, b)
	}

	// Regenerate our index.
	lb.w = make(map[Weight]*Backend)
	for _, b := range keep {
		lb.AddBackend(b)
	}
}

// Handle selects a weighted random backend to handle a request.
func (lb *LoadBalancer) Handle() {
	lb.m.RLock()
	defer lb.m.RUnlock()
	choice := Weight(rand.Intn(len(lb.w) - 1))
	lb.w[choice].Handle()
}
