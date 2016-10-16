package quicksort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func randSlice(n int) []int {
	r := make([]int, n)
	for i := range r {
		r[i] = rand.Intn(100)
	}
	return r
}

var sortTests = []struct {
	length int
}{
	{10},
	{20},
	{1},
	{2},
	{4},
}

func TestSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for _, tt := range sortTests {
		data := randSlice(tt.length)

		t.Log("Sorting ", data)
		qsort(data)
		if !sort.IntsAreSorted(data) {
			t.Errorf("%v is not sorted", data)
		}
	}
}
