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

	ints := []int{0}
	max := 0
	for _, l := range lines {
		v, _ := strconv.Atoi(l)
		ints = append(ints, v)
		if v > max {
			max = v
		}
	}

	rating := max + 3
	ints = append(ints, rating)
	sort.Ints(ints)

	// path := make(map[int]struct{})
	// for _, i := range ints {
	// 	path[i] = struct{}{}
	// }
	// fmt.Println(traverse(0, path))
	fmt.Println(dp(0, ints))
}

var paths = make(map[int]int)

func traverse(current int, path map[int]struct{}) int {
	if c, ok := paths[current]; ok {
		return c
	}

	count := 0
	found := false
	for i := 1; i <= 3; i++ {
		next := current + i
		if _, ok := path[next]; ok {
			count += traverse(next, path)
			found = true
		}
	}

	if !found {
		count++
	}
	paths[current] = count
	return count
}

var cache = make(map[int]int)

func dp(i int, ints []int) int {
	if i == len(ints)-1 {
		return 1
	}
	if v, ok := cache[i]; ok {
		return v
	}
	count := 0
	for j := i + 1; j < len(ints); j++ {
		if ints[j]-ints[i] <= 3 {
			count += dp(j, ints)
		}
	}
	cache[i] = count
	return count
}
