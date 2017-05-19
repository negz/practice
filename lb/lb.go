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
	w map[int]*Backend
}

func NewLoadBalancer(backends []*Backend) *LoadBalancer {
	rand.Seed(time.Now().UnixNano())
	lb := &LoadBalancer{w: make(map[int]*Backend)}
	i := 0
	for _, b := range backends {
		for j := 0; j < b.Weight; j++ {
			lb.w[i] = b
			i++
		}
	}
	return lb
}

func (lb *LoadBalancer) Next() *Backend {
	// Pick a random number from 0 -> totalWeight
	b := lb.w[rand.Intn(len(lb.w))]
	b.Handled++
	return b
}
