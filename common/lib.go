package common

import (
	"io/ioutil"
	"log"
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