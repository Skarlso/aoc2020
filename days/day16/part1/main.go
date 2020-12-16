package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// var rules = `class: 1-3 or 5-7
// row: 6-11 or 33-44
// seat: 13-40 or 45-50`
//
// var yourTicket = `your ticket:
// 7,1,14`

var rules = `departure location: 47-874 or 885-960
departure station: 25-616 or 622-964
departure platform: 42-807 or 825-966
departure track: 36-560 or 583-965
departure date: 37-264 or 289-968
departure time: 27-325 or 346-954
arrival location: 37-384 or 391-950
arrival station: 35-233 or 244-963
arrival platform: 26-652 or 675-949
arrival track: 41-689 or 710-954
class: 27-75 or 81-952
duration: 45-784 or 807-967
price: 40-350 or 374-970
route: 30-892 or 904-968
row: 47-144 or 151-957
seat: 28-750 or 773-973
train: 30-456 or 475-950
type: 34-642 or 648-968
wagon: 42-486 or 498-970
zone: 37-152 or 167-973`

var myTicket = `83,137,101,73,67,61,103,131,151,127,113,107,109,89,71,139,167,97,59,53`

type validator func(in int) bool

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, l := range split {
		lines = append(lines, l)
	}

	validators := make([]validator, 0)

	// Parse the rules
	for _, l := range strings.Split(rules, "\n") {
		fmt.Println("line: ", l)
		var (
			lowerMin, lowerMax int
			upperMin, upperMax int
		)
		split := strings.Split(l, ": ")
		fields := split[1]

		fmt.Sscanf(fields, "%d-%d or %d-%d", &lowerMin, &lowerMax, &upperMin, &upperMax)

		f := func(in int) bool {
			return (in >= lowerMin && in <= lowerMax) || (in >= upperMin && in <= upperMax)
		}

		validators = append(validators, f)
	}

	// Parse my ticket... But later.

	// Now go through all the lines, parse them, and run all validators.
	// If one of them fails, the ticket is invalid.
	sum := 0
	for _, l := range lines {
		fields := strings.Split(l, ",")
		for _, f := range fields {
			i, _ := strconv.Atoi(f)
			valid := false
			for _, v := range validators {
				if v(i) {
					valid = true
					break
				}
			}

			if !valid {
				sum += i
			}
		}
	}

	fmt.Println(sum)

}
