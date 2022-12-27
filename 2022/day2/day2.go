package day2

import (
	"bufio"
	"io"
	"strings"
)

const (
	Win  = 6
	Draw = 3
	Loss = 0
)

var Points = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,

	"X": 1,
	"Y": 2,
	"Z": 3,
}

func Score(r io.Reader, outcomeFunc func(player, opponent string) int) int {
	var sum int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")
		opponent := split[0]
		player := split[1]
		outcome := outcomeFunc(player, opponent)
		sum += outcome
	}
	return sum
}
