package lb2

import (
	"reflect"
	"sort"
	"testing"
)

var lbTests = []struct {
	backends []*Backend
	requests int
}{
	{
		backends: []*Backend{
			&Backend{weight: 10},
			&Backend{weight: 20},
			&Backend{weight: 30},
		},
		requests: 1000,
	},
	{
		backends: []*Backend{
			&Backend{weight: 100},
			&Backend{weight: 200},
			&Backend{weight: 300},
		},
		requests: 1000,
	},
	{
		backends: []*Backend{
			&Backend{weight: 100},
			&Backend{weight: 300},
		},
		requests: 1000,
	},
}

type ByWeight []*Backend

func (b ByWeight) Len() int           { return len(b) }
func (b ByWeight) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByWeight) Less(i, j int) bool { return b[i].weight < b[j].weight }

type ByHandled []*Backend

func (b ByHandled) Len() int           { return len(b) }
func (b ByHandled) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByHandled) Less(i, j int) bool { return b[i].handled < b[j].handled }

func TestLB(t *testing.T) {
	t.Run("Handle", func(t *testing.T) {
		for _, tt := range lbTests {
			if len(tt.backends) < 2 {
				t.Fatal("test cases must have at least 2 backends")
			}
			// Add all but the last LoadBalancer.
			lb := NewLoadBalancer(tt.backends[0 : len(tt.backends)-1])

			// Then add the last one!
			lb.AddBackend(tt.backends[len(tt.backends)-1])

			// Handle some requests.
			for i := 0; i < tt.requests; i++ {
				lb.Handle()
			}

			// Sort backends by weight
			byWeight := make([]*Backend, len(tt.backends))
			copy(byWeight, tt.backends)
			sort.Sort(ByWeight(byWeight))

			// Sort backends by how many requests they handled.
			byHandled := make([]*Backend, len(tt.backends))
			copy(byHandled, tt.backends)
			sort.Sort(ByHandled(byHandled))

			if !reflect.DeepEqual(byWeight, byHandled) {
				t.Errorf("\nwant %s\ngot  %s", byWeight, byHandled)
				for w, b := range lb.w {
					t.Logf("%d: %s", w, b)
				}
			}
		}
	})

	t.Run("RemoveBackend", func(t *testing.T) {
		for _, tt := range lbTests {
			if len(tt.backends) < 2 {
				t.Fatal("test cases must have at least 2 backends")
			}
			// Add all but the last LoadBalancer.
			lb := NewLoadBalancer(tt.backends)

			// Then remove the first one!
			lb.AddBackend(tt.backends[0])

			// Handle some requests.
			for i := 0; i < tt.requests; i++ {
				lb.Handle()
			}

			// Sort backends by weight
			byWeight := make([]*Backend, len(tt.backends))
			copy(byWeight, tt.backends)
			sort.Sort(ByWeight(byWeight))

			// Sort backends by how many requests they handled.
			byHandled := make([]*Backend, len(tt.backends))
			copy(byHandled, tt.backends)
			sort.Sort(ByHandled(byHandled))

			if !reflect.DeepEqual(byWeight, byHandled) {
				t.Errorf("\nwant %s\ngot  %s", byWeight, byHandled)
				for w, b := range lb.w {
					t.Logf("%d: %s", w, b)
				}
			}
		}
	})
}
