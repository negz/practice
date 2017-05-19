package lb

import (
	"reflect"
	"sort"
	"testing"
)

var lbTests = []struct {
	requests int
	backends []*Backend
}{
	{
		requests: 1000,
		backends: []*Backend{
			&Backend{Name: "one", Weight: 300},
			&Backend{Name: "two", Weight: 700},
		},
	},
	{
		requests: 100,
		backends: []*Backend{
			&Backend{Name: "one", Weight: 200},
			&Backend{Name: "two", Weight: 800},
		},
	},
	{
		requests: 100,
		backends: []*Backend{
			&Backend{Name: "two", Weight: 800},
			&Backend{Name: "one", Weight: 200},
		},
	},
	{
		requests: 10000,
		backends: []*Backend{
			&Backend{Name: "one", Weight: 500},
			&Backend{Name: "two", Weight: 1000},
			&Backend{Name: "three", Weight: 600},
		},
	},
}

type ByWeight []*Backend

func (b ByWeight) Len() int           { return len(b) }
func (b ByWeight) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByWeight) Less(i, j int) bool { return b[i].Weight < b[j].Weight }

type ByHandled []*Backend

func (b ByHandled) Len() int           { return len(b) }
func (b ByHandled) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByHandled) Less(i, j int) bool { return b[i].Handled < b[j].Handled }

func TestBackendWeight(t *testing.T) {
	for _, tt := range lbTests {
		lb := NewLoadBalancer(tt.backends)

		for i := 0; i < tt.requests; i++ {
			lb.Next()
		}

		byWeight := make([]*Backend, len(tt.backends))
		copy(byWeight, tt.backends)
		sort.Sort(ByWeight(byWeight))

		byHandled := make([]*Backend, len(tt.backends))
		copy(byHandled, tt.backends)
		sort.Sort(ByHandled(byHandled))

		if !reflect.DeepEqual(byWeight, byHandled) {
			t.Errorf("LBs sorted by handled requests != LBs sorted by weight")
		}

		for _, b := range byWeight {
			t.Logf("LB %v", b)
		}
	}
}

func BenchmarkLoadBalancer(b *testing.B) {
	lb := NewLoadBalancer(lbTests[0].backends)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lb.Next()
	}
}
