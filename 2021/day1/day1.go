package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readInput(path string) ([]int, error) {
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
	ints := make([]int, len(text))
	for i, s := range text {
	  ints[i], _ = strconv.Atoi(s)
	}
	return ints, scanner.Err()
  }
  
  func main() {
	ints, err := readInput("input.txt")
	if err != nil {
	  log.Fatalf("readLines: %s", err)
	}
	increase := 0
	
	for i, output := range ints {
	  if i == 0 {
		continue
	  }
	  if output > ints[i - 1] {
		increase ++
	  }
	}
	print(increase)
  }