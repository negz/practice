package lb

import (
	"math/rand"
	"time"
)

// Implement a load balancer that given a set of requests distributes those
// requests to an arbitrary number of weighted backends

type Backend struct {
	Name    string
	Weight  int
	Handled int
}

type LoadBalancer struct {
	b           []*Backend
	totalWeight int
}

func NewLoadBalancer(backends []*Backend) *LoadBalancer {
	rand.Seed(time.Now().UnixNano())
	lb := &LoadBalancer{b: backends, totalWeight: 0}
	for _, backend := range lb.b {
		lb.totalWeight += backend.Weight
	}
	return lb
}

func (lb *LoadBalancer) Next() *Backend {
	// Pick a random number from 0 -> totalWeight
	choice := rand.Intn(lb.totalWeight)

	currentWeight := 0
	// If that number falls in an LB's 'weight range' pick it
	for _, backend := range lb.b {
		currentWeight += backend.Weight
		if choice <= currentWeight {
			backend.Handled++
			return backend
		}
	}
	return nil
}
