package main

import "fmt"

func main() {
	//seen := map[int]int{
	//	3: 1,
	//	2: 2,
	//	1: 3,
	//}
	seen := map[int]int{
		9: 1,
		3: 2,
		1: 3,
		0: 4,
		8: 5,
		4: 6,
	}
	// copy the map
	seen2 := map[int]int{}
	for k, v := range seen{
		seen2[k]=v
	}

	part1 := playGame(seen, 2020)
	part2 := playGame(seen2, 30000000)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func playGame(seen map[int]int, iterations int) int {
	lastPos := seen[len(seen)-1]
	newPos := 0
	for i := len(seen) + 1; i <= iterations; i++ {
		lastPos = newPos
		oldVal, seenBefore := seen[lastPos]
		if seenBefore {
			newPos = i - oldVal
		} else {
			newPos = 0
		}
		seen[lastPos] = i
	}
	return lastPos
}

