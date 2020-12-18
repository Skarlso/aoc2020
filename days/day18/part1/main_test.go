package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	s := "1 + 2 + 3 * 5"
	n := solve(s)
	assert.Equal(t, 30, n)
	s = "10 + 15 + 5 + 6"
	n = solve(s)
	assert.Equal(t, 36, n)
	s = "1+2"
	n = solve(s)
	assert.Equal(t, 3, n)
	s = "11+22"
	n = solve(s)
	assert.Equal(t, 33, n)
}

func TestSolveParanthesis(t *testing.T) {
	s := "1 + 2 + (3 * 5)"
	n := solveParanthesis(s)
	assert.Equal(t, "1 + 2 + 15", n)
}
