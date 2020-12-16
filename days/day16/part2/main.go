package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// var rules = `class: 0-1 or 4-19
// row: 0-5 or 8-19
// seat: 0-13 or 16-19`

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

// var myTicket = "11,12,13"

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
		// fmt.Println("line: ", l)
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

	// Parse my ticket... But later.

	// Now go through all the lines, parse them, and run all validators.
	// If one of them fails, the ticket is invalid.
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

	// I have my valid tickets

	// there are as many columns as there are validators of course. For each column there
	// is a validator.
	// fmt.Println("Valid tickets: ", validTickets)

	// Save the field once we find the valid column with name and the column it was valid for.
	fields := make([]field, 0)
	// We know that eventually we will find a valid column
	// because we threw out all the invalid tickets.
	column := 0 // column tracker. which column we are testing ATM for all tickets.
	for {
		allFieldsOfTheValidTicket := make([]int, 0)
		for _, ticket := range validTickets {
			split := strings.Split(ticket, ",")
			i, _ := strconv.Atoi(split[column])
			allFieldsOfTheValidTicket = append(allFieldsOfTheValidTicket, i)
		}

		// now run all validators on these numbers the one we find that validates them all
		// is the validator for that column.
		for i := 0; i < len(validators); i++ {
			v := validators[i]
			valid := true
			for _, n := range allFieldsOfTheValidTicket {
				if !v.f(n) {
					valid = false
					break
				}
			}

			if valid {
				fields = append(fields, field{name: v.name, column: column})
				validators = append(validators[:i], validators[i+1:]...)
				break
			}
		}

		// fmt.Println(fields)
		column++
		// if column == 20 {
		// 	break
		// }
		if len(validators) == 0 {
			break
		}
	}

	fmt.Println(fields)

	multi := 1
	for _, field := range fields {
		myTicketFields := strings.Split(myTicket, ",")
		if strings.HasPrefix(field.name, "departure") {
			i, _ := strconv.Atoi(myTicketFields[field.column])
			fmt.Println(i)
			multi *= i
		}
	}

	fmt.Println(multi)
}
