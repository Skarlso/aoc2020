package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type point struct {
	x, y int
}

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

	plane := make([][]string, 0)
	for _, l := range lines {
		s := strings.Split(l, "")
		plane = append(plane, s)
	}

	seats := make(map[point]string)
	for x, row := range plane {
		for y, col := range row {
			if col == "L" {
				seats[point{x: x, y: y}] = "L"
			}
		}
	}

	displayPlane(plane)
	for {
		noChange := true
		nplane := copyPlane(plane)

		for y := 0; y < len(plane); y++ {
			for x := 0; x < len(plane[y]); x++ {
				newValue := getNewValue(point{x: x, y: y}, plane)
				if plane[y][x] != newValue {
					noChange = false
				}
				nplane[y][x] = newValue
			}
		}

		// displayPlane(nplane)

		if noChange {
			break
		}

		plane = nplane
	}

	occupied := 0
	for _, y := range plane {
		for _, x := range y {
			if x == "#" {
				occupied++
			}
		}
	}

	fmt.Println(occupied)
}

func copyPlane(plane [][]string) [][]string {
	nplane := make([][]string, 0)
	for _, x := range plane {
		row := make([]string, len(x))
		copy(row, x)
		nplane = append(nplane, row)
	}
	return nplane
}

func displayPlane(plane [][]string) {
	for _, y := range plane {
		for _, x := range y {
			fmt.Print(x)
		}
		fmt.Println()
	}
	fmt.Println()
}

// x, y
var directions = [][]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
}

// returns the new value of the seat.
func getNewValue(p point, plane [][]string) string {
	if plane[p.y][p.x] == "." {
		return "."
	}
	var empty bool

	if plane[p.y][p.x] == "L" {
		empty = true
	}
	occupied := 0
	for _, d := range directions {
		x := p.x + d[0]
		y := p.y + d[1]

		if y < len(plane) && y >= 0 && x < len(plane[y]) && x >= 0 && plane[y][x] == "#" {
			occupied++
		}
	}

	if empty && occupied == 0 {
		return "#"
	} else if !empty && occupied >= 4 {
		return "L"
	}

	// no change
	return plane[p.y][p.x]
}
