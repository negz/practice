package b62generator

import "testing"

var b62tests = []struct {
	next     uint64
	path     string
	overflow bool
}{
	{next: 1, path: "1"},
	{next: 61, path: "Z"},
	{next: 62, path: "10"},
	{next: 3521614606207, path: "ZZZZZZZ"},
	{next: 3521614606208, overflow: true},
	{next: 1000000, path: "4c92"},
}

func TestBase62Generator(t *testing.T) {
	g, err := New()
	if err != nil {
		t.Fatalf("New(): %v", err)
	}
	for _, tt := range b62tests {
		got, err := g.Generate(tt.next)
		if err != nil {
			if tt.overflow {
				continue
			}
			t.Errorf("g.Generate(%v): %v", tt.next, err)
			continue
		}
		if tt.path != got {
			t.Errorf("g.Generate(%v): want %#v, got %#v", tt.next, tt.path, got)
		}
	}
}
