package main

import (
	"bufio"
	"log"
	"math"
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
func max(a []int) int {
	max := a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}
func sum_up_to(num int) int {
	return int(0.5 * float64((num-1)*num))
}

func part1() {
	ints, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	minFuel := 999999999
	maximum := max(ints)
	for newPosition := 0; newPosition < maximum; newPosition++ {
		fuel := 0
		for pos := 0; pos < len(ints); pos++ {
			fuel += int(math.Abs(float64(newPosition - ints[pos])))
		}
		if minFuel > fuel {
			minFuel = fuel
		}
	}
	println(minFuel)
}

func part2() {
	ints, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	minFuel := 999999999
	maximum := max(ints)
	for newPosition := 0; newPosition < maximum; newPosition++ {
		fuel := 0
		for pos := 0; pos < len(ints); pos++ {
			fuel += sum_up_to(int(math.Abs(float64(newPosition-ints[pos]))) + 1)
		}
		// print(fuel)
		if minFuel > fuel {
			minFuel = fuel
		}
	}
	println(minFuel)
}

func main() {
	part1()
	part2()
}
