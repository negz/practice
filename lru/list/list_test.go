package list

import (
	"reflect"
	"testing"
)

var listTests = []struct {
	s []uint32
	t []uint32
	h []uint32
}{
	{
		s: []uint32{0, 1, 2, 3, 4},
		t: []uint32{0, 2, 3, 4, 1},
		h: []uint32{1, 2, 3, 4, 0},
	},
}

func TestList(t *testing.T) {
	for _, tt := range listTests {
		l := NewFromSlice(tt.s)

		t.Run("ToSlice", func(t *testing.T) {
			actual := l.ToSlice()
			if !reflect.DeepEqual(tt.s, actual) {
				t.Errorf("want %v, got %v", tt.s, actual)
			}
		})

		t.Run("MoveToTail", func(t *testing.T) {
			mtt := NewFromSlice(tt.s)
			mtt.MoveToTail(mtt.Head.Next)
			actual := mtt.ToSlice()
			if !reflect.DeepEqual(tt.t, actual) {
				t.Errorf("want %v, got %v", tt.t, actual)
			}
		})

		t.Run("MoveHeadToTail", func(t *testing.T) {
			mhtt := NewFromSlice(tt.s)
			mhtt.MoveToTail(mhtt.Head)
			actual := mhtt.ToSlice()
			if !reflect.DeepEqual(tt.h, actual) {
				t.Errorf("want %v, got %v", tt.h, actual)
			}
		})

		t.Run("Append", func(t *testing.T) {
			tt.s = append(tt.s, 42)
			l.Append(42)
			actual := l.ToSlice()
			if !reflect.DeepEqual(tt.s, actual) {
				t.Errorf("want %v, got %v", tt.s, actual)
			}
		})

		t.Run("TrimHead", func(t *testing.T) {
			l.TrimHead()
			actual := l.ToSlice()
			if !reflect.DeepEqual(tt.s[1:], actual) {
				t.Errorf("want %v, got %v", tt.s, actual)
			}
		})

		t.Run("DeleteTail", func(t *testing.T) {
			l.Delete(l.Tail)
			actual := l.ToSlice()
			if !reflect.DeepEqual(tt.s[1:len(tt.s)-1], actual) {
				t.Errorf("want %v, got %v", tt.s[1:len(tt.s)-1], actual)
			}
		})

		t.Run("DeleteAll", func(t *testing.T) {
			e := l.Head
			for e != nil {
				l.Delete(e)
				e = e.Next
			}
			actual := l.ToSlice()
			if !reflect.DeepEqual([]uint32{}, actual) {
				t.Errorf("want %v, got %v", []uint32{}, actual)
			}
		})
	}
}
