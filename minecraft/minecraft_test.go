package minecraft

import "testing"

var minecraftTests = []struct {
	grid      Grid
	item      Item
	locations []Location
}{
	{
		grid:      NewGrid(0, 1).InsertItem(0, 0, NewTorch()),
		item:      NewTorch(),
		locations: []Location{Location{0, 0}},
	},
	{
		grid:      NewGrid(10, 10).InsertItem(5, 5, NewTorch()),
		item:      NewTorch(),
		locations: []Location{Location{5, 5}},
	},
	{
		grid:      NewGrid(100, 100).InsertItem(9, 9, NewPineappleTree()),
		item:      NewPineappleTree(),
		locations: []Location{Location{9, 9}},
	},
}

func TestMinecraft(t *testing.T) {
	for _, tt := range minecraftTests {
		found := tt.grid.FindItem(tt.item)

		for _, l := range tt.locations {
			if !found[l] {
				t.Errorf("%v not found at %v in %v", tt.item, l, tt.grid)
			}
		}
	}
}
