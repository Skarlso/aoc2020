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

func TestHasMatchingSide(t *testing.T) {
	tests := map[string]struct {
		main  *tile
		other *tile
		want  bool
	}{
		"no rotation": {
			main: &tile{
				id: 1,
				pixels: [][]string{
					{"*", "*", ".", "*"},
					{".", ".", ".", "*"},
					{"*", ".", "*", "."},
					{"*", ".", ".", "."},
				},
			},
			other: &tile{
				id: 2,
				pixels: [][]string{
					{"*", "*", ".", "*"},
					{".", ".", ".", "*"},
					{"*", ".", "*", "."},
					{"*", ".", ".", "."},
				},
			},
			want: true,
		},
		"one rotation": {
			main: &tile{
				id: 1,
				pixels: [][]string{
					{"*", ".", ".", "."},
					{"*", ".", ".", "."},
					{"*", ".", ".", "."},
					{"*", ".", ".", "."},
				},
			},
			other: &tile{
				id: 2,
				pixels: [][]string{
					{"-", "-", "-", "-"},
					{"-", ".", ".", "-"},
					{"-", ".", ".", "-"},
					{"*", "*", "*", "*"},
				},
			},
			want: true,
		},
		"two rotations": {
			main: &tile{
				id: 1,
				pixels: [][]string{
					{"*", ".", ".", "."},
					{"*", ".", ".", "."},
					{"*", ".", ".", "."},
					{"*", ".", ".", "."},
				},
			},
			other: &tile{
				id: 2,
				pixels: [][]string{
					{"-", "-", "-", "*"},
					{"-", ".", ".", "*"},
					{"-", ".", ".", "*"},
					{"-", "-", "-", "*"},
				},
			},
			want: true,
		},
		"one flip": {
			main: &tile{
				id: 1,
				pixels: [][]string{
					{"*", ".", ".", "."},
					{"*", ".", ".", "."},
					{".", ".", ".", "."},
					{"*", ".", ".", "."},
				},
			},
			other: &tile{
				id: 2,
				pixels: [][]string{
					{"*", "-", "-", "*"},
					{".", ".", ".", "*"},
					{"*", ".", ".", "*"},
					{"*", "-", "-", "*"},
				},
			},
			want: true,
		},
		// "one flip and a rotations": {},
		"no matching side": {
			main: &tile{
				id: 1,
				pixels: [][]string{
					{"*", ".", ".", "."},
					{"*", ".", ".", "."},
					{"*", ".", ".", "."},
					{"*", ".", ".", "."},
				},
			},
			other: &tile{
				id: 2,
				pixels: [][]string{
					{"-", "-", "-", "-"},
					{"-", ".", ".", "-"},
					{"-", ".", ".", "-"},
					{"-", "-", "-", "-"},
				},
			},
			want: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.main.hasMatchingSideWith(tc.other)
			if tc.want != got {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
