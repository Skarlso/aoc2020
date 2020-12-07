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

	c := traverse(bags, "shiny gold")
	fmt.Println("count: ", c-1)
}

func traverse(bags map[string][]bag, current string) int {
	c := 1
	for _, b := range bags[current] {
		c += b.contains * traverse(bags, b.name)
	}
	return c
}
