package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var test = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
`

var neededFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	// content := test
	split := strings.Split(string(content), "\n")

	lines := make([]string, 0)
	for _, l := range split {
		lines = append(lines, l)
	}

	passports := make([][]string, 0)
	p := make([]string, 0)

	for _, l := range lines {
		if l == "" {
			passports = append(passports, p)
			p = make([]string, 0)
			continue
		}

		p = append(p, l)
	}
	// add the last part
	passports = append(passports, p)

	valids := 0
	for _, pass := range passports {
		fields := make(map[string]string)
		for _, f := range pass {
			split := strings.Split(f, " ")
			for _, s := range split {
				split2 := strings.Split(s, ":")
				fields[split2[0]] = split2[1]
			}
			// fmt.Println(fields)
		}
		valid := true
		for _, f := range neededFields {
			if _, ok := fields[f]; !ok {
				valid = false
				break
			}
		}

		for _, f := range neededFields {
			if v, ok := fields[f]; ok && !validateField(f, v) {
				fmt.Println("Was found not valid: ", pass)
				valid = false
				break
			}
		}

		if valid {
			fmt.Println("Found valid: ", pass)
			valids++
		}
	}

	fmt.Println("Valids: ", valids)
}

func validateField(f, v string) bool {
	// fmt.Println("received fields to validate: ", f, v)
	if len(v) == 0 {
		return false
	}

	switch f {
	case "byr":
		if len(v) != 4 {
			fmt.Println("Not 4 byr")
			return false
		}
		i := convertToInt("byr", v, 0)
		if !inBetween(i, 1920, 2002) {
			return false
		}
	case "iyr":
		if len(v) != 4 {
			fmt.Println("Not 4 iyr")
			return false
		}
		i := convertToInt("iyr", v, 0)
		if !inBetween(i, 2010, 2020) {
			return false
		}
	case "eyr":
		if len(v) != 4 {
			return false
		}
		i := convertToInt("eyr", v, 0)
		if !inBetween(i, 2020, 2030) {
			return false
		}
	case "hgt":
		if strings.HasSuffix(v, "cm") {
			v = strings.TrimSuffix(v, "cm")
			i := convertToInt("hgt cm", v, 0)
			if i < 150 || i > 193 {
				fmt.Println("Not between hgt cm i: ", i)
				return false
			}
		} else if strings.HasSuffix(v, "in") {
			v = strings.TrimSuffix(v, "in")
			i := convertToInt("hgt in", v, 0)
			if i < 59 || i > 76 {
				fmt.Println("Not between hgt in i: ", i)
				return false
			}
		} else {
			return false
		}
	case "hcl":
		if v[0] != '#' {
			return false
		}
		// trim the #
		r := regexp.MustCompile(`^[a-f|0-9]+$`)
		v = string(v[1:])

		if !r.MatchString(v) {
			fmt.Println("No match hcl: ", v)
			return false
		}
	case "ecl":
		allowed := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		if !contains(allowed, v) {
			return false
		}
	case "pid":
		if len(v) != 9 {
			fmt.Println("Pid length not 9: ", v)
			return false
		}
	}

	return true
}

// convertToInt converts a string to int and prints the error in red
// then returns a default value.
func convertToInt(name, v string, def int) int {
	colorRed := "\033[31m"
	colorReset := "\033[0m"

	i, err := strconv.Atoi(v)
	if err != nil {
		fmt.Printf("err converting to int in %s. v: %s, err: %s%v%s\n", name, v, string(colorRed), err, string(colorReset))
		return def
	}
	return i
}

func contains(arr []string, v string) bool {
	for _, a := range arr {
		if v == a {
			return true
		}
	}

	return false
}

func inBetween(v, min, max int) bool {
	return v >= min && v <= max
}
