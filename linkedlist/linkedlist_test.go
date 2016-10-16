package linkedlist

import "testing"

var llTests = []struct {
	s []int
	n []int
	i []int
}{
	{
		[]int{100, 1, 2, 3, 4, 5},
		[]int{100, 1, 32, 5, 6},
		[]int{5, 1, 100},
	},
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equalsSlice(e *Element, s []int, t *testing.T) {
	if e == nil && len(s) != 0 {
		t.Errorf("First element is nil")
		return
	}

	i := 0
	for e != nil {
		if i > len(s) {
			t.Errorf("List is longer than slice (len = %v)", len(s))
			return
		}
		if e.Get() != s[i] {
			t.Errorf("Want %v, got %v", s[i], e.Get())
		}

		e = e.Next()
		i++
	}
}

func TestLinkedList(t *testing.T) {
	for _, tt := range llTests {
		// Make a linked list from the slice
		l := LinkedListFromSlice(tt.s)

		// Make a reversed copy of the test slice
		rs := make([]int, len(tt.s))
		copy(rs, tt.s)
		reverseSlice(rs)

		t.Run("Create", func(t *testing.T) {
			equalsSlice(l, tt.s, t)
		})
		t.Run("Reverse", func(t *testing.T) {
			equalsSlice(Reverse(l), rs, t)
		})
		t.Run("RecurseReverse", func(t *testing.T) {
			equalsSlice(RecurseReverse(l), rs, t)
		})
		t.Run("Intersection", func(t *testing.T) {
			ln := LinkedListFromSlice(tt.n)
			equalsSlice(Intersection(l, ln), tt.i, t)
		})
	}
}
