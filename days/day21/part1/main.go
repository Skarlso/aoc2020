package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type food struct {
	// ingredients contains a list of allergies
	ingredients map[string]map[string]struct{}
	allergies   map[string]struct{}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// global list of ingredients to keep track of all of them separately.
	// we store what allergenes are contained in which ingredients here.
	// but we also store which food line has which ingredient in the food line itself
	// so we can compare them.
	ingredients := make(map[string]map[string]struct{})
	foods := make([]food, 0)
	for _, line := range bytes.Split(content, []byte("\n")) {
		// fmt.Println("line: ", string(line))
		split := strings.Split(string(line), "(contains")
		fmt.Println(split)
		ing := split[0]
		all := split[1]
		f := food{
			ingredients: make(map[string]map[string]struct{}),
			allergies:   make(map[string]struct{}),
		}
		for _, i := range strings.Split(ing, " ") {
			if i == "" {
				continue
			}
			f.ingredients[i] = make(map[string]struct{})
			// only do this if we didn't encounter this ingredient yet...
			if _, ok := ingredients[i]; !ok {
				ingredients[i] = make(map[string]struct{})
			}
		}
		for _, a := range strings.Split(all, ", ") {
			trimmed := strings.TrimRight(a, ")")
			f.allergies[trimmed] = struct{}{}
			// add allergene to ingredients? // I would have to know which ones are in this food item...
			for k := range f.ingredients {
				// to this ingredient in the global list, add this possible allergene as well
				// since we encountered it in this list
				ingredients[k][trimmed] = struct{}{}
			}
		}
		foods = append(foods, f)
	}
	fmt.Printf("FOODS: %+v\n", foods)
	// fmt.Printf("Ingredients: %+v\n", ingredients)
	for k, v := range ingredients {
		fmt.Println("k: ", k)
		fmt.Printf("v: %+v\n", v)
	}
	// loop until all ingredients have 0 or 1 allergies
	// that's flawed because they start out with 0.. this would quit immediately.

	// first loop, assign all allergenes to all ingredients in a first run.
	// then start trimming them down.
}
