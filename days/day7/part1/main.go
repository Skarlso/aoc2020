package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type bag struct {
	name     string
	contains int
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, l := range split {
		lines = append(lines, l)
	}

	bags := make(map[string][]bag)
	for _, line := range lines {
		rules := strings.Split(line, "contain")
		bagUnSplit := rules[0]
		contain := strings.TrimSuffix(rules[1], ".")

		s := strings.Split(bagUnSplit, " ")
		bagName := strings.Join(s[:2], " ")

		containedBags := strings.Split(contain, ",")
		var addBags []bag

		// These will signify the end of the chain...
		if contain == " no other bags" {
			bags[bagName] = nil
			continue
		}
		for _, cbs := range containedBags {
			s := strings.Split(cbs, " ")
			count, _ := strconv.Atoi(s[1])
			name := strings.Join(s[2:4], " ")
			addBags = append(addBags, bag{name: name, contains: count})
		}

		bags[bagName] = addBags
	}

	// Go through all the bags and see if they can eventually end up containing a shiny gold bag.
	// If so... colors ++.

	colors := 0

	for k := range bags {
		if traverse("shiny gold", bags, k) {
			colors++
		}
	}

	fmt.Println("colors: ", colors)
}

func traverse(goal string, bags map[string][]bag, current string) bool {
	if bags[current] == nil {
		return false
	}
	check := make([]string, 0)
	for _, b := range bags[current] {
		if b.name == goal {
			return true
		}
		check = append(check, b.name)
	}
	// we haven't found it, check the other names
	for _, name := range check {
		if traverse(goal, bags, name) {
			return true
		}
	}
	return false
}
