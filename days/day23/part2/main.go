package main

import (
	"fmt"
	"strconv"
)

// maybe use this...
// I really don't want to build a tree.
type cup struct {
	label int
	value int
}

var (
	data = []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	// moves   = 10
	cc      = 3 // index of the current cup
	ccIndex = 0
	// data  = []int{1, 5, 8, 9, 3, 7, 4, 6, 2}
	moves = 10000000
	// cc    = 1
)

func main() {
	// fmt.Println(data)
	for i := 0; i < 1000000; i++ {
		data = append(data, i+9)
	}
	cm := 0 // current moves
	// I should know the index and could possibly eliminate some findIndexOfs... +3 or -3 or start from there perhaps and not from 0.
	for cm < moves {
		// take three
		three := make([]int, 0)
		threeIndex := make([]int, 0)
		// ccIndex := findIndexOfCup(cc, data)
		for i := 0; i < 3; i++ {
			ci := (i + ccIndex + 1) % len(data)
			three = append(three, data[ci])
			// fmt.Printf("ci: %d, value: %d\n", ci, data[ci])
			threeIndex = append(threeIndex, ci)
		}
		// delete all three with one append. wouldn't work if it wraps around...
		// for _, c := range three {
		// 	// this would be a bit better if this would be a tree, but I don't care for now.

		// 	// Of course, the index will be the same from upward.. so just delete +3 from the first index of the three.
		// 	// unless it wraps around.
		// 	index := findIndexOfCup(c, data, -1)
		// 	fmt.Printf("index: %d, value: %d\n", index, c)
		// 	data = append(data[:index], data[index+1:]...)
		// }

		data = append(data[:threeIndex[0]], data[threeIndex[0]+1:]...)
		if threeIndex[1]-1 < 0 {
			data = append(data[:0], data[1:]...)
		} else {
			data = append(data[:threeIndex[1]-1], data[(threeIndex[1]-1)+1:]...)
		}
		if threeIndex[2]-2 < 0 {
			data = append(data[:0], data[1:]...)
		} else {
			data = append(data[:threeIndex[2]-2], data[(threeIndex[2]-2)+1:]...)
		}

		// for _, c := range threeIndex {
		// 	// if any of the indexes is greater than the one coming after it, it won't fall into place
		// 	data = append(data[:c], data[c+1:]...)
		// }
		dc := findDestinationCup(cc, data, ccIndex)
		data = append(data[:dc+1], append(three, data[dc+1:]...)...)
		// fmt.Println(len(data))
		ccIndex = findIndexOfCup(cc, data, -1)
		ccIndex = (ccIndex + 1) % len(data)
		cc = data[ccIndex]
		cm++
		// fmt.Println(cm)
	}

	indexOfCup := findIndexOfCup(1, data, -1)
	fromCup := gatherTwoFromCup(indexOfCup, data)
	fmt.Println("from cup 1: ", fromCup)
}

// findDestinationCup returns the index of the destination cup.
func findDestinationCup(cc int, cups []int, startingPoint int) int {
	ccIndex := findIndexOfCup(cc, cups, startingPoint)
	// fmt.Printf("cc: %d, cc index: %d, cups: %+v\n", cc, ccIndex, cups)
	value := cups[ccIndex] - 1
	if value == 0 {
		value = 1000000 - 1
	}
	for {
		if i := findIndexOfCup(value, cups, -1); i != -1 {
			// fmt.Printf("found destination index: %d with value: %d\n", i, cups[i])
			return i
		}

		value--
		if value == 0 {
			value = 1000000 - 1
		}
		// fmt.Println("new value: ", value)
	}
}

func gatherTwoFromCup(cup int, cups []int) []int {
	two := make([]int, 0)

	for i := 0; i < 2; i++ {
		currentIndex := (i + cup + 1) % len(cups)
		two = append(two, cups[currentIndex])
	}

	return two
}

func gatherFromCup(cup int, cups []int) string {
	var result string

	for i := 0; i < len(cups); i++ {
		currentIndex := (i + cup) % len(cups)
		result += strconv.Itoa(cups[currentIndex])
	}

	return result
}

func findIndexOfCup(cup int, cups []int, startingPoint int) int {
	// fmt.Println("in index: ", cups)
	if startingPoint != -1 {
		start := startingPoint - 10
		if start < 0 {
			start = 0
		}
		end := startingPoint + 10
		if end >= len(cups) {
			end = len(cups)
		}
		for i := start; i < end; i++ {
			if cups[i] == cup {
				return i
			}
		}
	}
	for i := range cups {
		if cups[i] == cup {
			return i
		}
	}
	return -1
}
