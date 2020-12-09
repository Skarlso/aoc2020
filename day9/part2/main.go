package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var target = 41682220

func main() {
	if len(os.Args) < 1 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, l := range split {
		lines = append(lines, l)
	}

	// convert to int
	numbers := make([]int, 0, len(lines))
	for _, l := range lines {
		n, _ := strconv.Atoi(l)
		numbers = append(numbers, n)
	}

	for i := 0; i < len(numbers)-1; i++ {
		sum := numbers[i]
		slidingNumber := []int{numbers[i]}
		for j := i + 1; j < len(numbers); j++ {
			slidingNumber = append(slidingNumber, numbers[j])
			sum += numbers[j]
			if sum == target {
				sort.Ints(slidingNumber)
				min := slidingNumber[0]
				max := slidingNumber[len(slidingNumber)-1]
				fmt.Println("min, max", min, max)
				fmt.Println("sum: ", min+max)
				return
			}
		}
	}
}
