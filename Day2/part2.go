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

	if err := scanner.Err(); err != nil {
		fmt.Println("File scanning error", err)
	}

	fmt.Println("Box Id: ", commonID)

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
