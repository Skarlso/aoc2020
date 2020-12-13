package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, l := range split {
		lines = append(lines, l)
	}

	heading := 1
	headingsX := []int{0, 1, 0, -1}
	headingsY := []int{1, 0, -1, 0}
	var (
		x, y int
	)
	directionsX := map[string]int{
		"N": 0,
		"S": 0,
		"E": 1,
		"W": -1,
	}
	directionsY := map[string]int{
		"N": 1,
		"S": -1,
		"E": 0,
		"W": 0,
	}
	for _, l := range lines {
		direction := string(l[0])
		a := l[1:]
		amount, _ := strconv.Atoi(a)

		switch direction {
		case "F":
			x += headingsX[heading] * amount
			y += headingsY[heading] * amount
		case "E", "W", "S", "N":
			x += directionsX[direction] * amount
			y += directionsY[direction] * amount
		case "R":
			degree := amount / 90
			heading = modLikePython(heading+degree, 4)
		case "L":
			degree := amount / 90
			heading = modLikePython(heading-degree, 4)
		}
	}

	fmt.Println(abs(x) + abs(y))
}

func modLikePython(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
