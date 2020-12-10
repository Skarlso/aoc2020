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

	ints := []int{0}
	max := 0
	for _, l := range lines {
		v, _ := strconv.Atoi(l)
		ints = append(ints, v)
		if v > max {
			max = v
		}
	}

	sort.Ints(ints)
	threeDiffs := 1 // starts at one because of max+3
	oneDiffs := 0

	for i := 0; i < len(ints)-1; i++ {
		if ints[i]+1 == ints[i+1] {
			oneDiffs++
		} else if ints[i]+3 == ints[i+1] {
			threeDiffs++
		}
	}

	fmt.Println("threes: ", threeDiffs)
	fmt.Println("ones: ", oneDiffs)
	fmt.Println("multi: ", oneDiffs*threeDiffs)
}
