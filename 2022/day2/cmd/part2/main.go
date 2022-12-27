package main

import (
	"fmt"
	"os"

	"github.com/ccarlfjord/advent-of-code/2022/day2"
)

var rock string = "A"
var paper string = "B"
var scissor string = "C"

var rules = map[string]map[string]string{
	"X": {
		rock:    scissor,
		paper:   rock,
		scissor: paper,
	},
	"Y": {
		rock:    rock,
		paper:   paper,
		scissor: scissor,
	},
	"Z": {
		rock:    paper,
		paper:   scissor,
		scissor: rock,
	},
}
var pointMap = map[string]int{
	"X": day2.Loss,
	"Y": day2.Draw,
	"Z": day2.Win,
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	score := day2.Score(f, outcome)

	fmt.Println(score)

}

func outcome(player, opponent string) int {
	hand := rules[player][opponent]
	points := day2.Points[hand]
	outcome := pointMap[player]
	return points + outcome
}
