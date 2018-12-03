package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	scanner := bufio.NewScanner(file)
	twoWords := 0
	threeWords := 0

	for scanner.Scan() {
		word := scanner.Text()
		letterList := strings.Split(word, "")
		letterDict := map[string]int{}

		counted2 := false
		counted3 := false

		for _, letter := range letterList {

			_, ok := letterDict[letter]

			if !ok {
				letterDict[letter] = 0
			}

			letterDict[letter]++
		}

		for _, frequency := range letterDict {
			if frequency == 2 && !counted2 {
				twoWords++
				counted2 = true
			} else if frequency == 3 && !counted3 {
				threeWords++
				counted3 = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("File scanning error", err)
	}

	fmt.Println("Checksum: ", twoWords*threeWords)

}
