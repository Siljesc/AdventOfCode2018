package main

import (
	"fmt"
	"strconv"

	"../util"
)

func scanLines(path string, freq int, freqDict map[int]bool) (int, error) {

	scanner, _ := util.ReadFile(path)

	var firstRepeated int
	var foundValue = false

	for scanner.Scan() {
		lineFreq, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return 0, err
		}

		freq += lineFreq

		_, freqItem := freqDict[freq]

		if freqItem {
			firstRepeated = freq
			foundValue = true
			break
		}

		freqDict[freq] = true
	}

	if !foundValue {
		return scanLines(path, freq, freqDict)
	}

	return firstRepeated, nil
}

func part1() int {

	scanner, _ := util.ReadFile("input.txt")
	freq := 0

	for scanner.Scan() {
		lineFreq, _ := strconv.Atoi(scanner.Text())

		freq += lineFreq
	}

	return freq
}

func part2() int {
	freq := 0
	freqDict := map[int]bool{0: true}

	firstRepeated, _ := scanLines("input.text", freq, freqDict)

	return firstRepeated
}

func main() {
	fmt.Println("[Sol 1] Frequency: ", part1())
	fmt.Println("[Sol 2] First Repeated Frequency: ", part2())
}
