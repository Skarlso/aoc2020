package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Skarlso/aoc2020/asm"
)

func main() {
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, l := range split {
		lines = append(lines, l)
	}

	// acc := 0
	// offset := 0

	// seen := make(map[int]struct{})
	// for {
	// 	if _, ok := seen[offset]; ok {
	// 		break
	// 	}
	// 	instruction := strings.Split(lines[offset], " ")
	// 	// fmt.Println(lines[offset])
	// 	seen[offset] = struct{}{}

	// 	op := instruction[0]
	// 	inst, _ := strconv.Atoi(instruction[1])
	// 	// fmt.Println(op, inst, acc)
	// 	switch op {
	// 	case "acc":
	// 		acc += inst
	// 		offset++
	// 	case "jmp":
	// 		offset += inst
	// 	case "nop":
	// 		// do nothing
	// 		offset++
	// 	}
	// }

	runner, _ := asm.NewASMRunner(lines)
	if _, err := runner.Run(); err != nil {
		log.Fatal(err)
	}
}
