package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Coords struct {
	X, Y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	scanner := bufio.NewScanner(file)
	fabric := make([][]int, 1000*1000)

	overlappedCells := map[string][]Coords{}
	overlappedClaims := map[int]bool{}

	var cleanClaim int

	for scanner.Scan() {
		line := scanner.Text()
		id, startCoords, endCoords := parseCoords(line)

		for i := startCoords.X; i <= endCoords.X; i++ {
			for j := startCoords.Y; j <= endCoords.Y; j++ {

				cell := fabric[i*1000+j]
				cellName := strconv.Itoa(i) + `x` + strconv.Itoa(j)

				fabric[i*1000+j] = append(cell, id)

				if len(fabric[i*1000+j]) > 1 {
					overlappedCells[cellName] = append(overlappedCells[cellName], Coords{i, j})

					for _, t := range fabric[i*1000+j] {
						overlappedClaims[t] = true
					}

				}

			}
		}
	}

	for i := 1; i <= 1295; i++ {

		fmt.Println(overlappedClaims[i])

		if overlappedClaims[i] {
			continue
		}

		cleanClaim = i
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("File scanning error", err)
	}

	fmt.Println("Overlapped Cells: ", len(overlappedCells))
	fmt.Println("Clean Claim: ", cleanClaim)

}

func parseCoords(str string) (int, Coords, Coords) {

	var re = regexp.MustCompile(`#(?m)([0-9]{0,4}) @ (?m)([0-9]{0,4}),(?m)([0-9]{0,4}): (?m)([0-9]{0,4})x(?m)([0-9]{0,4})`)
	stringParsed := re.FindStringSubmatch(str)

	id, _ := strconv.Atoi(stringParsed[1])
	leftGap, _ := strconv.Atoi(stringParsed[2])
	topGap, _ := strconv.Atoi(stringParsed[3])
	width, _ := strconv.Atoi(stringParsed[4])
	height, _ := strconv.Atoi(stringParsed[5])

	startCoords := Coords{leftGap + 1, topGap + 1}
	endCoords := Coords{leftGap + width, topGap + height}

	return id, startCoords, endCoords
}
