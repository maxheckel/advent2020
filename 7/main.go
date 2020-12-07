package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func lineToBag(line string) *Bag {
	line = strings.ReplaceAll(line, "bags", "")
	line = strings.ReplaceAll(line, "bag", "")
	line = strings.ReplaceAll(line, ".", "")
	split := strings.Split(line, "contain")
	rootBagColor := split[0]
	bag := &Bag{
		color:        strings.Trim(rootBagColor, " "),
		count:        1,
		children:     []*Bag{},
		flatChildren: []Bag{},
	}
	if strings.Contains(line, "no other") {
		return bag
	}
	contains := strings.Split(split[1], ",")

	for _, child := range contains{
		child = strings.Trim(child, " ")
		bgRe := regexp.MustCompile(`^(\d+) (\D+)$`)
		matches := bgRe.FindAllSubmatch([]byte(child), -1)
		childBag := &Bag{}
		for _, m := range matches {
			num, _ := strconv.Atoi(string(m[1]))
			childBag.count = num
			childBag.color = string(m[2])
		}
		bag.flatChildren = append(bag.flatChildren, *childBag)
	}
	return bag
}


type Bag struct{
	color        string
	count        int
	children     []*Bag
	flatChildren []Bag
}

func main() {
	file, err := os.Open("./7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)


	bags := []*Bag{}
	for scanner.Scan() {
		bags = append(bags, lineToBag(scanner.Text()))
	}
	// Build the tree
	for _, bag := range bags{
		for i, check := range bags{
			for _, childBag := range check.flatChildren{
				if childBag.color == bag.color {
					bags[i].children = append(bags[i].children, bag)
				}
			}
		}
	}

	goldCount := 0
	goldSubCount := 0
	for _, bag := range bags {
		if bag.color == "shiny gold" {
			for _, child := range bag.children{
				count := findCountWithinBag(bag, child)
				goldSubCount += count + (count * findNumBagsInBag(child))
			}
		}

		if hasGoldAncestors(bag.children) {
			goldCount++
		}

	}

	fmt.Printf("Part 1: %d\n", goldCount)
	fmt.Printf("Part 2: %d\n", goldSubCount)
}

func findCountWithinBag(bag *Bag, child *Bag) int {
	count := 0
	for _, fchild := range bag.flatChildren {
		if fchild.color == child.color {
			count = fchild.count
		}
	}
	return count
}

func findNumBagsInBag(bag *Bag) int {
	if len(bag.children) == 0 {
		return 0
	}
	total := 0
	for _, child := range bag.children{
		count := findCountWithinBag(bag, child)
		total += count + (count * findNumBagsInBag(child))
	}
	return total
}

func hasGoldAncestors(bags []*Bag) bool{
	for _, bag := range bags {
		if bag.color == "shiny gold"{
			return true
		}
		if len(bag.children) > 0 {
			if hasGoldAncestors(bag.children){
				return true
			}
		}
	}
	return false

}
