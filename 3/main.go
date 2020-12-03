package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Rule struct{
	x int
	y int
}

func main() {
	file, err := os.Open("./3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}


	rules := []Rule{
		{
			x: 1,
			y: 1,
		},
		{
			x: 3,
			y: 1,
		},
		{
			x: 5,
			y: 1,
		},
		{
			x: 7,
			y: 1,
		},
		{
			x: 1,
			y: 2,
		},
	}
	total := []int{}
	for _, r := range rules {
		numTrees := 0
		leftIndex := 0
		for i, currentLine := range rows{
			if r.y > 1 && i % r.y != 0{
				continue
			}
			for len(currentLine) <= leftIndex && i > 0 {
				currentLine+=currentLine
			}
			if string(currentLine[leftIndex]) == "#" {
				numTrees++
			}
			leftIndex+=r.x
		}
		total = append(total, numTrees)
	}
	fmt.Println(fmt.Sprintf("Part 1: %d", total[1]))
	fmt.Println(fmt.Sprintf("Part 2: %d", total[0]*total[1]*total[2]*total[3]*total[4]))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
