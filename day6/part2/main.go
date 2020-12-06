package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, l := range split {
		lines = append(lines, l)
	}

	groups := make([][]string, 0)
	group := make([]string, 0)
	for _, l := range lines {
		if l == "" {
			groups = append(groups, group)
			group = make([]string, 0)
			continue
		}

		group = append(group, l)
	}

	// add the rest
	groups = append(groups, group)

	totalYes := 0
	for _, g := range groups {
		// Read group. Count characters. The number of
		// yes will be the same letter in which every group
		// had it. So it's the letter which is there in every group.
		// If there is only one group, then it's all unique letters.

		yess := make(map[rune]int)
		yes := 0

		for _, ig := range g {
			for _, c := range ig {
				yess[c]++
			}
		}

		for _, y := range yess {
			if y == len(g) {
				yes++
			}
		}
		totalYes += yes
	}
	fmt.Println(totalYes)
}
