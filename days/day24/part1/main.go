package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type point struct {
	q, r, s int
}

var directions = map[string]point{
	"nw": {q: 0, r: -1, s: 1},
	"ne": {q: 1, r: -1, s: 0},
	"sw": {q: -1, r: 1, s: 0},
	"se": {q: 0, r: 1, s: -1},
	"w":  {q: -1, r: 0, s: 1},
	"e":  {q: 1, r: 0, s: -1},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: <binary> <test.txt>")
		os.Exit(1)
	}
	f := os.Args[1]
	content, err := os.ReadFile(f)
	if err != nil {
		fmt.Println("Failed to read file: ", err)
		os.Exit(1)
	}

	paths := make([][]string, 0)
	for _, line := range bytes.Split(content, []byte("\n")) {
		path := parsePath(string(line))
		paths = append(paths, path)
	}

	grid := move(paths)
	allBlacks := 0
	for _, v := range grid {
		if v == 1 {
			allBlacks++
		}
	}
	fmt.Println("all black tiles: ", allBlacks)
}

func move(paths [][]string) map[point]int {
	grid := make(map[point]int)
	for _, path := range paths {
		pos := point{q: 0, s: 0, r: 0}
		// fmt.Println("path: ", path)
		for _, p := range path {
			// fmt.Println("p: ", p)
			dir, ok := directions[p]
			if !ok {
				fmt.Println("Invalid direction: ", p)
				os.Exit(1)
			}
			pos.q += dir.q
			pos.r += dir.r
			pos.s += dir.s
			// fmt.Printf("pos: %+v\n", pos)
		}
		grid[pos] = grid[pos] ^ 1
	}
	return grid
}

func parsePath(s string) []string {
	var result []string
	split := strings.Split(s, "")
	for i := 0; i < len(split); i++ {
		r := string(split[i])
		if r == "n" || r == "s" {
			r += split[i+1]
			i++
		}
		result = append(result, r)
	}
	return result
}
