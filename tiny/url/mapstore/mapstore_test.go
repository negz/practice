package mapstore

import "testing"

var mapStoreTests = []struct {
	path string
	url  string
}{
	{path: "a", url: "http://example.org"},
	{path: "b", url: "http://example.com"},
	{path: "c", url: "http://example.net"},
}

func TestMapStore(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Fatalf("New(): %v", err)
	}
	i, err := s.GetIndex()
	if err != nil {
		t.Fatalf("s.GetIndex(): %v", err)
	}
	for _, tt := range mapStoreTests {
		i, err = s.Put(tt.path, tt.url, i)
		if err != nil {
			t.Errorf("s.Put(%v, %v, %v): %v", tt.path, tt.url, i, err)
			continue
		}
		got, err := s.Get(tt.path)
		if err != nil {
			t.Errorf("s.Get(%v): %v", tt.path, err)
			continue
		}
		if got != tt.url {
			t.Errorf("s.Get(%v): want %v, got %v", tt.path, tt.url, got)
		}
	}
}
