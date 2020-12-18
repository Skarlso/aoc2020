package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, l := range split {
		l = strings.Trim(l, "\n")
		lines = append(lines, l)
	}

	sum := 0
	for _, l := range lines {
		l = solveParanthesis(l)
		l = solveAdditions(l)
		sum += solve(l)
	}

	fmt.Println(sum)
}

func solveParanthesis(expression string) string {
	re := regexp.MustCompile(`\([^\(\)]+\)`)

	for {
		ms := re.FindAllString(expression, -1)
		if len(ms) == 0 {
			break
		}

		for _, m := range ms {
			expression = solveAdditions(expression)
			replaceWith := solve(m[1 : len(m)-1])
			expression = strings.ReplaceAll(expression, m, fmt.Sprintf("%d", replaceWith))
		}
	}

	return expression
}

func solveAdditions(expression string) string {
	re := regexp.MustCompile(`\d+( \+ \d+)+`)

	for {
		ms := re.FindAllString(expression, -1)
		if len(ms) == 0 {
			break
		}

		for _, m := range ms {
			replaceWith := solve(m)
			expression = strings.ReplaceAll(expression, m, fmt.Sprintf("%d", replaceWith))
		}
	}

	return expression
}

// solve.
func solve(expression string) int {
	sum := 0
	currentOp := '+'
	// if we reached the end or a closing paranthesis, we return
	// the count we calculated thus far.
	var trackNumber []byte
	for _, c := range expression {
		switch c {
		case '+', '*':
			// add our track number
			if currentOp == '+' {
				// fmt.Println(string(trackNumber))
				n, _ := strconv.Atoi(string(trackNumber))
				sum += int(n)
				trackNumber = []byte("")
			} else if currentOp == '*' {
				// fmt.Println(string(trackNumber))
				n, _ := strconv.Atoi(string(trackNumber))
				sum *= int(n)
				trackNumber = []byte("")
			}
			currentOp = c
		}
		if unicode.IsNumber(c) {
			trackNumber = append(trackNumber, byte(c))
		}
	}

	// add the last number
	if currentOp == '+' {
		// fmt.Println(string(trackNumber))
		n, _ := strconv.Atoi(string(trackNumber))
		sum += int(n)
		trackNumber = []byte("")
	} else if currentOp == '*' {
		// fmt.Println(string(trackNumber))
		n, _ := strconv.Atoi(string(trackNumber))
		sum *= int(n)
		trackNumber = []byte("")
	}

	return sum
}
