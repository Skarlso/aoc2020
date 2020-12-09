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
	if len(os.Args) < 2 {
		displayHelpThenExit()
	}
	skipCmd := flag.NewFlagSet("", flag.ExitOnError)
	skipCmd.IntVar(&year, "year", 2020, "the current year")
	skipCmd.IntVar(&day, "day", -1, "the current day")
	skipCmd.StringVar(&solution, "solution", "", "the solution")
	skipCmd.IntVar(&part, "part", -1, "part of the day")
	if err := skipCmd.Parse(os.Args[2:]); err != nil {
		log.Fatal(err)
	}
}

func displayHelpThenExit() {
	fmt.Println("Available commands:")
	fmt.Println("download -- download an input for a given year, day and part")
	fmt.Println("submit -- submit a solution for a given year, day and part")
	os.Exit(1)
}

func main() {
	cmd := os.Args[1]
	if cmd == "" {
		displayHelpThenExit()
	}
	switch cmd {
	case "download":
		fmt.Printf("Downloading year %d day %d\n", year, day)
		output, err := aoc.DownloadInput(year, day)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Done... %s\n", output)
	case "submit":
		fmt.Printf("Submitting year %d day %d part %d solution %s\n", year, day, part, solution)
		ok, err := aoc.SubmitSolution(year, day, part, solution)
		if err != nil {
			log.Fatal(err)
		}
		if ok {
			fmt.Println(":)")
		} else {
			fmt.Println(":(")
		}
	default:
		displayHelpThenExit()
	}
}
