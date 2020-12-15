package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyMask(t *testing.T) {
	mask := []byte("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
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

// func TestCombination(t *testing.T) {
// 	a := []string{"0", "1"}
// 	combine(a, "", len(a), 3)
// 	log.Println(combinations)
// }

// func TestGenerateMasks(t *testing.T) {
// 	mask := "000000000000000000000000000000X1001X"
// 	indexes := make([]int, 0)
// 	for i, c := range mask {
// 		if c == 'X' {
// 			indexes = append(indexes, i)
// 		}
// 	}
// 	got := generateMasks(indexes, mask)
// 	want := []string{"000000000000000000000000000000010010", "000000000000000000000000000000010011", "000000000000000000000000000000110010", "000000000000000000000000000000110011"}
// 	assert.Equal(t, want, got)
// 	// Test if we can run again and get a different outcome so that the
// 	// global slice is overwritten

// 	mask = "000000000000000000000000000000X10011"
// 	indexes = make([]int, 0)
// 	for i, c := range mask {
// 		if c == 'X' {
// 			indexes = append(indexes, i)
// 		}
// 	}
// 	got = generateMasks(indexes, mask)
// 	want = []string{"000000000000000000000000000000010011", "000000000000000000000000000000110011"}
// 	assert.Equal(t, want, got)
// }

func TestNewCombine(t *testing.T) {
	ret := combine(3)
	log.Println(ret)
}
