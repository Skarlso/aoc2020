package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Skarlso/aoc2020/pkg/asm"
)

func main() {
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, l := range split {
		lines = append(lines, l)
	}

	runner, _ := asm.NewASMRunner(lines)
	if _, err := runner.Run(); err != nil {
		log.Fatal(err)
	}
}
