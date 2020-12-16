package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

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

type sValidator struct {
	f    validator
	name string
}

type field struct {
	name   string
	column int
}

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

	validators := make([]sValidator, 0)

	// Parse the rules
	for _, l := range strings.Split(rules, "\n") {
		var (
			lowerMin, lowerMax int
			upperMin, upperMax int
		)
		split := strings.Split(l, ": ")
		fields := split[1]
		name := split[0]

		fmt.Sscanf(fields, "%d-%d or %d-%d", &lowerMin, &lowerMax, &upperMin, &upperMax)

		f := func(in int) bool {
			return (in >= lowerMin && in <= lowerMax) || (in >= upperMin && in <= upperMax)
		}

		validators = append(validators, sValidator{
			f:    f,
			name: name,
		})
	}

	validTickets := make([]string, 0)
	for _, l := range lines {
		validTicket := true
		fields := strings.Split(l, ",")
		for _, f := range fields {
			i, _ := strconv.Atoi(f)
			valid := false
			for _, v := range validators {
				if v.f(i) {
					valid = true
					break
				}
			}

			if !valid {
				validTicket = false
				break
			}
		}

		if validTicket {
			validTickets = append(validTickets, l)
		}
	}

	// Add my own ticket to the valid tickets
	validTickets = append(validTickets, myTicket)

	// Save the field once we find the valid column with name and the column it was valid for.
	fields := make([]field, 0)
	assignedColumns := make(map[int]struct{})
	maxColumns := len(validators)
	// we go until all validators have been found
	for len(validators) != 0 {
		for i := 0; i < len(validators); i++ {
			v := validators[i]
			// and if it's ONLY valid for a single column then we say it's that column
			validForCount := 0
			index := 0

			for c := 0; c < maxColumns; c++ {
				if _, ok := assignedColumns[c]; ok {
					// skip already found columns
					continue
				}

				valid := true
				for _, t := range validTickets {
					values := strings.Split(t, ",")
					value := values[c]
					n, _ := strconv.Atoi(value)
					if !v.f(n) {
						valid = false
						break
					}
				}

				if valid {
					validForCount++
					index = c
				}
			}

			if validForCount == 1 {
				assignedColumns[index] = struct{}{}
				fields = append(fields, field{name: v.name, column: index})
				// delete and begin again
				validators = append(validators[:i], validators[i+1:]...)
				break
			}
		}
	}

	multi := 1
	for _, field := range fields {
		myTicketFields := strings.Split(myTicket, ",")
		if strings.HasPrefix(field.name, "departure") {
			i, _ := strconv.Atoi(myTicketFields[field.column])
			multi *= i
		}
	}

	fmt.Println(multi)
}
