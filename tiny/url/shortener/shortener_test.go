package shortener

import (
	"testing"

	"github.com/negz/practice/tiny/url/b62generator"
	"github.com/negz/practice/tiny/url/mapstore"
)

var shortenerTests = []struct {
	url string
}{
	{url: "http://example.org/"},
	{url: "http://example.org/404"},
}

func TestShortener(t *testing.T) {
	st, err := mapstore.New()
	if err != nil {
		t.Fatalf("mapstore.New(): %v", err)
	}
	g, err := b62generator.New()
	if err != nil {
		t.Fatalf("b62generator.New(): %v", err)
	}
	s, err := New(g, st)
	if err != nil {
		t.Fatalf("NewShortener(%v, %v): %v", g, st, err)
	}

	for _, tt := range shortenerTests {
		path, err := s.Create(tt.url)
		if err != nil {
			t.Errorf("s.Create(%v): %v", tt.url, err)
			continue
		}
		got, err := s.Get(path)
		if err != nil {
			t.Errorf("s.Get(%v): %v", path, err)
			continue
		}

		if tt.url != got {
			t.Errorf("s.Get(%v): want %v, got %v", path, tt.url, got)
		}

	}

}
