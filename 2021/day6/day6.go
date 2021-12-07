package main

import (
	"bufio"
	// "fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := strings.Split(scanner.Text(), ",")
	ints := make([]int, len(text))
	for i, s := range text {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints, scanner.Err()
}

func fish(numberOfDays int) {
	ints, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	ages := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, age := range ints {
		ages[age] ++
	}
	// fmt.Printf("%v\n", ages)
	for day := 0; day < numberOfDays; day++ {
		newFish := ages[0]
		ages[0] = 0
		for age := 1; age < len(ages); age ++ {
			ages[age - 1] = ages[age]
			ages[age] = 0
		}
		ages[6] += newFish
		ages[8] += newFish
		// fmt.Printf("day %v: %v\n", day, ages)
	}

	totalFish := 0
	for _, groups := range ages {
		totalFish += groups
	}
	println(totalFish)
}
func main() {
	//part 1
	fish(80)
	//part 2
	fish(256)
}
