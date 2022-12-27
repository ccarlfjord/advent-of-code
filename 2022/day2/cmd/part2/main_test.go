package main

import (
	"strings"
	"testing"

	"github.com/ccarlfjord/advent-of-code/2022/day2"
)

var testRounds = `A Y
B X
C Z
`

func TestMain(t *testing.T) {
	r := strings.NewReader(testRounds)
	score := day2.Score(r, outcome)
	t.Log(score)
}
