package main

import (
	"fmt"
	"github.com/maxheckel/advent2020/common"
	"sort"
)

func main() {
	input := common.IntListFromFile("./10/example1.txt")
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)
	part1Res := part1(input)
	count := part2(input)

	fmt.Printf("Part 1: %d\n", part1Res)
	fmt.Printf("Part 2: %d", count)
}

func part2(input []int) int {
	memo := map[int]int{0: 1}
	for _, v := range input[1:] {
		memo[v] = memo[v-1] + memo[v-2] + memo[v-3]

	}
	fmt.Println(memo)
	return memo[input[len(input)-1]]
}

func part1(input []int) int {
	threeJumps := 0
	oneJumps := 0
	twoJumps := 0
	currentVal := 0
	for _, val := range input {
		if val == currentVal {
			continue
		}
		if currentVal+1 == val {
			oneJumps++
			currentVal = val
			continue
		}
		if currentVal+2 == val {
			twoJumps++
			currentVal = val
			continue
		}
		if currentVal+3 == val {
			threeJumps++
			currentVal = val
			continue
		}
	}
	return oneJumps * threeJumps
}
