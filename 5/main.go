package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("./5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)


	seatNumbers := []int{}
	for scanner.Scan() {
		seatNumber := getSeatNumber(scanner.Text())
		seatNumbers = append(seatNumbers, seatNumber)
	}
	sort.Ints(seatNumbers)
	mySeat := 0
	for i, sn := range seatNumbers{
		if i == len(seatNumbers)-1 {
			continue
		}
		if seatNumbers[i+1] != sn+1 {
			mySeat = sn+1
			break
		}
	}
	fmt.Println(fmt.Sprintf("Part 1: %d", seatNumbers[len(seatNumbers)-1]))
	fmt.Println(fmt.Sprintf("Part 2: %d", mySeat))

}

type AvailableSeats struct {
	startRow float64
	endRow float64
	startCol float64
	endCol float64
}

func getSeatNumber(directions string) int {
	seats := AvailableSeats{
		startRow: 0,
		endRow:   127,
		startCol: 0,
		endCol:   7,
	}
	for _, char := range directions{
		switch string(char) {
		case "F":
			seats.endRow -= math.Ceil((seats.endRow-seats.startRow)/2)
		case "B":
			seats.startRow += math.Ceil((seats.endRow-seats.startRow)/2)
		case "L":
			seats.endCol -= math.Ceil((seats.endCol-seats.startCol)/2)
		case "R":
			seats.startCol += math.Ceil((seats.endCol-seats.startCol)/2)
		}
	}

	return int(math.Min(seats.endRow, seats.startRow) * 8 + math.Min(seats.startCol, seats.endCol))
}
