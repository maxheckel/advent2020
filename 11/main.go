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
			if seat == empty && adjacentSeatsOccupied(x, y, rows) == 0 {
				newRow += string(occupied)
				continue
			}
			if seat == occupied && adjacentSeatsOccupied(x, y, rows) >= 4 {
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
			if seat == empty && visibleSeatsOccupied(x, y, rows) == 0 {
				newRow += string(occupied)
				continue
			}
			if seat == occupied && visibleSeatsOccupied(x, y, rows) >= 5 {
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

func adjacentSeatsOccupied(x, y int, input []string) int {
	count := 0
	adjacentSpots := getAdjacentSeats(x, input, y)
	for _, spot := range adjacentSpots {
		if spot == occupied {
			count++
		}
	}
	return count
}


func visibleSeatsOccupied(x, y int, input []string) int {
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
	left := uint8(0)
	right := uint8(0)

	top := uint8(0)
	bottom := uint8(0)
	bottomLeft := uint8(0)
	topLeft := uint8(0)
	bottomRight := uint8(0)
	topRight := uint8(0)
	inputLength := len(input)
	// left
	for i := 0; i < len(input[0]); i++ {
		xChar := input[y][i]
		if i < x && xChar != floor{
			left = xChar
		}
		if i == x && left != 0 {
			visibleSeats = append(visibleSeats, left)
		}
		if i > x && xChar != floor && right == 0 {
			right = xChar
			visibleSeats = append(visibleSeats, right)
		}

		yChar := input[i][x]
		if i < y && yChar != floor {
			top = yChar
		}
		if i == y && top != 0 {
			visibleSeats = append(visibleSeats, top)
		}
		if i > y && yChar != floor && bottom == 0 {
			bottom = yChar
			visibleSeats = append(visibleSeats, bottom)
		}
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
