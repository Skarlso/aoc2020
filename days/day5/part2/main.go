package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

// var test = "FFFBBBFRRR"

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, l := range split {
		lines = append(lines, l)
	}

	rows := make([]int, 0)
	for i := 0; i < 128; i++ {
		rows = append(rows, i)
	}

	columns := make([]int, 0)
	for i := 0; i < 8; i++ {
		columns = append(columns, i)
	}

	seats := make([]int, 0)
	for _, line := range lines {
		var (
			row, column int
		)
		rws := make([]int, len(rows))
		copy(rws, rows)
		firstHalf := line[:7]
		lastHalf := line[7:]

		// do the binary search
		for _, c := range firstHalf {
			i := len(rws) / 2
			if c == 'F' {
				rws = rws[:i]
			} else if c == 'B' {
				rws = rws[i:]
			}
		}
		if len(rws) != 1 {
			log.Fatal("Ops.. ", rws)
		}
		row = rws[0]

		// do the other binary search
		cols := make([]int, len(columns))
		copy(cols, columns)
		for _, c := range lastHalf {
			i := len(cols) / 2
			if c == 'L' {
				cols = cols[:i]
			} else if c == 'R' {
				cols = cols[i:]
			}
		}

		column = cols[0]
		seat := (row * 8) + column
		seats = append(seats, seat)
	}

	sort.Ints(seats)
	for i := 0; i < len(seats)-1; i++ {
		if seats[i]+1 != seats[i+1] {
			fmt.Println(seats[i-1], seats[i], seats[i+1])
			break
		}
	}
}
