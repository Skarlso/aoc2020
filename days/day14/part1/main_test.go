package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyMask(t *testing.T) {
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	number := 11

	got := applyMask(number, mask)

	assert.Equal(t, int64(73), got)

	number = 101
	got = applyMask(number, mask)
	assert.Equal(t, int64(101), got)

	number = 0
	got = applyMask(number, mask)
	assert.Equal(t, int64(64), got)
}
