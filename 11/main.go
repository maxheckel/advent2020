package main

import (
	"fmt"
	"github.com/maxheckel/advent2020/common"
)

const (
	// uint8 of "#"
	occupied = 35
	// uint8 of "L"
	empty = 76
	// uint8 of "."
	floor = 46
)

func main() {
	rows := common.StringListFromFile("./11/input.txt")
	part1EndState := part1(rows)
	part2EndState := part2(rows)
	part1Occupied := 0
	part2Occupied := 0
	for _, row := range part1EndState {
		for _, char := range row {
			if char == occupied {
				part1Occupied++
			}
		}
	}
	for _, row := range part2EndState {
		for _, char := range row {
			if char == occupied {
				part2Occupied++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1Occupied)
	fmt.Printf("Part 2: %d\n", part2Occupied)
}

func part1(rows []string) []string {
	var postRun []string
	for y, row := range rows {
		newRow := ""
		for x, seat := range row {
			if seat == floor {
				newRow += string(floor)
				continue
			}
			if seat == empty && part1NumAdjacentSeatsOccupied(x, y, rows) == 0 {
				newRow += string(occupied)
				continue
			}
			if seat == occupied && part1NumAdjacentSeatsOccupied(x, y, rows) >= 4 {
				newRow += string(empty)
				continue
			}
			newRow += string(seat)
		}
		postRun = append(postRun, newRow)
	}
	for i, inp := range postRun {
		if rows[i] != inp {
			return part1(postRun)
		}
	}
	return postRun
}

func part2(rows []string) []string {
	var postRun []string
	for y, row := range rows {
		newRow := ""
		for x, seat := range row {
			if seat == floor {
				newRow += string(floor)
				continue
			}
			if seat == empty && part2NumAdjacentSeatsOccupied(x, y, rows) == 0 {
				newRow += string(occupied)
				continue
			}
			if seat == occupied && part2NumAdjacentSeatsOccupied(x, y, rows) >= 5 {
				newRow += string(empty)
				continue
			}
			newRow += string(seat)
		}
		postRun = append(postRun, newRow)
	}
	for i, inp := range postRun {
		if rows[i] != inp {
			return part2(postRun)
		}
	}
	return postRun
}

func part1NumAdjacentSeatsOccupied(x, y int, input []string) int {
	count := 0
	adjacentSpots := getAdjacentSeats(x, input, y)
	for _, spot := range adjacentSpots {
		if spot == occupied {
			count++
		}
	}
	return count
}


func part2NumAdjacentSeatsOccupied(x, y int, input []string) int {
	count := 0
	adjacentSpots := getVisibleSeats(x,  y, input)
	for _, spot := range adjacentSpots {
		if spot == occupied {
			count++
		}
	}
	return count
}


func getVisibleSeats(x, y int, input []string) []uint8 {
	var visibleSeats []uint8
	appendLeft := uint8(0)
	for index, char := range input[y]{
		if char == floor {
			continue
		}
		if index < x {
			appendLeft = input[y][index]
		}
		if index == x && appendLeft != 0 {
			visibleSeats = append(visibleSeats, appendLeft)
		}
		if index > x {
			visibleSeats = append(visibleSeats, input[y][index])
			break
		}
	}

	appendTop := uint8(0)
	for j, row := range input{
		if row[x] == floor {
			continue
		}
		if j < y {
			appendTop = row[x]
		}
		if j == y && appendTop != 0 {
			visibleSeats = append(visibleSeats, appendTop)
		}
		if j > y {
			visibleSeats = append(visibleSeats, row[x])
			break
		}
	}

	//
	bottomLeft := uint8(0)
	topLeft := uint8(0)
	bottomRight := uint8(0)
	topRight := uint8(0)
	inputLength := len(input)
	// left
	for i := 0; i < len(input[0]); i++ {
		if i < x {
			// Bottom left
			bottomLeftIndex := y - i + x
			if bottomLeftIndex >= 0 && bottomLeftIndex < inputLength && input[bottomLeftIndex][i] != floor{
				bottomLeft = input[bottomLeftIndex][i]
			}
			// Top left
			topLeftIndex := y - (x - i)
			if topLeftIndex >= 0 && topLeftIndex < inputLength && input[topLeftIndex][i] != floor {
				topLeft = input[topLeftIndex][i]
			}
		}
		if i == x {
			visibleSeats = append(visibleSeats, bottomLeft)
			visibleSeats = append(visibleSeats, topLeft)
		}

		if i > x {
			topRightIndex := y - (i - x)
			if topRightIndex >= 0 && topRightIndex < inputLength && input[topRightIndex][i] != floor && topRight == 0 {
				topRight = input[topRightIndex][i]
			}
			bottomRightIndex := y + (i - x)

			if bottomRightIndex >= 0 && bottomRightIndex < inputLength && input[bottomRightIndex][i] != floor && bottomRight == 0 {
				bottomRight = input[bottomRightIndex][i]
			}
		}
		if bottomRight != 0 && topRight != 0 {
			break
		}
	}
	visibleSeats = append(visibleSeats, topRight)
	visibleSeats = append(visibleSeats, bottomRight)


	return visibleSeats
}

func getAdjacentSeats(x int, input []string, y int) []uint8 {
	adjacentSpots := []uint8{}
	if x > 0 {
		adjacentSpots = append(adjacentSpots, input[y][x-1])
	}
	if y > 0 {
		adjacentSpots = append(adjacentSpots, input[y-1][x])
	}
	if x < len(input[0])-1 {
		adjacentSpots = append(adjacentSpots, input[y][x+1])
	}
	if y < len(input)-1 {
		adjacentSpots = append(adjacentSpots, input[y+1][x])
	}
	if x > 0 && y > 0 {
		adjacentSpots = append(adjacentSpots, input[y-1][x-1])
	}
	if x < len(input[0])-1 && y < len(input)-1 {
		adjacentSpots = append(adjacentSpots, input[y+1][x+1])
	}
	if y > 0 && x < len(input[0])-1 {
		adjacentSpots = append(adjacentSpots, input[y-1][x+1])
	}
	if x > 0 && y < len(input)-1 {
		adjacentSpots = append(adjacentSpots, input[y+1][x-1])
	}
	return adjacentSpots
}
