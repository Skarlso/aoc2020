package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Skarlso/aoc2020/pkg/perf"
)

type sum struct {
	a, b int
}

func main() {
	defer perf.Duration(perf.Track("Part1"))

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

	sums := make(map[sum]int)

	// prefill
	// slidingNumbers is the tracker for the sliding numbers.
	var slidingNumbers []int
	slidingNumbers, numbers = numbers[:25], numbers[25:]

	for i := 0; i < len(slidingNumbers)-1; i++ {
		for j := i + 1; j < len(slidingNumbers); j++ {
			sums[sum{a: slidingNumbers[i], b: slidingNumbers[j]}] = slidingNumbers[i] + slidingNumbers[j]
		}
	}

	// fmt.Println(sums)
	for _, n := range numbers {
		isSum := false
		for _, v := range sums {
			if n == v {
				isSum = true
				break
			}
		}

		if !isSum {
			fmt.Println("No sum found for: ", n)
			return
		}

		// update our sums and remove the first one
		var first int
		first, slidingNumbers = slidingNumbers[0], slidingNumbers[1:]

		// add the new number to the end of the sliding numbers
		slidingNumbers = append(slidingNumbers, n)

		// Now go through the sums and remove every combination which has `first` in it.
		for k := range sums {
			if k.a == first || k.b == first {
				delete(sums, k)
			}
		}

		// add the combination of the new number to the sums
		for _, sn := range slidingNumbers {
			if sn != n {
				sums[sum{a: sn, b: n}] = sn + n
			}
		}
	}

	fmt.Println("Every number was a sum of the previous numbers.")
}
