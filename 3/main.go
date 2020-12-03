package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	leftIndex := 0
	line := 0
	numTrees := 0
	for scanner.Scan() {

		currentLine := scanner.Text()
		for len(currentLine) < leftIndex && line > 0 {
			currentLine+=currentLine
		}
		if string(currentLine[leftIndex]) == "#" {
			numTrees++
		}
		line++
		leftIndex+=3
	}
	fmt.Println(numTrees)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
