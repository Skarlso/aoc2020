package main

import "fmt"

var (
	// dec1 = []int{9, 2, 6, 3, 1}
	// dec2 = []int{5, 8, 4, 7, 10}
	dec1 = []int{24, 22, 26, 6, 14, 19, 27, 17, 39, 34, 40, 41, 23, 30, 36, 11, 28, 3, 10, 21, 9, 50, 32, 25, 8}
	dec2 = []int{48, 49, 47, 15, 42, 44, 5, 4, 13, 7, 20, 43, 12, 37, 29, 18, 45, 16, 1, 46, 38, 35, 2, 33, 31}
)

func main() {
	var (
		rounds int
		winner string
		score  int
	)
	for {
		if len(dec1) == 0 {
			winner = "player 2"
			for i, j := len(dec2)-1, 1; i >= 0; i, j = i-1, j+1 {
				score += dec2[i] * j
			}
			break
		}
		if len(dec2) == 0 {
			winner = "player 1"
			for i, j := len(dec1)-1, 1; i >= 0; i, j = i-1, j+1 {
				score += dec1[i] * j
			}
			break
		}

		var (
			card1, card2 int
		)

		card1, dec1 = dec1[0], dec1[1:]
		card2, dec2 = dec2[0], dec2[1:]
		if card1 > card2 {
			dec1 = append(dec1, card1, card2)
		} else if card2 > card1 {
			dec2 = append(dec2, card2, card1)
		}
		rounds++
	}

	fmt.Println(winner, rounds, score)
}
