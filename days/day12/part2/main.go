package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type sboat struct {
	x, y int
}

type swaypoint struct {
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

	waypointDirectionX := map[string]int{
		"N": 0,
		"S": 0,
		"E": 1,
		"W": -1,
	}
	waypointDirectionY := map[string]int{
		"N": 1,
		"S": -1,
		"E": 0,
		"W": 0,
	}

	boat := sboat{}
	waypoint := swaypoint{
		x: 10,
		y: 1,
	}

	for _, l := range lines {
		direction := string(l[0])
		a := l[1:]
		amount, _ := strconv.Atoi(a)

		switch direction {
		case "E", "W", "S", "N":
			waypoint.x += waypointDirectionX[direction] * amount
			waypoint.y += waypointDirectionY[direction] * amount
		case "L":
			for i := 0; i < amount/90; i++ {
				waypoint.x, waypoint.y = -waypoint.y, waypoint.x
			}
		case "R":
			for i := 0; i < amount/90; i++ {
				waypoint.x, waypoint.y = waypoint.y, -waypoint.x
			}
		case "F":
			boat.x += waypoint.x * amount
			boat.y += waypoint.y * amount
		}
	}

	fmt.Println(abs(boat.x) + abs(boat.y))

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
