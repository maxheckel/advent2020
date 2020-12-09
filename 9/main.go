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
	file, err := ioutil.ReadFile("./9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var cypher []int
	split := strings.Split(string(file), "\n")
	for _, row := range split{
		intVal, _ := strconv.Atoi(row)
		cypher = append(cypher, intVal)
	}

	algoStart := time.Now()
	preambleSize := 25
	badNumber := 0
	for i, _ := range cypher {
		if i+preambleSize == len(cypher) {
			break
		}
		subSlice := cypher[i:i+preambleSize]
		numToCheck := cypher[i+preambleSize]
		if !canBeSummed(numToCheck, subSlice) {
			badNumber = numToCheck

			globalElapsed := time.Since(globalStart)
			part1Elapsed := time.Since(algoStart)
			fmt.Printf("Part 1 took overall %s\n", globalElapsed)
			fmt.Printf("Part1 Algo took %s\n", part1Elapsed)
			fmt.Printf("Part 1: %d\n", numToCheck)
			break
		}
	}
	left, right := 0, 0
	sum := 0
	for right < len(cypher) {
		if sum < badNumber {
			sum += cypher[right]
			right++
		} else if sum > badNumber {
			sum -= cypher[left]
			left++
		} else {
			sort.Ints(cypher[left:right])
			globalElapsed := time.Since(globalStart)
			part2Elapsed := time.Since(algoStart)
			fmt.Printf("\n\nPart 2 took overall %s\n", globalElapsed)
			fmt.Printf("Part 2 Algo took %s\n", part2Elapsed)
			fmt.Printf("Part 2: %d", cypher[left:right][0] + cypher[left:right][len(cypher[left:right])-1])
			break
		}
	}


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
