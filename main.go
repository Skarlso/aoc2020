package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Skarlso/aoc2020/aoc"
)

var (
	year     int
	day      int
	part     int
	solution string
)

func init() {
	flag.IntVar(&year, "year", -1, "the current year")
	flag.IntVar(&day, "day", -1, "the current day")
	flag.IntVar(&part, "part", -1, "part of the day")
	flag.StringVar(&solution, "solution", "", "the solution")
}

func main() {
	cmd := os.Args[1]
	if cmd == "" {
		fmt.Println("Available commands:")
		fmt.Println("download -- download an input for a given year, day and part")
		fmt.Println("submit -- submit a solution for a given year, day and part")
		os.Exit(1)
	}
	switch cmd {
	case "download":
		output, err := aoc.DownloadInput(year, day, part)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Done... %s\n", output)
	case "submit":
		ok, err := aoc.SubmitSolution(year, day, part, solution)
		if err != nil {
			log.Fatal(err)
		}
		if ok {
			fmt.Println(":)")
		} else {
			fmt.Println(":(")
		}
	}
}
