package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")

	split := strings.Split(string(content), "\n")
	n := make([]int, 0)
	for _, l := range split {
		i, _ := strconv.Atoi(strings.Trim(l, "\n"))
		n = append(n, i)
	}
	for i := 0; i < len(n)-2; i++ {
		for j := i + 1; j < len(n)-1; j++ {
			for g := j + 1; g < len(n); g++ {
				if n[i]+n[j]+n[g] == 2020 {
					fmt.Println(n[i] * n[j] * n[g])
					break
				}
			}
		}
	}
}
