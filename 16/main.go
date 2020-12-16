package main

import (
	"fmt"
	"github.com/maxheckel/advent2020/common"
	"io/ioutil"
	"sort"
	"strings"
)

type Rule struct {
	name string
	validPositions map[int]bool
	validNums map[int]bool
}
func main() {
	file, _ := ioutil.ReadFile("./16/input.txt")
	parts := strings.Split(string(file), "\n\n")
	rawRules := strings.Split(parts[0], "\n")
	rules := getRules(rawRules)
	total, validTickets := part1(rules, parts)
	validTicketInts := [][]int{}
	for _, ticket := range validTickets{
		ticketStrings := strings.Split(ticket, ",")
		newTicketInts := []int{}
		for _, i := range ticketStrings{
			newTicketInts = append(newTicketInts, common.IntVal(i))
		}
		validTicketInts = append(validTicketInts, newTicketInts)
	}

	for x := 0; x < len(validTicketInts[0]); x++{

		for i, rule := range rules{
			validRule := true
			for i, ticket := range validTicketInts{
				if rule.name == "type" && i == 93 && x == 4{
					fmt.Println(ticket[x])
				}
				if !rule.validNums[ticket[x]] {
					validRule = false
					break
				}
			}
			if validRule {
				rules[i].validPositions[x] = true
			}

		}
	}


	for i := 0; i < 100; i++{
		for i, rule := range rules{
			if len(rule.validPositions) > 1{
				continue
			}
			removePosition := -1
			for r := range rule.validPositions{
				removePosition = r
			}
			for j := range rules{
				if i == j {
					continue
				}
				delete(rules[j].validPositions, removePosition)
			}
		}
	}
	sort.SliceStable(rules, func(i, j int) bool {
		return len(rules[i].validPositions) < len(rules[j].validPositions)
	})
	//for _, rule := range rules{
	//	fmt.Println(rule.name, rule.validPositions)
	//}

	fmt.Printf("Part 1: %d", total)
}

func difference(r1, r2 Rule) []int {
	foundDifs := []int{}
	for k := range r1.validPositions{
		if !r2.validPositions[k] {
			foundDifs = append(foundDifs, k)
		}
	}
	return foundDifs
}

func part1(rules []Rule, parts []string) (int, []string) {
	total := 0
	allTheRules := map[int]bool{}
	for _, rule := range rules {
		for k := range rule.validNums {
			allTheRules[k] = true
		}
	}
	validTickets := []string{}
	for _, ticket := range strings.Split(parts[2], "\n")[1:] {
		ticketNums := strings.Split(ticket, ",")
		isValid := true
		for _, num := range ticketNums {
			intNum := common.IntVal(num)
			if !allTheRules[intNum] {
				total += intNum
				isValid = false
			}
		}
		if isValid {
			validTickets = append(validTickets, ticket)
		}
	}

	return total, validTickets
}

func getRules(rawRules []string) []Rule {
	var rules []Rule
	for _, rule := range rawRules {
		groups := common.RE(`^([a-z\s]+): (\d+)-(\d+) or (\d+)-(\d+)$`).Groups(rule)

		rule := Rule{
			name: groups[0],
			validNums: map[int]bool{},
		}
		for x := common.IntVal(groups[1]); x <= common.IntVal(groups[2]); x++{
			rule.validNums[x] = true
		}
		for x := common.IntVal(groups[3]); x <= common.IntVal(groups[4]); x++{
			rule.validNums[x] = true
		}
		rule.validPositions = map[int]bool{}
		rules = append(rules, rule)
	}
	return rules
}
