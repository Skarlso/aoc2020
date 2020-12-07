package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	name     string
	contains int
}

func main() {
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
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

	c := traverse(bags, "shiny gold", 1)
	// Why is there a -1? Because the last return somehow pushes a 0 through somewhere. Meh.
	fmt.Println("count: ", c-1)
}

func traverse(bags map[string][]bag, current string, count int) int {
	if bags[current] == nil {
		return count
	}

	c := 0
	for _, b := range bags[current] {
		c += b.contains * traverse(bags, b.name, count)
	}
	count += c
	return count
}
