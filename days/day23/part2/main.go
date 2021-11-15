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
	head  *cup
	cache map[int]*cup
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
	if cc, ok := c.cache[label]; ok {
		return cc
	}

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

// up this sucks that had to duplicate it. Maybe there is some better way.
func (c *cup) search(label int) *cup {
	curr := c.next
	for {
		if curr.label == label {
			return curr
		}
		if curr.label == c.label {
			return nil
		}
		curr = curr.next
	}
}

func (c *cup) follow(n int) *cup {
	target := c

	for i := 0; i < n; i++ {
		target = target.next
	}

	return target
}

func (c *circle) shuffle() {
	// select the hand
	hand := c.head.next
	// update the circle to point to the following 4th after the 3rd cup
	c.head.next = c.follow(4)
	// create a new circle from the hand
	hand.follow(2).next = hand

	// get the circle's end 4 over from the current one
	// move the whole circle's ref to the DC
	// update the head to point to the new point
	// you don't remove and insert in a linked chain, you move the whole chain x times over.

	var (
		dc      *cup
		dcLabel = c.head.label
	)
	for {
		dcLabel--
		if dcLabel < 1 {
			dcLabel = max
		}
		// if the hand contains the target label, skip
		if ok := hand.search(dcLabel); ok != nil {
			continue
		}
		// fmt.Println("before: ", dcLabel)
		dc = c.search(dcLabel)
		// fmt.Println("after")
		break
	}

	// bring the hand back at the dc's location
	hand.follow(2).next = dc.next
	dc.next = hand

	// update the circle
	c.head = c.head.next
}

var (
	// data = []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	// moves = 100
	max = 1000000
	// cc  = 3
	data  = []int{1, 5, 8, 9, 3, 7, 4, 6, 2}
	moves = 10000000
	cc    = 1
)

func main() {
	// fmt.Println(data)
	for i := 10; i <= 1000000; i++ {
		data = append(data, i)
	}
	fmt.Println(data[len(data)-1])
	first := &cup{
		label: cc,
	}
	c := &circle{
		head:  first,
		cache: make(map[int]*cup),
	}
	last := first
	// Create the cups
	for _, d := range data[1:] {
		current := &cup{
			label: d,
		}
		last.next = current
		last = current
		c.cache[d] = current

	}
	last.next = first

	// fmt.Println(c.display(3))

	// I will need to cache somewhere.
	for i := 0; i < moves; i++ {
		c.shuffle()
	}

	next := c.search(1).next
	value := next.label * next.next.label
	fmt.Println(value)
}
