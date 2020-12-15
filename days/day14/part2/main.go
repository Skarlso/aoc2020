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

var bits = []string{"0", "1"}

// returns all masks based on the combinations from the X-es.
func generateMasks(indexes []int, originalMask []byte) [][]byte {
	combinations = make([][]byte, 0)
	combine(bits, "", len(bits), len(indexes))
	masks := make([][]byte, 0)

	for _, c := range combinations {
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

var combinations [][]byte

func combine(a []string, prefix string, n, k int) {
	if k == 0 {
		combinations = append(combinations, []byte(prefix))
		return
	}
	for i := 0; i < n; i++ {
		newPrefix := prefix + a[i]
		combine(a, newPrefix, n, k-1)
	}
}
