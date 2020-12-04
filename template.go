package main

import (
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

}
