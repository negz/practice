package lru

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var lruTests = []struct {
	max int
	s   []uint32
}{
	{
		max: 3,
		s:   []uint32{0, 1, 2, 3, 4},
	},
	{
		max: 1,
		s:   []uint32{0, 1, 2, 3, 4},
	},
	{
		max: 5,
		s:   []uint32{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	},
}

func key(i uint32) string {
	return fmt.Sprintf("key %d", i)
}

func TestLRU(t *testing.T) {
	for _, tt := range lruTests {
		rand.Seed(time.Now().UnixNano())
		c := New(tt.max)

		t.Run("Insert()", func(t *testing.T) {
			for _, i := range tt.s {
				c.Insert(key(i), i)
			}

			want := tt.s[len(tt.s)-tt.max:]
			got := c.ToSlice()
			if !reflect.DeepEqual(want, got) {
				t.Errorf("want %v, got %v", want, got)
			}
		})

		t.Run("Get()", func(t *testing.T) {
			backwards := make([]uint32, 0, len(tt.s))
			for i := len(tt.s) - 1; i >= 0; i-- {
				c.Get(key(tt.s[i]))
				backwards = append(backwards, tt.s[i])
			}

			want := backwards[:len(backwards)-(len(backwards)-tt.max)]
			got := c.ToSlice()
			if !reflect.DeepEqual(want, got) {
				t.Errorf("want %v, got %v", want, got)
			}
		})
	}
}
