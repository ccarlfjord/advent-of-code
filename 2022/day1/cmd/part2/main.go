package main

import (
	"fmt"
	"os"

	"github.com/ccarlfjord/advent-of-code/2022/day1"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	sum := day1.Parse(file)
	topThree := day1.ThreeLargest(sum)
	topThreeSum := day1.Sum(topThree)

	fmt.Println(topThreeSum)
}
