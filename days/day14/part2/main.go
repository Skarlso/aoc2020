package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := bytes.Split(content, []byte("\n"))

	lines := make([][]byte, 0)
	for _, l := range split {
		lines = append(lines, l)
	}

	var mask []byte
	memory := make(map[int]int64)
	for _, l := range lines {
		if bytes.HasPrefix(l, []byte("mask")) {
			mask = bytes.Split(l, []byte(" = "))[1]
			continue
		}

		var (
			index  int
			number int
		)

		fmt.Sscanf(string(l), "mem[%d] = %d", &index, &number)

		// The values are NOT modified. Only the indexes are modified. So we
		// write the given value to the random indexes.

		// Take the original mask... apply a new pattern
		// That will come back with a new mask
		// Use that NEW mask to generate

		newMask := applyMask(index, mask)
		// fmt.Println(newMask)
		indexes := make([]int, 0)
		for i, c := range newMask {
			if c == 'X' {
				indexes = append(indexes, i)
			}
		}
		masks := generateMasks(indexes, newMask)
		for _, m := range masks {
			i, _ := strconv.ParseInt(string(m), 2, 64)
			memory[int(i)] = int64(number)
		}
	}

	var sum int64
	for _, v := range memory {
		sum += v
	}
	fmt.Println(sum)
}

// returns all masks based on the combinations from the X-es.
func generateMasks(indexes []int, originalMask []byte) [][]byte {
	masks := make([][]byte, 0)
	for _, c := range combine(len(indexes)) {
		result := make([]byte, len(originalMask))
		copy(result, originalMask)
		for i, r := range c {
			result[indexes[i]] = r
		}
		masks = append(masks, result)
	}

	return masks
}

// Needs 32 bit integer representation
func applyMask(n int, mask []byte) []byte {
	var result string
	binary := fmt.Sprintf("%036b", n)
	for i, c := range binary {
		if mask[i] == '0' {
			result += string(c)
		} else if mask[i] == '1' {
			result += "1"
		} else if mask[i] == 'X' {
			result += "X"
		}
	}
	return []byte(result)
}

// 0 = 000
// 1 = 001
// 2 = 010
// 3 = 011
// 4 = 100
// 5 = 101
// 6 = 110
// 7 = 111
func combine(n int) [][]byte {
	combinations := make([][]byte, 0)

	// 2 to the power of n
	for i := 0; i < 1<<n; i++ {
		binary := fmt.Sprintf("%0*b", n, i)
		combinations = append(combinations, []byte(binary))
	}

	return combinations
}
