package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	timeTable := strings.Split(lines[1], ",")

	keys := make([]int, 0)
	for _, t := range timeTable {
		if t == "x" {
			continue
		}

		id, _ := strconv.Atoi(t)
		keys = append(keys, id)
	}

	k := keys[0]
	d := keys[0]

	for dt, id := range keys {
		for {
			if (k+dt)%id == 0 {
				break
			}
			k += d
		}
		d = lcm(d, id)
	}

	fmt.Println(k)
}

func lcm(a, b int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}
	return a * b / gcd(a, b)
}
