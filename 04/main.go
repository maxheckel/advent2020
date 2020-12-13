package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
}

func ParseLine(line string) (Passport, error) {
	segments := strings.Split(line, " ")
	passport := Passport{}
	for _, segment := range segments {
		keyVal := strings.Split(segment, ":")
		switch keyVal[0] {
		case "byr":
			intVal, err := strconv.Atoi(keyVal[1])
			if err != nil {
				return Passport{}, err
			}
			passport.byr = intVal
			break
		case "iyr":
			intVal, err := strconv.Atoi(keyVal[1])
			if err != nil {
				return Passport{}, err
			}
			passport.iyr = intVal
			break
		case "eyr":
			intVal, err := strconv.Atoi(keyVal[1])
			if err != nil {
				return Passport{}, err
			}
			passport.eyr = intVal
			break
		case "hgt":
			passport.hgt = keyVal[1]
			break
		case "hcl":
			passport.hcl = keyVal[1]
			break
		case "ecl":
			passport.ecl = keyVal[1]
			break
		case "pid":
			passport.pid = keyVal[1]
			break
		}
	}
	return passport, nil
}

func (p *Passport) isValid() bool {
	if p.byr < 1920 || p.byr > 2002 {
		return false
	}
	if p.iyr < 2010 || p.iyr > 2020 {
		return false
	}

	if p.eyr < 2020 || p.eyr > 2030 {
		return false
	}
	if strings.Contains(p.hgt, "cm") {
		intVal, err := strconv.Atoi(strings.Replace(p.hgt, "cm", "", -1))
		if err != nil {
			return false
		}
		if intVal < 150 || intVal > 193 {
			return false
		}
	}

	if strings.Contains(p.hgt, "in") {
		intVal, err := strconv.Atoi(strings.Replace(p.hgt, "in", "", -1))
		if err != nil {
			return false
		}
		if intVal < 59 || intVal > 76 {
			return false
		}
	}

	if !strings.Contains(p.hgt, "in") && !strings.Contains(p.hgt, "cm") {
		return false
	}
	if len(p.hcl) > 7 {
		return false
	}
	var re = regexp.MustCompile(`(?m)#[0-9a-f]{6}`)
	if len(re.FindAllString(p.hcl, -1)) == 0 {
		return false
	}
	validEyeColors := map[string]bool{
		"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true,
	}
	if !validEyeColors[p.ecl] {
		return false
	}
	if len(p.pid) != 9 {
		return false
	}
	for _, c := range p.pid {
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("./4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	passportText := ""
	missingCredentials := 0
	badCredentials := 0

	for scanner.Scan() {
		currentRow := scanner.Text()
		passportText += currentRow + "\n"
	}
	passports := strings.Split(passportText, "\n\n")

	for i, passport := range passports {
		passport = strings.Replace(passport, "\n", " ", -1)
		passports[i] = passport
		if !hasRequiredFields(passport) {
			missingCredentials++
			continue
		}
		passportObj, _ := ParseLine(passport)
		if !passportObj.isValid() {
			badCredentials++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Part 1: %d", len(passports)-missingCredentials))
	fmt.Println(fmt.Sprintf("Part 2: %d", len(passports)-missingCredentials-badCredentials))

}

func hasRequiredFields(workingCredentials string) bool {
	requiredFields := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}
	for _, r := range requiredFields {
		if !strings.Contains(workingCredentials, r) {
			return false
		}
	}
	return true
}
