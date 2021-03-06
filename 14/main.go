package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)
type Assignment struct {
	location uint64
	runeVal  []rune
	value    string
	intVal   uint64
}
func main() {

	file, err := os.Open("./14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var currentMask string
	mem := map[uint64]uint64{}
	mem2 := map[string]uint64{}
	algoStart := time.Now()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "mask") {
			fmt.Sscanf(line, "mask = %s", &currentMask)
		}
		if strings.Contains(line, "mem") {
			assn := part1(line, currentMask, mem)
			part2(line, assn, currentMask, mem2)
		}
	}

	total := uint64(0)
	for _, val := range mem {
		total += val
	}
	total2 := uint64(0)
	for _, val := range mem2 {
		total2 += val
	}

	elapsed := time.Since(algoStart)
	fmt.Printf("Part 1: %d\n", total)
	fmt.Printf("Part 2: %d\n", total2)

	fmt.Printf("Time taken %s\n", elapsed)
	PrintMemUsage()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func part2(line string, assn Assignment, currentMask string, mem2 map[string]uint64) {
	fmt.Sscanf(line, "mem[%d] = %d", &assn.location, &assn.intVal)
	unpaddedBinary := strconv.FormatUint(assn.location, 2)
	locationBinary := []rune(fmt.Sprintf("%0"+strconv.Itoa(36)+"s", unpaddedBinary))
	for i, r := range currentMask {
		if string(r) == "0" {
			continue
		}
		locationBinary[i] = r
	}

	numX := strings.Count(string(locationBinary), "X")
	permutations := uint64(math.Pow(2, float64(numX)))
	for i := uint64(0); i < permutations; i++ {
		newLocation := locationBinary
		injection := fmt.Sprintf("%0"+strconv.Itoa(numX)+"s", strconv.FormatUint(i, 2))
		for _, p := range injection {
			newLocation = []rune(strings.Replace(string(newLocation), "X", string(p), 1))
		}
		mem2[string(newLocation)] = assn.intVal
	}
}

func part1(line string, currentMask string, mem map[uint64]uint64) Assignment {
	assn := Assignment{}
	var val uint64
	fmt.Sscanf(line, "mem[%d] = %d", &assn.location, &val)
	assn.runeVal = []rune(fmt.Sprintf("%0"+strconv.Itoa(36)+"s", strconv.FormatUint(val, 2)))
	for i, r := range currentMask {
		if string(r) == "X" {
			continue
		}
		assn.runeVal[i] = r
	}
	assn.value = string(assn.runeVal)
	mem[assn.location], _ = strconv.ParseUint(assn.value, 2, 64)
	return assn
}
