package minecraft

import (
	"fmt"
	"math/rand"
	"time"
)

// Given an n*n grid of resources, find a torch, i.e. a coal resource directly
// above a wood resource.

type Resource int

const (
	Empty Resource = iota
	Coal
	Wood
	Pineapple
)

func (r Resource) String() string {
	switch r {
	case Coal:
		return "Coal"
	case Wood:
		return "Wood"
	case Pineapple:
		return "Pineapple"
	default:
		return "Empty"
	}
}

var Resources = []Resource{Empty, Coal, Wood, Pineapple}

type Location struct {
	x int
	y int
}

func (l Location) String() string {
	return fmt.Sprintf("(%v,%v)", l.x, l.y)
}

type Grid map[Location]Resource

func randomResource() Resource {
	rand.Seed(time.Now().UnixNano())
	return Resources[rand.Intn(len(Resources))]
}

// NewGrid randomly generates a grid with bounded x and y
func NewGrid(x, y int) Grid {
	g := Grid{}
	for i := 0; i < x; i++ {
		for n := 0; n < y; n++ {
			g[Location{i, n}] = randomResource()
		}
	}
	return g
}

// An Item is just a small grid.
type Item Grid

// A NewTorch looks like:
//
//   c
//   w
//
func NewTorch() Item {
	i := Item{}
	i[Location{0, 0}] = Coal
	i[Location{0, 1}] = Wood
	return i
}

// A NewPineappleTree looks like:
//
//  ppp
//   w
//
func NewPineappleTree() Item {
	i := Item{}
	i[Location{0, 0}] = Pineapple
	i[Location{1, 0}] = Pineapple
	i[Location{2, 0}] = Pineapple
	i[Location{1, 1}] = Wood
	return i
}

func (g Grid) InsertItem(x, y int, i Item) Grid {
	for l, r := range i {
		g[Location{x + l.x, y + l.y}] = r
	}
	return g
}

func (g Grid) TestItem(l Location, i Item) bool {
	r, ok := g[l]
	if !ok {
		// This location is outside the grid
		return false
	}

	if r != i[Location{0, 0}] {
		// This location does not match the item's (0,0) location
		return false
	}

	for il, ir := range i {
		r, ok := g[Location{l.x + il.x, l.y + il.y}]
		if !ok {
			// We can't match this item because one of its coords would be
			// outside the grid.
			return false
		}
		if r != ir {
			// The resource at this grid location does not match the resource at
			// the item location.
			return false
		}
	}

	// All of the item's locations matched with an existing grid location.
	return true
}

func (g Grid) FindItem(i Item) map[Location]bool {
	// This won't be super efficient because we'll test grid locations for a
	// match to the item's 0,0 even if we've previously accessed their resource
	// to test for a match against another location in the item.
	found := make(map[Location]bool)
	for l, _ := range g {
		if g.TestItem(l, i) {
			found[l] = true
		}
	}
	return found
}
