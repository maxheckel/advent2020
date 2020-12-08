package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Operation struct {
	opType string
	opValue int
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
		op := &Operation{}
		fmt.Sscanf(scanner.Text(), "%s %d", &op.opType, &op.opValue)
		ops = append(ops, op)
	}
	result, _ := runOps(ops)
	fmt.Printf("Part 1: %d\n", result)

	terminated := false
	iterator := 0
	for !terminated {
		for ops[iterator].opType == "acc" {
			iterator++
		}
		swapper := strings.NewReplacer("jmp", "nop", "nop", "jmp")
		ops[iterator].opType = swapper.Replace(ops[iterator].opType)
		result, terminated = runOps(ops)
		// Switch it back
		ops[iterator].opType = swapper.Replace(ops[iterator].opType)
		iterator++
	}
	fmt.Printf("Part 2: %d", result)

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
