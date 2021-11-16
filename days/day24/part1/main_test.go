package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePath(t *testing.T) {
	testCases := []struct {
		path string
		want []string
	}{
		{
			path: "esenee",
			want: []string{"e", "se", "ne", "e"},
		},
		{
			path: "esew",
			want: []string{"e", "se", "w"},
		},
		{
			path: "nwwswee",
			want: []string{"nw", "w", "sw", "e", "e"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(tt *testing.T) {
			got := parsePath(tc.path)
			if !reflect.DeepEqual(got, tc.want) {
				tt.Errorf("wanted: %+v got: %+v", tc.want, got)
			}
		})
	}
}

func TestMoving(t *testing.T) {
	paths := [][]string{
		{"nw", "w", "sw", "e", "e"},
	}
	grid := move(paths)
	assert.Equal(t, 1, grid[point{0, 0}])
}
