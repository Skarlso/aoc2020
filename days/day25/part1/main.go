package main

import "fmt"

var (
	// cardPK = 5764801
	// doorPK = 17807724

	subjectNumber = 7
	cardPK        = 14082811
	doorPK        = 5249543
)

func main() {

	cardLoopSize := 0
	doorLoopSize := 0
	value := 1
	for i := 1; ; i++ {
		value *= subjectNumber
		value %= 20201227

		if value == cardPK {
			fmt.Println("Found cardLoopSize: ", i)
			cardLoopSize = i
		}
		if value == doorPK {
			fmt.Println("Found doorLoopSize: ", i)
			doorLoopSize = i
		}

		if doorLoopSize > 0 && cardLoopSize > 0 {
			fmt.Println("Found both.")
			break
		}
	}

	subjectNumber = doorPK
	value = 1
	for i := 0; i < cardLoopSize; i++ {
		value *= subjectNumber
		value %= 20201227
	}
	fmt.Println("encryption key: ", value)
}
