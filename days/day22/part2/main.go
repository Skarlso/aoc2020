package main

import (
	"fmt"
	"strconv"
)

const (
	player1 = iota
	player2
)

// var (
// dec1 = []int{9, 2, 6, 3, 1}
// dec2 = []int{5, 8, 4, 7, 10}
// dec1 = []int{24, 22, 26, 6, 14, 19, 27, 17, 39, 34, 40, 41, 23, 30, 36, 11, 28, 3, 10, 21, 9, 50, 32, 25, 8}
// dec2 = []int{48, 49, 47, 15, 42, 44, 5, 4, 13, 7, 20, 43, 12, 37, 29, 18, 45, 16, 1, 46, 38, 35, 2, 33, 31}
// )

func main() {
	// scoring must happen after the game has determined who won
	// dec1 := []int{9, 2, 6, 3, 1}
	// dec2 := []int{5, 8, 4, 7, 10}
	dec1 := []int{24, 22, 26, 6, 14, 19, 27, 17, 39, 34, 40, 41, 23, 30, 36, 11, 28, 3, 10, 21, 9, 50, 32, 25, 8}
	dec2 := []int{48, 49, 47, 15, 42, 44, 5, 4, 13, 7, 20, 43, 12, 37, 29, 18, 45, 16, 1, 46, 38, 35, 2, 33, 31}
	won := 0

	var (
		cardCacheDeck1 = make(map[string]struct{})
		cardCacheDeck2 = make(map[string]struct{})
	)
	for {
		if len(dec1) == 0 {
			won = player2
			break
		}
		if len(dec2) == 0 {
			won = player1
			break
		}

		var (
			card1, card2 int
		)

		card1, dec1 = dec1[0], dec1[1:]
		card2, dec2 = dec2[0], dec2[1:]

		if _, ok := cardCacheDeck1[hashDeck(dec1)]; ok {
			won = player1
			break
		}

		if _, ok := cardCacheDeck2[hashDeck(dec2)]; ok {
			won = player1
			break
		}

		if len(dec1) >= card1 && len(dec2) >= card2 {
			copyDec1 := make([]int, 0)
			copyDec1 = append(copyDec1, dec1[:card1]...)
			copyDec2 := make([]int, 0)
			copyDec2 = append(copyDec2, dec2[:card2]...)
			if winner := game(copyDec1, copyDec2); winner == player1 {
				dec1 = append(dec1, card1, card2)
			} else {
				dec2 = append(dec2, card2, card1)
			}
		} else {
			if card1 > card2 {
				dec1 = append(dec1, card1, card2)
			} else if card2 > card1 {
				dec2 = append(dec2, card2, card1)
			}
		}

		cardCacheDeck1[hashDeck(dec1)] = struct{}{}
		cardCacheDeck2[hashDeck(dec2)] = struct{}{}
	}

	var score int
	if won == player1 {
		score = scoreDeck(dec1)
	} else {
		score = scoreDeck(dec2)
	}
	fmt.Println("player: ", score)
}

// game runs a game of recursive combat.
func game(deck1, deck2 []int) int {
	// fmt.Println(deck1, deck2)
	var (
		cardCacheDeck1 = make(map[string]struct{})
		cardCacheDeck2 = make(map[string]struct{})
	)
	for {
		if len(deck1) == 0 {
			return player2
		}
		if len(deck2) == 0 {
			return player1
		}

		var (
			card1, card2 int
		)

		card1, deck1 = deck1[0], deck1[1:]
		card2, deck2 = deck2[0], deck2[1:]

		if _, ok := cardCacheDeck1[hashDeck(deck1)]; ok {
			return player1
		}

		if _, ok := cardCacheDeck2[hashDeck(deck2)]; ok {
			return player1
		}

		if len(deck1) >= card1 && len(deck2) >= card2 {
			copyDec1 := make([]int, 0)
			copyDec1 = append(copyDec1, deck1[:card1]...)
			copyDec2 := make([]int, 0)
			copyDec2 = append(copyDec2, deck2[:card2]...)
			if winner := game(copyDec1, copyDec2); winner == player1 {
				deck1 = append(deck1, card1, card2)
			} else {
				deck2 = append(deck2, card2, card1)
			}
		} else {
			if card1 > card2 {
				deck1 = append(deck1, card1, card2)
			} else if card2 > card1 {
				deck2 = append(deck2, card2, card1)
			}
		}

		cardCacheDeck1[hashDeck(deck1)] = struct{}{}
		cardCacheDeck2[hashDeck(deck2)] = struct{}{}
	}
}

func scoreDeck(deck []int) int {
	var score int
	for i, j := len(deck)-1, 1; i >= 0; i, j = i-1, j+1 {
		score += deck[i] * j
	}
	return score
}

func hashDeck(deck []int) string {
	var result string
	for _, i := range deck {
		result += strconv.Itoa(i)
	}
	return result
}
