package internal

import (
	"fmt"
	"strings"
)

func SolveDay01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return fmt.Sprintf("Number of lines: %d", len(lines))
}

func init() {
	Register(1, SolveDay01)
}
