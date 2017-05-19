// Boolean islands. Given an arbitrarily sized 2D grid of 0s and 1s, determine
// how many contiguous islands of 1s exist.

package boolands

import (
	"math/rand"
	"time"
)

type Location struct {
	X int
	Y int
}

type Grid struct {
	m map[Location]bool
	x int
	y int
}

func NewRandomGrid(x, y int) Grid {
	rand.Seed(time.Now().UnixNano())
	m := make(map[Location]bool)
	// Think of Y as horizontal axis, X as vertical
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			m[Location{X: i, Y: j}] = rand.Intn(1) != 0
		}
	}
	return Grid{m, x, y}
}

func NewZeroGrid(x, y int) Grid {
	m := make(map[Location]bool)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			m[Location{X: i, Y: j}] = false
		}
	}
	return Grid{m, x, y}
}

func (g Grid) Set(x, y int, b bool) Grid {
	g.m[Location{x, y}] = b
	return g
}

func (g Grid) existingAdjacentIsland(visited []Location) bool {
	for _, v := range visited {
		if a, ok := g.m[v]; ok && a {
			// This adjacent location exists and is part of a counted
			// island. Continue without incrementing.
			return true
		}
	}
	return false
}

func (g Grid) Islands() int {
	islands := 0
	for i := 0; i < g.x; i++ {
		for j := 0; j < g.y; j++ {
			if !g.m[Location{i, j}] {
				// This location is a 0, continue without incrementing.
				continue
			}

			visited := []Location{
				Location{i, j - 1},     // Directly left
				Location{i - 1, j - 1}, // Above left
				Location{i - 1, j + 1}, // Above right
				Location{i - 1, j},     // Directly above
			}

			if g.existingAdjacentIsland(visited) {
				continue
			}

			// This location is a 1, and none of its visited adjacent neighbors are
			// part of an existing island.
			islands++
		}
	}
	return islands
}
