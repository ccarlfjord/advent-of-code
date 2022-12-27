package main

import (
	"fmt"
	"os"

	"github.com/ccarlfjord/advent-of-code/2022/day2"
)

var rules = map[string]map[string]int{
	"X": {"A": day2.Draw, "B": day2.Loss, "C": day2.Win},
	"Y": {"A": day2.Win, "B": day2.Draw, "C": day2.Loss},
	"Z": {"A": day2.Loss, "B": day2.Win, "C": day2.Draw},
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
	hand := day2.Points[player]
	outcome := rules[player][opponent]
	return hand + outcome
}
