package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	E = iota
	S
	W
	N
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

	heading := 0
	var (
		e, s, w, n int
	)
	headings := []int{
		E, S, W, N,
	}
	for _, l := range lines {
		direction := string(l[0])
		a := l[1:]
		amount, _ := strconv.Atoi(a)

		switch direction {
		case "F":
			switch headings[heading] {
			case E:
				e += amount
			case S:
				s += amount
			case W:
				w += amount
			case N:
				n += amount
			}
		case "E":
			e += amount
		case "W":
			w += amount
		case "S":
			s += amount
		case "N":
			n += amount
		case "R":
			degree := amount / 90
			heading = modLikePython(heading+degree, len(headings))
		case "L":
			degree := amount / 90
			heading = modLikePython(heading-degree, len(headings))
		}
	}

	fmt.Println(e, w, s, n)
	fmt.Println(abs(e-w) + abs(s-n))
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
