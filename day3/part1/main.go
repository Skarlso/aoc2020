package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(content), "\n")

	area := make([]string, 0)
	for _, l := range split {
		area = append(area, l)
	}

	right := 3
	down := 1
	side := 0
	bottom := 0
	trees := 0
	for bottom < len(area)-1 {
		if area[bottom][side] == '#' {
			trees++
		}

		side = (side + right) % len(area[bottom])
		bottom += down
	}
	fmt.Println(trees)
}
