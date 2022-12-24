package main

import (
	"fmt"
	"os"

	"github.com/ccarlfjord/advent-of-code/2022/day2"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	score := day2.Score(f)

	fmt.Println(score)

}
