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
	// data  = []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	// moves = 100
	// cc    = 3 // index of the current cup
	data  = []int{1, 5, 8, 9, 3, 7, 4, 6, 2}
	moves = 100
	cc    = 1
)

func main() {
	// fmt.Println(data)
	cm := 0 // current moves
	for cm < moves {
		// take three
		three := make([]int, 0)
		ccIndex := findIndexOfCup(cc, data)
		for i := 0; i < 3; i++ {
			ci := (i + ccIndex + 1) % len(data)
			three = append(three, data[ci])
		}
		for _, c := range three {
			// this would be a bit better if this would be a tree, but I don't care for now.
			index := findIndexOfCup(c, data)
			data = append(data[:index], data[index+1:]...)
		}
		dc := findDestinationCup(cc, data)
		data = append(data[:dc+1], append(three, data[dc+1:]...)...)
		// update the index of CC after it has been moved around by the insert again.
		ccIndex = findIndexOfCup(cc, data)
		ccIndex = (ccIndex + 1) % len(data)
		cc = data[ccIndex]
		cm++
	}

	indexOfCup := findIndexOfCup(1, data)
	fromCup := gatherFromCup(indexOfCup, data)
	fmt.Println("from cup 1: ", fromCup[1:])
}

// findDestinationCup returns the index of the destination cup.
func findDestinationCup(cc int, cups []int) int {
	ccIndex := findIndexOfCup(cc, cups)
	// fmt.Printf("cc: %d, cc index: %d, cups: %+v\n", cc, ccIndex, cups)
	value := cups[ccIndex] - 1
	if value == 0 {
		value = 9
	}
	// fmt.Println("initial value: ", value)
	for {
		if i := findIndexOfCup(value, cups); i != -1 {
			// fmt.Printf("found destination index: %d with value: %d\n", i, cups[i])
			return i
		}

		value--
		if value == 0 {
			value = 9
		}
		// fmt.Println("new value: ", value)
	}
}

func gatherFromCup(cup int, cups []int) string {
	var result string

	for i := 0; i < len(cups); i++ {
		currentIndex := (i + cup) % len(cups)
		result += strconv.Itoa(cups[currentIndex])
	}

	return result
}

func findIndexOfCup(cup int, cups []int) int {
	for i := range cups {
		if cups[i] == cup {
			return i
		}
	}
	return -1
}
