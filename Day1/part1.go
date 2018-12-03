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

	scanner := bufio.NewScanner(file)
	freq := 0

	for scanner.Scan() {
		lineFreq, lineErr := strconv.Atoi(scanner.Text())

		if lineErr != nil {
			fmt.Println("String casting to int error", err)
			return
		}

		freq += lineFreq
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("File scanning error", err)
	}

	fmt.Println("Frequency: ", freq)

}
