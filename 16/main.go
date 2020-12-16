package main

import (
	"fmt"
	"github.com/maxheckel/advent2020/common"
	"io/ioutil"
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
	myTicketRaw := parts[1]
	myTicket := strings.Split(strings.Split(myTicketRaw, "\n")[1], ",")
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
			for _, ticket := range validTicketInts{
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

	numRuns := len(rules[0].validPositions)
	for i := 0; i < numRuns; i++{
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

	runningVal := 1
	for _, rule := range rules{
		if strings.Contains(rule.name, "departure") {
			locVal := -1
			for v := range rule.validPositions{
				locVal = v
			}
			runningVal*=common.IntVal(myTicket[locVal])
		}
	}
	fmt.Printf("Part 1: %d\n", total)
	fmt.Printf("Part 2: %d\n", runningVal)
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
