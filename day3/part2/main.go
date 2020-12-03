package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const test = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(content), "\n")

	area := make([]string, 0)
	for _, l := range split {
		area = append(area, l)
	}

	slopes := [][]int{
		{
			1, 1,
		},
		{
			3, 1,
		},
		{
			5, 1,
		},
		{
			7, 1,
		},
		{
			1, 2,
		},
	}

	sum := 1
	for _, s := range slopes {
		right := s[0]
		down := s[1]
		trees := 0
		i := 0
		j := 0
		for {
			if i >= len(area) {
				break
			}
			if area[i][j] == '#' {
				trees++
			}

			j = (j + right) % len(area[i])
			i += down
		}
		sum *= trees
	}
	fmt.Println(sum)
}
