package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(filename string) (st string) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	st = strings.TrimSuffix(string(f), "\n")
	return
}

func part1(input string) {
	var res int
	for i, r := range input {
		if input[i] == input[(i+1)%len(input)] {
			res += int(r - '0')
		}
	}
	fmt.Println(res)
}

func part2(input string) {
	var res int
	for i, r := range input {
		if input[i] == input[(i+(len(input)/2))%len(input)] {
			res += int(r - '0')
		}
	}
	fmt.Println(res)
}

func main() {
	input := readFile("input.txt")
	part1(input)
	part2(input)
}
