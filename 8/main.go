package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	opType string
	opValue int
}

func OperationFromString(input string) *Operation {
	op := &Operation{}
	split := strings.Split(input, " ")
	op.opType = split[0]

	op.opValue, _ = strconv.Atoi(split[1][1:len(split[1])])
	if strings.Contains(split[1], "-") {
		op.opValue *= -1
	}
	return op
}

func main(){
	file, err := os.Open("./8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var ops []*Operation
	for scanner.Scan() {
		ops = append(ops, OperationFromString(scanner.Text()))
	}
	terminated := false

	result, _ := runOps(ops)
	fmt.Printf("Part 1: %d\n", result)
	iterator := 0
	for !terminated {
		for ops[iterator].opType == "acc" {
			iterator++
		}
		swapOp(ops, iterator)
		result, terminated = runOps(ops)
		// Switch it back
		swapOp(ops, iterator)
		iterator++
	}
	fmt.Printf("Part 2: %d", result)

}

func swapOp(copiedOps []*Operation, position int) {
	switch copiedOps[position].opType {
	case "jmp":
		copiedOps[position].opType = "nop"
		break
	case "nop":
		copiedOps[position].opType = "jmp"
	}
}

func runOps(ops []*Operation) (int, bool) {
	currentVal := 0
	currentPos := 0
	visited := map[*Operation]bool{}
	for visited[ops[currentPos]] == false {
		visited[ops[currentPos]] = true
		switch ops[currentPos].opType {
		case "nop":
			currentPos++
			break
		case "acc":
			currentVal += ops[currentPos].opValue
			currentPos++
			break
		case "jmp":
			currentPos += ops[currentPos].opValue
		}
		// Next instruction terminates
		if currentPos == len(ops) {
			return currentVal, true
		}
	}
	return currentVal, false
}
