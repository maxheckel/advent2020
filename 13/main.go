package main

import (
	"fmt"
	"github.com/maxheckel/advent2020/common"
	"math"
	"strconv"
	"strings"
)
const inputFile = "./13/sample3.txt"

func main() {
	goal, busIDs := getGoalAndIDs()
	solution := part1(busIDs, goal)
	fmt.Printf("Part 1: %d\n", solution)
	//allIDs := getIDsAndDiffs()
	//for i, busId := range allIDs{
	//	if busId == 0{
	//		continue
	//	}
	//
	//}
	//fmt.Println(endingNum)

}

type NumOrDifference struct {
	num float64
	diff int
}

func getIDsAndDiffs() map[int]int64  {
	lines := common.StringListFromFile(inputFile)
	var busIDs map[int]int64
	for i, id := range strings.Split(lines[1], ",") {
		if id == "x" {
			busIDs[i] = 0
		}
		busID, _ := strconv.Atoi(id)
		busIDs[i] = int64(busID)
	}
	return busIDs
}

func findSubsequentTime(num1 float64, num2 float64, diff int, minNum int) float64{
	found := false
	i := math.Ceil(float64(minNum) / num1)
	for !found{
		if math.Mod((num1*i)+float64(diff), num2) == 0 {
			break
		}
		i++
	}
	return num1*i+1
}

func part1(busIDs []float64, goal float64) int {
	max := float64(0)
	var smallestDifID float64
	for _, id := range busIDs {
		divisions := goal / id
		mod := math.Mod(divisions, 1)
		if mod > max {
			max = mod
			smallestDifID = id
		}
	}
	diff := int(goal / smallestDifID)
	solution := ((diff+1)*int(smallestDifID) - int(goal)) * int(smallestDifID)
	return solution
}

func getGoalAndIDs() (float64, []float64) {
	lines := common.StringListFromFile(inputFile)
	goal, _ := strconv.Atoi(lines[0])
	var busIDs []float64
	for _, id := range strings.Split(lines[1], ",") {
		if id == "x" {
			continue
		}
		busID, _ := strconv.Atoi(id)
		busIDs = append(busIDs, float64(busID))
	}
	return float64(goal), busIDs
}
