package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var text []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text, scanner.Err()
}
func part1 () {
	text, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	horizontal := 0
	depth := 0
	for i := 0; i < len(text); i++ {
		if strings.HasPrefix(text[i], "forward ") {
			forward := 0
			forward, _ = strconv.Atoi(strings.Replace(text[i], "forward ", "", 1))
			horizontal += forward
		} else if strings.HasPrefix(text[i], "down ") {
			down := 0
			down, _ = strconv.Atoi(strings.Replace(text[i], "down ", "", 1))
			depth += down
		} else if strings.HasPrefix(text[i], "up ") {
			up := 0
			up, _ = strconv.Atoi(strings.Replace(text[i], "up ", "", 1))
			depth -= up
		} 
	}
	println(horizontal * depth)
}

func part2 () {
	text, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	horizontal := 0
	depth := 0
	aim := 0
	for i := 0; i < len(text); i++ {
		if strings.HasPrefix(text[i], "forward ") {
			forward := 0
			forward, _ = strconv.Atoi(strings.Replace(text[i], "forward ", "", 1))
			horizontal += forward
			depth += aim * forward
		} else if strings.HasPrefix(text[i], "down ") {
			down := 0
			down, _ = strconv.Atoi(strings.Replace(text[i], "down ", "", 1))
			aim += down
		} else if strings.HasPrefix(text[i], "up ") {
			up := 0
			up, _ = strconv.Atoi(strings.Replace(text[i], "up ", "", 1))
			aim -= up
		} 
	}
	println(horizontal * depth)
}

func main() {
	part1()
	part2()
}
