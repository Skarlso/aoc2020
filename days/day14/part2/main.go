package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, l := range split {
		lines = append(lines, l)
	}

	var mask string
	memory := make(map[int]int64)
	for _, l := range lines {
		if strings.HasPrefix(l, "mask") {
			mask = strings.Split(l, " = ")[1]
			continue
		}

		var (
			index  int
			number int
		)

		fmt.Sscanf(l, "mem[%d] = %d", &index, &number)

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
			i, _ := strconv.ParseInt(m, 2, 64)
			// newIndex := applyMask(index, m)
			// fmt.Println("New Index: ", i)
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
func generateMasks(indexes []int, originalMask string) []string {
	combinations = make([]string, 0)
	combine(bits, "", len(bits), len(indexes))
	masks := make([]string, 0)

	for _, c := range combinations {
		result := []rune(originalMask)
		for i, r := range c {
			result[indexes[i]] = r
		}
		masks = append(masks, string(result))
	}

	return masks
}

// Needs 32 bit integer representation
func applyMask(n int, mask string) string {
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
	return result
}

var combinations []string

func combine(a []string, prefix string, n, k int) {
	if k == 0 {
		combinations = append(combinations, prefix)
		return
	}
	for i := 0; i < n; i++ {
		newPrefix := prefix + a[i]
		combine(a, newPrefix, n, k-1)
	}
}
