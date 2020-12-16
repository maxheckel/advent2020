package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func IntListFromFile(filepath string) []int{
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	var intlist []int
	split := strings.Split(string(file), "\n")
	for _, row := range split {
		intVal, _ := strconv.Atoi(row)
		intlist = append(intlist, intVal)
	}
	return intlist
}

func StringListFromFile(filepath string) []string{
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	var stringList []string
	split := strings.Split(string(file), "\n")
	for _, row := range split {
		stringList = append(stringList, row)
	}
	return stringList
}


type regex struct {
	r *regexp.Regexp
}

func RE(pattern string) *regex {
	return &regex{regexp.MustCompile(pattern)}
}

func (r *regex) Test(s string) bool {
	return r.r.MatchString(s)
}

func (r *regex) Groups(s string) []string {
	match := r.r.FindStringSubmatch(s)
	if match == nil {
		panic(fmt.Sprintf("\"%s\" did not match pattern \"%v\"", s, r.r))
	}
	if len(match) == 0 {
		panic("No groups found.")
	}
	return match[1:]
}

func IntVal(input string) int{
	res, _ := strconv.Atoi(input)
	return res
}
