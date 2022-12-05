package main

import (
	"fmt"
	"os"

	"github.com/ccarlfjord/advent-of-code/2022/day1"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	sum := day1.Parse(input)
	largest, _ := day1.Largest(sum)
	fmt.Println(largest)
}
