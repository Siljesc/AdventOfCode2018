package main

import (
	"fmt"
	"strings"

	"../util"
)

func part1() int {
	scanner, _ := util.ReadFile("input.txt")

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

	return twoWords * threeWords

}

func checkIds(id1 []string, id2 []string) string {
	var commonLetters strings.Builder

	for index, letter := range id1 {
		if letter == id2[index] {
			commonLetters.WriteString(letter)
		}
	}

	return commonLetters.String()
}

func part2() string {
	scanner, _ := util.ReadFile("input.txt")

	var lines [][]string
	var checkLines [][]string
	var commonID string

	for scanner.Scan() {
		word := scanner.Text()
		lettersList := strings.Split(word, "")
		lines = append(lines, lettersList)
		checkLines = append(checkLines, lettersList)
	}

	for i, letterList1 := range lines {
		if len(commonID) == 25 || i == (len(lines)-1) {
			break
		}

		checkLines = append(lines[:i], lines[i+1:]...)

		for _, letterList2 := range checkLines {
			letters := checkIds(letterList1, letterList2)

			if len(letters) == 25 {
				commonID = letters
				break
			}
		}
	}

	return commonID

}

func main() {
	fmt.Println("[Sol 1] Checksum: ", part1())
	fmt.Println("[Sol 2] Common ID: ", part2())
}
