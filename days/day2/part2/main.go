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

		var i, j int
		fmt.Sscanf(restrict, "%d-%d", &i, &j)

		if (string(password[i-1]) == letter || string(password[j-1]) == letter) && !(string(password[i-1]) == letter && string(password[j-1]) == letter) {
			valid++
		}

	}
	fmt.Println(valid)
}
