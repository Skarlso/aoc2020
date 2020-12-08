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
		newLines := make([]string, len(lines))
		copy(newLines, lines)

		newLines[i] = strings.ReplaceAll(newLines[i], "jmp", "nop")

		runner, _ := asm.NewASMRunner(newLines)
		if ok, err := runner.Run(); err != nil {
			log.Fatal(err)
		} else if ok {
			return
		}
	}

	// second, try every nop
	for _, i := range nops {
		newLines := make([]string, len(lines))
		copy(newLines, lines)
		newLines[i] = strings.ReplaceAll(newLines[i], "nop", "jmp")

		runner, _ := asm.NewASMRunner(newLines)
		if ok, err := runner.Run(); err != nil {
			log.Fatal(err)
		} else if ok {
			return
		}
	}

}
