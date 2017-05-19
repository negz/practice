package boolands

import "testing"

var boolandsTests = []struct {
	g Grid
	i int
}{
	{
		g: NewZeroGrid(5, 5).
			Set(0, 0, true).
			Set(2, 2, true).
			Set(2, 3, true),
		i: 2,
	},
	{
		g: NewZeroGrid(5, 5).
			Set(3, 3, true).
			Set(2, 3, true),
		i: 1,
	},
	{
		g: NewZeroGrid(5, 5).
			Set(2, 1, true).
			Set(4, 4, true),
		i: 2,
	},
	{
		g: NewZeroGrid(3, 3).
			Set(0, 0, true).
			Set(0, 1, true).
			Set(0, 2, true).
			Set(1, 0, true).
			Set(1, 1, true).
			Set(1, 2, true).
			Set(2, 0, true).
			Set(2, 1, true).
			Set(2, 2, true),
		i: 1,
	},
	{
		g: NewZeroGrid(0, 0),
		i: 0,
	},
	{
		g: NewZeroGrid(1, 1),
		i: 0,
	},
	{
		g: NewZeroGrid(1, 1).Set(0, 0, true),
		i: 1,
	},
}

func TestIslands(t *testing.T) {
	for _, tt := range boolandsTests {
		got := tt.g.Islands()
		if got != tt.i {
			t.Errorf("tt.g.Islands(): want %v, got %v", tt.i, got)
		}
	}
}
