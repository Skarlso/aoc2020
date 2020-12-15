package main

import (
	"fmt"
	"strconv"
	"strings"
)

var test = "0,3,6"
var test1 = "2,1,3"
var test2 = "1,2,3"
var test3 = "2,3,1"
var input = "0,14,1,3,7,9"

type number struct {
	turn           int
	lastTurn       int
	mostRecentTurn int
	occurred       int
}

func (n *number) String() string {
	return fmt.Sprintf("turn: %d, lastTurn: %d, mostRecentTurn: %d, occurred: %d\n", n.turn, n.lastTurn, n.mostRecentTurn, n.occurred)
}

func main() {
	split := strings.Split(input, ",")

	numbers := make(map[int]*number)
	turn := 1
	last := -1
	for _, s := range split {
		i, _ := strconv.Atoi(s)
		numbers[i] = &number{turn: turn, mostRecentTurn: turn, occurred: 0, lastTurn: turn}
		last = i
		turn++
	}

	for turn <= 30000000 { // part 1 2020
		if n, ok := numbers[last]; ok {
			if n.occurred == 0 {
				last = 0
			} else if n.occurred > 0 {
				last = n.mostRecentTurn - n.lastTurn
			}
		}

		if n, ok := numbers[last]; !ok {
			// we said a new number
			numbers[last] = &number{turn: turn, mostRecentTurn: turn, occurred: 1, lastTurn: turn}
		} else {
			n.lastTurn, n.mostRecentTurn = n.mostRecentTurn, turn
			n.occurred++
		}
		turn++
	}

	fmt.Println("Last number: ", last)
}
