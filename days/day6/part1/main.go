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
		yess := make(map[rune]struct{})

		for _, ig := range g {
			for _, c := range ig {
				yess[c] = struct{}{}
			}
		}

		totalYes += len(yess)
	}
	fmt.Println(totalYes)
}
