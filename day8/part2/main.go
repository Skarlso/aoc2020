package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	nops := make([]int, 0)
	jmps := make([]int, 0)
	for i, l := range split {
		if strings.HasPrefix(l, "nop") {
			nops = append(nops, i)
		} else if strings.HasPrefix(l, "jmp") {
			jmps = append(jmps, i)
		}
		lines = append(lines, l)
	}

	// first, try every jmp
	for _, i := range jmps {
		acc := 0
		offset := 0
		seen := make(map[int]struct{})
		newLines := make([]string, len(lines))
		copy(newLines, lines)

		newLines[i] = strings.ReplaceAll(newLines[i], "jmp", "nop")

		for {
			if _, ok := seen[offset]; ok {
				break
			}
			instruction := strings.Split(newLines[offset], " ")
			// fmt.Println(lines[offset])
			seen[offset] = struct{}{}
			op := instruction[0]
			inst, _ := strconv.Atoi(instruction[1])
			// fmt.Println(op, inst, acc)
			switch op {
			case "acc":
				acc += inst
				offset++
			case "jmp":
				offset += inst
			case "nop":
				// do nothing
				offset++
			}

			if offset >= len(newLines) {
				fmt.Println("We have a winner: ", acc)
				return
			}
		}
	}

	// second, try every nop
	// first, try every jmp
	for _, i := range nops {
		acc := 0
		offset := 0
		seen := make(map[int]struct{})
		newLines := make([]string, len(lines))
		copy(newLines, lines)

		newLines[i] = strings.ReplaceAll(newLines[i], "nop", "jmp")

		for {
			if _, ok := seen[offset]; ok {
				break
			}
			instruction := strings.Split(newLines[offset], " ")
			// fmt.Println(lines[offset])
			seen[offset] = struct{}{}
			op := instruction[0]
			inst, _ := strconv.Atoi(instruction[1])
			// fmt.Println(op, inst, acc)
			switch op {
			case "acc":
				acc += inst
				offset++
			case "jmp":
				offset += inst
			case "nop":
				// do nothing
				offset++
			}

			if offset >= len(newLines) {
				fmt.Println("We have a winner: ", acc)
				return
			}
		}
	}

}
