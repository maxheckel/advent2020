package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	globalStart := time.Now()
	cypher := getCypher()

	algoStart := time.Now()
	badNumber := part1(cypher)
	globalElapsed := time.Since(globalStart)
	part1Elapsed := time.Since(algoStart)
	fmt.Printf("Part 1 took overall %s\n", globalElapsed)
	fmt.Printf("Part1 Algo took %s\n", part1Elapsed)
	fmt.Printf("Part 1: %d\n", badNumber)

	for i, val := range cypher{
		if val > badNumber {
			cypher = cypher[0:i]
			break
		}
	}
	result := part2(cypher, badNumber)
	globalElapsed = time.Since(globalStart)
	part2Elapsed := time.Since(algoStart)
	fmt.Printf("\n\nPart 2 took overall %s\n", globalElapsed)
	fmt.Printf("Part 2 Algo took %s\n", part2Elapsed)
	fmt.Printf("Part 2: %d", result)
}

func getCypher() []int {
	file, err := ioutil.ReadFile("./9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var cypher []int
	split := strings.Split(string(file), "\n")
	for _, row := range split {
		intVal, _ := strconv.Atoi(row)
		cypher = append(cypher, intVal)
	}
	return cypher
}

func part2(cypher []int, badNumber int) int {
	oleft, oright, left, right, sum, osum := 0, 0, len(cypher)-1, len(cypher)-1, 0, 0
	for true {
		if osum < badNumber {
			osum += cypher[oright]
			oright++
		} else if osum > badNumber {
			osum -= cypher[oleft]
			oleft++
		}

		if sum < badNumber {
			sum += cypher[right]
			right--
		} else if sum > badNumber {
			sum -= cypher[left]
			left--
		}

		if osum == badNumber {
			sort.Ints(cypher[oleft:oright])
			return cypher[oleft:oright][0] + cypher[oleft:oright][len(cypher[oleft:oright])-1]
		}
		if sum == badNumber {
			sort.Ints(cypher[right:left])
			return cypher[right:left][0] + cypher[right:left][len(cypher[right:left])-1]
		}
	}
	return 0
}

func part1(cypher []int) int {
	preambleSize := 25
	badNumber := 0
	for i, _ := range cypher {
		if i+preambleSize == len(cypher) {
			break
		}
		subSlice := cypher[i : i+preambleSize]
		numToCheck := cypher[i+preambleSize]
		if !canBeSummed(numToCheck, subSlice) {
			badNumber = numToCheck
			break
		}
	}
	return badNumber
}

func canBeSummed(goal int, nums []int) bool {
	for i, num1 := range nums{
		for _, num2 := range nums[i:]{
			if num1+num2 == goal {
				return true
			}
		}
	}
	return false
}
