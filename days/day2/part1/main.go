package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")

	split := strings.Split(string(content), "\n")

	valid := 0
	for _, l := range split {
		instructions := strings.Split(l, " ")
		restrict := instructions[0]
		letter := strings.Trim(instructions[1], ":")
		password := instructions[2]

		var min, max int
		fmt.Sscanf(restrict, "%d-%d", &min, &max)

		count := strings.Count(password, letter)
		if count >= min && count <= max {
			valid++
		}
	}
	fmt.Println(valid)
}
