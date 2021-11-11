package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
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
	// allergenes := make(map[string]map[string]struct{})
	foods := make([]food, 0)
	for _, line := range bytes.Split(content, []byte("\n")) {
		// fmt.Println("line: ", string(line))
		split := strings.Split(string(line), "(contains")
		// fmt.Println(split)
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
			trimmed = strings.TrimSpace(trimmed)
			if a == "" {
				continue
			}
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

	// initial cleanup loop

	for _, food := range foods {
		for a := range food.allergies {
			for i, alls := range ingredients {
				// everything that has this allergenes but is NOT in the list of ingredients in the food, remove that allergenes
				if _, ok := alls[a]; ok {
					if _, ok := food.ingredients[i]; !ok {
						delete(ingredients[i], a)
					}
				}
			}
		}
	}

	// NOW: Go over all the ingredients which have ONE and remove that ONE allergene from the OTHER ingredients allergies
	// because it has been taken.

	// remove stragglers
	for {
		for k, v := range ingredients {
			if len(v) == 1 {
				for a := range v {
					for i, all := range ingredients {
						if k != i {
							delete(all, a)
						}
					}
				}
			}
		}

		done := true
		for _, v := range ingredients {
			if len(v) > 1 {
				done = false
				break
			}
		}
		if done {
			break
		}
	}

	// I think we are possibly done...
	// Get all the one which have no allergies
	all := make([]string, 0)
	for _, a := range ingredients {
		if len(a) == 1 {
			// fmt.Println(a)
			for v := range a {
				all = append(all, v)
			}
		}
	}

	// fmt.Println(all)
	fmt.Println(ingredients)

	sort.Strings(all)

	ones := make([]string, 0)
	for _, a := range all {
		for k, v := range ingredients {
			if _, ok := v[a]; ok {
				ones = append(ones, k)
			}
		}
	}

	fmt.Println(strings.Join(ones, ","))
}
