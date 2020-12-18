package main

import (
	"fmt"
	"github.com/maxheckel/advent2020/common"
	"io/ioutil"
	"strconv"
	"strings"
)

type Equation struct {
	left     int
	right    *Equation
	operator string
}

func main() {
	input, _ := ioutil.ReadFile("./18/sample.txt")
	equations := strings.Split(string(input), "\n")
	equations = equations[:len(equations)-1]
	total := 0
	for _, eq := range equations {
		eq := strings.Replace(eq, " ", "", -1)
		for strings.Contains(eq, "("){
			eq = solve(eq)
		}
		total += solveSimple(eq)
	}
	fmt.Printf("Part 1: %d", total)
}
type ParenLoc struct {
	paren string
	loc int
}
func solve(eq string) string {

	if strings.Contains(eq, ")") {
		left := []int{}
		right := []int{}
		stack := []ParenLoc{}
		replacements := map[string]string{}
		for i, char := range eq {
			if string(char) == "(" {
				left = append(left, i+1)
				stack = append(stack, ParenLoc{
					paren: "(",
					loc: i+1,
				})
			}
			if string(char) == ")" {
				right = append(right, i)
				stack = append(stack, ParenLoc{
					paren: ")",
					loc: i,
				})
			}

			for len(stack) > 1 && stack[len(stack)-1].paren == ")" && stack[len(stack)-2].paren == "(" {
				l := stack[len(stack)-2].loc
				r := stack[len(stack)-1].loc
				stack = stack[:len(stack)-2]
				if strings.Contains(eq[l:r], "(") {
					replacements[eq[l-1:r+1]] = "("+solve(eq[l:r])+")"
				} else {
					replacements[eq[l-1:r+1]] = strconv.Itoa(solveSimple(eq[l:r]))
				}
			}
		}

		for k, v := range replacements {
			eq = strings.Replace(eq, k, v, 1)
		}
	}
	return eq
}

func solveSimple(input string) int {
	nums := strings.FieldsFunc(input, func(r rune) bool {
		return string(r) == "*" || string(r) == "+"
	})
	operators := []string{}
	for _, char := range input {
		if string(char) == "*" || string(char) == "+" {
			operators = append(operators, string(char))
		}
	}
	left := common.IntVal(nums[0])
	nums = nums[1:]
	for _, num := range nums {
		switch operators[0] {
		case "+":
			left += common.IntVal(num)
		case "*":
			left *= common.IntVal(num)
		}
		operators = operators[1:]
	}
	return left
}

func solveComplex(input string) int {
	nums := strings.FieldsFunc(input, func(r rune) bool {
		return string(r) == "*" || string(r) == "+"
	})
	operators := []string{}
	for _, char := range input {
		if string(char) == "*"{
			operators = append(operators, string(char))
		}

		if string(char) == "+" {
			operators = append(operators, string(char))
		}
	}
	stack := []string{}
	for i, o := range operators{
		if len(stack) == 0 {
			stack = append(stack, nums[i], nums[i+1], o)
		} else {
			stack = append(stack, nums[i+1], o)
		}
	}
	fmt.Println(stack)
	return 0
}
