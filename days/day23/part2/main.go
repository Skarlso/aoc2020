package main

import (
	"fmt"
	"strconv"
)

type cup struct {
	label int
	next  *cup
}

type circle struct {
	head *cup
}

func (c *circle) display(from int) string {
	cp := c.search(from)
	next := cp.next
	var result string

	for next.label != cp.label {
		result += strconv.Itoa(next.label)
		next = next.next
	}

	return result
}

func (c *circle) search(label int) *cup {
	// cache some shit here
	curr := c.head
	for {
		if curr.label == label {
			return curr
		}
		curr = curr.next
	}
}

func (c *circle) follow(n int) *cup {
	target := c.head

	for i := 0; i < n; i++ {
		target = target.next
	}

	return target
}

func (c *circle) shuffle() {

}

var (
	data  = []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	moves = 10
	max   = 10
	// data  = []int{1, 5, 8, 9, 3, 7, 4, 6, 2}
	// moves = 10000000
	// cc    = 1
)

func main() {
	// fmt.Println(data)
	first := &cup{
		label: 3,
	}
	c := &circle{
		head: first,
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

	// result := first.Display(1)
	// fmt.Println(result)
	for i := 0; i < moves; i++ {
		c.shuffle()
	}
	result := c.display(1)
	fmt.Println(result)
}
