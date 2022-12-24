package day2

import (
	"bufio"
	"io"
	"strings"
)

var Points = map[string]int{
	// Rock
	"A": 1,
	// Paper
	"B": 2,
	// Scissor
	"C": 3,

	// Rock
	"X": 1,
	// Paper
	"Y": 2,
	// Scissor
	"Z": 3,
}

const (
	loss = 0
	draw = 3
	win  = 6
)

func Score(r io.Reader) int {
	var sum int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")
		opponent := split[0]
		player := split[1]
		outcome := outcome(player, opponent)
		sum += outcome
	}
	return sum
}

func outcome(player, opponent string) int {
	point := Points[player]

	rules := map[string]map[string]int{
		"X": {"A": draw, "B": loss, "C": win},
		"Y": {"A": win, "B": draw, "C": loss},
		"Z": {"A": loss, "B": win, "C": draw},
	}
	return point + rules[player][opponent]

}
