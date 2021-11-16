package main

import (
	"reflect"
	"testing"
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
