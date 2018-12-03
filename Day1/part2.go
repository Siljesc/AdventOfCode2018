package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	freq := 0
	freqDict := map[int]bool{0: true}

	firstRepeated, err := scanLines(file, freq, freqDict)

	if err != nil {
		fmt.Println("File scanning error", err)
	}

	fmt.Println("First Repeated: ", firstRepeated)

}

func scanLines(file *os.File, freq int, freqDict map[int]bool) (int, error) {

	scanner := bufio.NewScanner(file)
	file.Seek(0, 0)

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
		return scanLines(file, freq, freqDict)
	}

	return firstRepeated, nil
}
