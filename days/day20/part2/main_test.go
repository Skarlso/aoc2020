package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	i := &tile{
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
	i := &tile{
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
	i := &tile{
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
	i := &tile{
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
	i := &tile{
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
	i := &tile{
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

func TestConstructSea(t *testing.T) {
	image = [][]*tile{
		{
			{
				id: 1, pixels: [][]string{
					{"*", "*", "*", "*"},
					{".", ".", ".", "."},
					{"*", ".", "*", "."},
					{"p", "p", "p", "p"}},
			},
		},
		{
			{
				id: 2, pixels: [][]string{
					{"*", ".", ".", "*"},
					{".", ".", ".", "."},
					{".", ".", "*", "*"},
					{"l", "l", "l", "l"}},
			},
		},
		{
			{
				id: 3, pixels: [][]string{
					{"*", "*", "*", "*"},
					{"*", "*", "*", "*"},
					{"*", "*", "*", "*"},
					{"o", "o", "o", "o"}},
			},
		},
		{
			{
				id: 4, pixels: [][]string{
					{"-", "-", "-", "-"},
					{"c", "c", "c", "c"},
					{"v", "v", "v", "v"},
					{"a", "a", "a", "a"}},
			},
		},
	}
	sea := constructSea()
	want := []string{
		"*****..*****----",
		"........****cccc",
		"*.*...******vvvv",
		"ppppllllooooaaaa",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}
	assert.Equal(t, want, sea)
}
