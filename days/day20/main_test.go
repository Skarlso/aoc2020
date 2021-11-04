package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	i := &image{
		id: 1,
		pixels: [][]string{
			{"*", ".", ".", "*"},
			{"*", ".", ".", "*"},
			{"*", ".", ".", "*"},
			{"*", ".", ".", "*"},
		},
	}
	i.rotate()
	pixels := i.pixels
	assert.Equal(t, [][]string{
		{"*", "*", "*", "*"},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
		{"*", "*", "*", "*"},
	}, pixels)
	// rotating again
	i.rotate()
	pixels = i.pixels
	assert.Equal(t, [][]string{
		{"*", ".", ".", "*"},
		{"*", ".", ".", "*"},
		{"*", ".", ".", "*"},
		{"*", ".", ".", "*"},
	}, pixels)
}

func TestFlip(t *testing.T) {
	i := &image{
		id: 1,
		pixels: [][]string{
			{"*", "*", ".", "*"},
			{"*", ".", ".", "*"},
			{"*", ".", "*", "*"},
			{"*", ".", "*", "*"},
		},
	}
	i.flip()
	pixels := i.pixels
	assert.Equal(t, [][]string{
		{"*", ".", "*", "*"},
		{"*", ".", "*", "*"},
		{"*", ".", ".", "*"},
		{"*", "*", ".", "*"},
	}, pixels)
	// flip again
	i.flip()
	pixels = i.pixels
	assert.Equal(t, [][]string{
		{"*", "*", ".", "*"},
		{"*", ".", ".", "*"},
		{"*", ".", "*", "*"},
		{"*", ".", "*", "*"},
	}, pixels)
}

func TestRightSide(t *testing.T) {
	i := &image{
		id: 1,
		pixels: [][]string{
			{"*", "*", ".", "*"},
			{".", ".", ".", "*"},
			{"*", ".", "*", "."},
			{"*", ".", "*", "*"},
		},
	}
	side := i.right()
	assert.Equal(t, "**.*", side)
}

func TestLeftSide(t *testing.T) {
	i := &image{
		id: 1,
		pixels: [][]string{
			{"*", "*", ".", "*"},
			{".", ".", ".", "*"},
			{"*", ".", "*", "."},
			{"*", ".", "*", "*"},
		},
	}
	side := i.left()
	assert.Equal(t, "*.**", side)
}

func TestTopSide(t *testing.T) {
	i := &image{
		id: 1,
		pixels: [][]string{
			{"*", "*", ".", "."},
			{".", ".", ".", "*"},
			{"*", ".", "*", "."},
			{"*", ".", "*", "*"},
		},
	}
	side := i.top()
	assert.Equal(t, "**..", side)
}

func TestBottomSide(t *testing.T) {
	i := &image{
		id: 1,
		pixels: [][]string{
			{"*", "*", ".", "*"},
			{".", ".", ".", "*"},
			{"*", ".", "*", "."},
			{"*", ".", ".", "."},
		},
	}
	side := i.bottom()
	assert.Equal(t, "*...", side)
}

func TestSideChecking(t *testing.T) {

}
