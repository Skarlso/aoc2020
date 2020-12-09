package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var test = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`

var neededFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
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
		fields := make(map[string]struct{})
		for _, f := range pass {
			split := strings.Split(f, " ")
			for _, s := range split {
				split2 := strings.Split(s, ":")
				fields[split2[0]] = struct{}{}
			}
		}
		valid := true
		for _, f := range neededFields {
			if _, ok := fields[f]; !ok {
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
