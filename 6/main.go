package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	answersText := ""

	for scanner.Scan() {
		currentRow := scanner.Text()
		answersText += currentRow + "\n"
	}
	answers := strings.Split(answersText, "\n\n")
	totalUnique := 0
	globallyUnique := 0
	for _, answer := range answers {
		answer = strings.Trim(answer, "\n")
		numAnswered := len(strings.Split(answer, "\n"))
		answer = strings.ReplaceAll(answer, "\n", "")
		charHistogram := map[string]int{}
		for _, c := range answer {
			cString := string(c)
			if _, ok := charHistogram[cString]; !ok {
				charHistogram[cString] = 0
			}
			charHistogram[cString]++
		}
		totalUnique += len(charHistogram)
		for _, answeredForQuestion := range charHistogram {
			if answeredForQuestion == numAnswered {
				globallyUnique++
			}

		}
	}

	fmt.Println(fmt.Sprintf("Part 1: %d", totalUnique))
	fmt.Println(fmt.Sprintf("Part 2: %d", globallyUnique))
}
