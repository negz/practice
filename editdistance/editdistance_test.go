package editdistance

import "testing"

var distanceTests = []struct {
	i string
	j string
	d int
}{
	{"kitten", "sitting", 3},
	{"sitting", "kitten", 3},
	{"cat", "dog", 3},
	{"hog", "dog", 1},
	{"frog", "frogfrog", 4},
	{"hog", "frog", 2},
	{"frog", "log", 2},
	{"intention", "execution", 5},
}

func TestDistance(t *testing.T) {
	for _, tt := range distanceTests {
		d := distance(tt.i, tt.j)
		if d != tt.d {
			t.Errorf("distance(%v, %v) Got %v, want %v", tt.i, tt.j, d, tt.d)
		}
	}
}
