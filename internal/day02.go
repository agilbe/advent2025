package internal

import (
	"strconv"
	"strings"
)

var idsWithDupes = []int{}

func SolveDay2(input string) string {
	idRanges := strings.Split(input, ",")
	for _, rangeString := range idRanges {
		start, _ := strconv.Atoi(strings.Split(rangeString, "-")[0])
		end, _ := strconv.Atoi(strings.Split(rangeString, "-")[1])
		checkForDupes(start, end)
	}
	sumOfDupes := 0
	for _, id := range idsWithDupes {
		sumOfDupes += id
	}
	return strconv.Itoa(sumOfDupes)
}

func checkForDupes(start int, end int) {
	for i := start; i <= end; i++ {
		// if hasDupesPart1(i) {
		// 	idsWithDupes = append(idsWithDupes, i)
		// }
		if hasDupesPart2(i) {
			idsWithDupes = append(idsWithDupes, i)
		}
	}
}

func hasDupesPart1(id int) bool {
	idRunes := []rune(strconv.Itoa(id))

	// only has a dupe if a string is repeated twice - it always has an even number of digits
	if (len(idRunes) % 2) != 0 {
		return false
	}
	// iterate over half the string, starting from 0 and also at half of its length
	halfwayPoint := len(idRunes) / 2
	for i := range halfwayPoint {
		if idRunes[i] != idRunes[i+halfwayPoint] {
			return false
		}
	}

	return true
}

func hasDupesPart2(id int) bool {
	return false
}

func init() {
	Register(2, SolveDay2)
}
