package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Password struct {
	MinAppearances int
	MaxAppearances int
	Letter string
	Password string
}

func (p *Password) isValid() bool{
	appearances := strings.Count(p.Password, p.Letter)
	return appearances >= p.MinAppearances && appearances <= p.MaxAppearances
}

func (p *Password) isValidPart2() bool{
	pos1Matches := string(p.Password[p.MinAppearances-1]) == p.Letter
	pos2Matches := string(p.Password[p.MaxAppearances-1]) == p.Letter
	return !(pos1Matches && pos2Matches) && (pos1Matches || pos2Matches)
}

func (p *Password) build(input string){
	halves := strings.Split(input, ": ")
	p.Password = halves[1]
	rules := strings.Split(halves[0], " ")
	p.Letter = rules[1]
	minMax := strings.Split(rules[0], "-")
	p.MinAppearances, _ = strconv.Atoi(minMax[0])
	p.MaxAppearances, _ = strconv.Atoi(minMax[1])
}

func main() {
	file, err := os.Open("./2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	part1ValidCount := 0
	part2ValidCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newPassword := Password{}
		newPassword.build(scanner.Text())
		if newPassword.isValid() {
			part1ValidCount++
		}
		if newPassword.isValidPart2() {
			part2ValidCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1ValidCount)
	fmt.Println(part2ValidCount)
}
