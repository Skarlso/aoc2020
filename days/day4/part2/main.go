package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type pass struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *pass) isValid() bool {
	// TODO Add cid when it matters.
	if p.byr == "" || p.iyr == "" || p.eyr == "" || p.hgt == "" || p.hcl == "" || p.ecl == "" || p.pid == "" {
		return false
	}
	return p.validateByr() && p.validateIyr() && p.validateEyr() && p.validateHgt() && p.validateHcl() && p.validateEcl() && p.validatePid() && p.validateCid()
}

func (p *pass) validateByr() bool {
	if len(p.byr) != 4 {
		fmt.Println("Not 4 byr")
		return false
	}
	i := convertToInt("byr", p.byr, 0)
	if !inBetween(i, 1920, 2002) {
		return false
	}
	return true
}

func (p *pass) validateIyr() bool {
	if len(p.iyr) != 4 {
		fmt.Println("Not 4 iyr")
		return false
	}
	i := convertToInt("iyr", p.iyr, 0)
	if !inBetween(i, 2010, 2020) {
		return false
	}
	return true
}

func (p *pass) validateEyr() bool {
	if len(p.eyr) != 4 {
		return false
	}
	i := convertToInt("eyr", p.eyr, 0)
	if !inBetween(i, 2020, 2030) {
		return false
	}
	return true
}

func (p *pass) validateHgt() bool {
	if strings.HasSuffix(p.hgt, "cm") {
		v := strings.TrimSuffix(p.hgt, "cm")
		i := convertToInt("hgt cm", v, 0)
		if i < 150 || i > 193 {
			fmt.Println("Not between hgt cm i: ", i)
			return false
		}
	} else if strings.HasSuffix(p.hgt, "in") {
		v := strings.TrimSuffix(p.hgt, "in")
		i := convertToInt("hgt in", v, 0)
		if i < 59 || i > 76 {
			fmt.Println("Not between hgt in i: ", i)
			return false
		}
	} else {
		return false
	}
	return true
}

func (p *pass) validateHcl() bool {
	if p.hcl[0] != '#' {
		return false
	}
	// trim the #
	r := regexp.MustCompile(`^[a-f|0-9]+$`)
	v := string(p.hcl[1:])

	if !r.MatchString(v) {
		fmt.Println("No match hcl: ", v)
		return false
	}
	return true
}

func (p *pass) validateEcl() bool {
	allowed := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	if !contains(allowed, p.ecl) {
		return false
	}
	return true
}

func (p *pass) validatePid() bool {
	if len(p.pid) != 9 {
		fmt.Println("Pid length not 9: ", p.pid)
		return false
	}
	return true
}

func (p *pass) validateCid() bool {
	return true
}

func newPass(lines []string) *pass {
	p := &pass{}
	for _, line := range lines {
		lsplit := strings.Split(line, " ")
		for _, s := range lsplit {
			split := strings.Split(s, ":")
			switch split[0] {
			case "byr":
				p.byr = split[1]
			case "iyr":
				p.iyr = split[1]
			case "eyr":
				p.eyr = split[1]
			case "hgt":
				p.hgt = split[1]
			case "hcl":
				p.hcl = split[1]
			case "ecl":
				p.ecl = split[1]
			case "pid":
				p.pid = split[1]
			case "cid":
				p.cid = split[1]
			default:
				log.Fatal("Unknown field: ", split[0])
			}
		}
	}
	return p
}

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
		p := newPass(pass)

		if p.isValid() {
			fmt.Println("Found valid: ", pass)
			valids++
		}
	}

	fmt.Println("Valids: ", valids)
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
