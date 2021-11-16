package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: <binary> <test.txt>")
		os.Exit(1)
	}
	f := os.Args[1]
	content, err := os.ReadFile(f)
	if err != nil {
		fmt.Println("Failed to read file: ", err)
		os.Exit(1)
	}

	paths := make([][]string, 0)
	for _, line := range bytes.Split(content, []byte("\n")) {
		path := parsePath(string(line))
		paths = append(paths, path)
	}
}

func parsePath(s string) []string {
	var result []string
	split := strings.Split(s, "")
	for i := 0; i < len(split); i++ {
		r := string(split[i])
		if r == "n" || r == "s" {
			r += split[i+1]
			i++
		}
		result = append(result, r)
	}
	return result
}
