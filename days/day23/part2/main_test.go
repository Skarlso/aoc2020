package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGatherFromCup(t *testing.T) {
	d := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	s := gatherFromCup(4, d)
	assert.Equal(t, "254673891", s)
}

func TestFindDestinationCup(t *testing.T) {
	d := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	dc := findDestinationCup(3, d)
	assert.Equal(t, 2, dc)
}
