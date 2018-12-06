package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../util"
)

func parseData(str string) (int, string) {
	re := regexp.MustCompile(`(?m)\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:([0-9]{2})\] (.*)`)

	stringParsed := re.FindStringSubmatch(str)

	minute, _ := strconv.Atoi(stringParsed[1])
	msg := stringParsed[2]

	return minute, msg
}

func getGuardID(str string) int {
	re := regexp.MustCompile(`Guard #([0-9]{0,4}) begins shift`)
	stringParsed := re.FindStringSubmatch(str)

	id, _ := strconv.Atoi(stringParsed[1])

	return id
}

func part1() int {

	scanner, _ := util.ReadFile("input.txt")

	guards := map[int]map[int]int{}

	var currentGuard int
	var sleepStart int

	for scanner.Scan() {
		line := scanner.Text()

		minute, msg := parseData(line)

		if strings.Contains(msg, "Guard") {
			currentGuard = getGuardID(msg)

			if _, ok := guards[currentGuard]; !ok {
				guards[currentGuard] = map[int]int{}
			}

		} else if msg == "falls asleep" {
			sleepStart = minute
		} else if msg == "wakes up" {

			for i := sleepStart; i <= minute; i++ {
				if _, ok := guards[currentGuard][i]; !ok {
					guards[currentGuard][i] = 1
				} else {
					guards[currentGuard][i]++
				}
			}
		}
	}

	fmt.Println(guards)

	return 1

}

func main() {
	fmt.Println("[Sol 1] Guard ID * Minute: ", part1())
}
