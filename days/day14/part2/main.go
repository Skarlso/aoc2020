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

		result := applyMask(number, mask)
		memory[index] = result
	}

	var sum int64
	for _, v := range memory {
		sum += v
	}
	fmt.Println(sum)
}

// Needs 32 bit integer representation
func applyMask(n int, mask string) int64 {
	var result string
	binary := fmt.Sprintf("%036b", n)
	for i, c := range binary {
		if mask[i] != 'X' {
			result += string(mask[i])
		} else {
			result += string(c)
		}
	}
	i, _ := strconv.ParseInt(result, 2, 64)
	return i
}
