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
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i]+n[j] == 2020 {
				fmt.Println(n[i] * n[j])
			}
		}
	}
}
