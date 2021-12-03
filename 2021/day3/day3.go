package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
func part1() {
	text, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	gammaRate := ""
	epsilonRate := ""
	for i := 0; i < len(text[0]); i++ {
		ones := 0
		zeros := 0
		for j := 0; j < len(text); j++ {
			if text[j][i] == '1' {
				ones++
			} else {
				zeros++
			}
		}
		if ones > zeros {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}
	gammaRateInt, _ := strconv.ParseInt(gammaRate,2,64)
	epsilonRateInt, _ := strconv.ParseInt(epsilonRate,2,64)
	println(gammaRateInt * epsilonRateInt)
}
func part2 (){
	text, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	oxGenSlice := oxGen(text,0)
	coScrubSlice := coScrub(text,0)
	oxGenBit := oxGenSlice[0]
	coScrubBit := coScrubSlice[0]
	oxGenInt, _ := strconv.ParseInt(oxGenBit,2,64)
	coScrubInt, _ := strconv.ParseInt(coScrubBit,2,64)
	println(oxGenInt * coScrubInt)
}
// for part 2
func oxGen(bits []string, index int) ([]string){
	if len(bits) == 1 {
		return bits
	} else {
		var newBits []string
		zeros := 0
		ones := 0
		for bit := 0; bit < len(bits); bit++ {
			if bits[bit][index] == '1' {
				ones ++
			} else {
				zeros ++
			}
		}
		if zeros <= ones {
			for bit := 0; bit < len(bits); bit++ {
				if bits[bit][index] == '1' {
					newBits = append(newBits,bits[bit])
				}
			}
		} else {
			for bit := 0; bit < len(bits); bit++ {
				if bits[bit][index] == '0' {
					newBits = append(newBits,bits[bit])
				}
			}
		}
		return oxGen(newBits, index +1)
	}
}
// for part 2
func coScrub(bits []string, index int) ([]string){
	if len(bits) == 1 {
		return bits
	} else {
		var newBits []string
		zeros := 0
		ones := 0
		for bit := 0; bit < len(bits); bit++ {
			if bits[bit][index] == '1' {
				ones ++
			} else {
				zeros ++
			}
		}
		if zeros <= ones {
			for bit := 0; bit < len(bits); bit++ {
				if bits[bit][index] == '0' {
					newBits = append(newBits,bits[bit])
				}
			}
		} else {
			for bit := 0; bit < len(bits); bit++ {
				if bits[bit][index] == '1' {
					newBits = append(newBits,bits[bit])
				}
			}
		}
		return coScrub(newBits, index +1)
	}
}

func main() {
	part1()
	part2()
}
