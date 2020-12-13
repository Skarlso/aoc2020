package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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
		lines = append(lines, l)
	}

	departure, _ := strconv.Atoi(lines[0])
	timeTable := strings.Split(lines[1], ",")

	keys := make([]int, 0)
	for _, t := range timeTable {
		if t == "x" {
			continue
		}

		id, _ := strconv.Atoi(t)
		keys = append(keys, id)
	}

	min := math.MaxInt64
	res := 0
	// Time skip
	for _, k := range keys {
		mins := k - (departure % k)
		if mins < min {
			min = mins
			res = mins * k
		}
	}
	fmt.Println(res)
}
