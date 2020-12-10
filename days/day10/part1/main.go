package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, l := range split {
		lines = append(lines, l)
	}

	ints := make([]int, 0)
	max := 0
	for _, l := range lines {
		v, _ := strconv.Atoi(l)
		ints = append(ints, v)
		if v > max {
			max = v
		}
	}

	rating := max + 3

	// pick +1 -> try if every adapter is used, then pick +2 try again... pick +3 try again...
	// backtracking.
	sort.Ints(ints)
	threeDiffs := 1 // starts at one because of max+3
	oneDiffs := 0

	for i := 0; i < len(ints)-1; i++ {
		if i == 0 && ints[i] == 1 {
			oneDiffs++
		}

		if ints[i]+1 == ints[i+1] {
			oneDiffs++
			continue
		}
		if ints[i]+2 == ints[i+1] {
			continue
		}
		if ints[i]+3 == ints[i+1] {
			threeDiffs++
			continue
		}

		if ints[i] >= rating {
			fmt.Println("we reached our max adapter")
			break
		}
	}

	fmt.Println("threes: ", threeDiffs)
	fmt.Println("ones: ", oneDiffs)
	fmt.Println("multi: ", oneDiffs*threeDiffs)
}
