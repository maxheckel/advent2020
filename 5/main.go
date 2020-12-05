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
		seatNumbers = append(seatNumbers, getSeatNumber(scanner.Text()))
	}
	sort.Ints(seatNumbers)
	mySeat := getMissingSeat(seatNumbers)
	fmt.Println(fmt.Sprintf("Part 1: %d", seatNumbers[len(seatNumbers)-1]))
	fmt.Println(fmt.Sprintf("Part 2: %d", mySeat))

}

func getMissingSeat(seatNumbers []int) int {

	mySeat := 0
	for i, sn := range seatNumbers {
		if i == len(seatNumbers)-1 {
			continue
		}
		if seatNumbers[i+1] != sn+1 {
			mySeat = sn + 1
			break
		}
	}
	return mySeat
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
		rowsChanged := math.Ceil((seats.endRow - seats.startRow) / 2)
		colsChanged := math.Ceil((seats.endCol - seats.startCol) / 2)
		switch string(char) {
		case "F":
			seats.endRow -= rowsChanged
		case "B":
			seats.startRow += rowsChanged
		case "L":
			seats.endCol -= colsChanged
		case "R":
			seats.startCol += colsChanged
		}
	}

	return int(seats.endRow * 8 + seats.endCol)
}
