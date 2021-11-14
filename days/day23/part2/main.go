package main

import (
	"fmt"
	"strconv"
)

// maybe use this...
// I really don't want to build a tree.
type cup struct {
	label int
	next  *cup
}

func (c *cup) Display(from int) string {
	cp := c.find(from)
	next := cp.next
	var result string

	for next.label != cp.label {
		result += strconv.Itoa(next.label)
		next = next.next
	}

	return result
}

func (c *cup) find(label int) *cup {
	// cache some shit here
	curr := c
	for {
		if curr.label == label {
			return curr
		}
		curr = curr.next
	}
}

var (
	data = []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	// moves   = 10
	// data  = []int{1, 5, 8, 9, 3, 7, 4, 6, 2}
	// moves = 10000000
	// cc    = 1
)

func main() {
	// fmt.Println(data)
	first := &cup{
		label: 3,
	}
	last := first
	// Create the cups
	for _, d := range data[1:] {
		current := &cup{
			label: d,
		}
		last.next = current
		last = current

	}
	last.next = first

	result := first.Display(1)
	fmt.Println(result)
}
