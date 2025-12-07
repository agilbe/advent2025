package internal

import (
	"strconv"
	"strings"
)

var numTimesTotalAt0 int = 0

func SolveDay01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	// dial starts at 50
	// turn left/right according to each line
	// if it crosses left of 0 once, goes to 99
	// count how many times it goes to 0
	numTimesEndingAt0 := 0
	position := 50
	for _, line := range lines {
		instruction := []rune(line)
		direction := string(instruction[0])
		movement, _ := strconv.Atoi(string(instruction[1:]))
		position = calculatePosition(position, direction, movement)
		if position == 0 {
			numTimesEndingAt0++
		}
	}

	// part1 := strconv.Itoa(numTimesEndingAt0)
	part2 := strconv.Itoa(numTimesTotalAt0)

	return part2
}

func calculatePosition(currentPosition int, direction string, movement int) (position int) {
	position = currentPosition
	for range movement {
		if direction == "R" {
			position++
			if position == 100 {
				position = 0
			}
		} else {
			position--
			if position == -1 {
				position = 99
			}
		}
		if position == 0 {
			numTimesTotalAt0++
		}
	}
	return
}

func init() {
	Register(1, SolveDay01)
}
