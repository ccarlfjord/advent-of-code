package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func massCalc(i int) int {
	mass := math.Floor(float64((i / 3) - 2))
	if mass > 0 {
		return int(mass)
	}
	return 0
}

func part1() int {
	/*
		For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
		For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
		For a mass of 1969, the fuel required is 654.
		For a mass of 100756, the fuel required is 33583.
	*/
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		fmt.Printf("%v", err)
	}
	var sum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("%v", err)
		}
		mass := massCalc(lineInt)
		sum += mass
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func part2() int {
	/*
		A module of mass 14 requires 2 fuel.
		This fuel requires no further fuel (2 divided by 3 and rounded down is 0, which would call for a negative fuel), so the total fuel required is still just 2.
		At first, a module of mass 1969 requires 654 fuel. Then, this fuel requires 216 more fuel (654 / 3 - 2).
		216 then requires 70 more fuel, which requires 21 fuel, which requires 5 fuel, which requires no further fuel. So, the total fuel required for a module of mass 1969 is 654 + 216 + 70 + 21 + 5 = 966.
		The fuel required by a module of mass 100756 and its fuel is: 33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346.

	*/
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		fmt.Printf("%v", err)
	}
	var sum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("%v", err)
		}
		fuel := massCalc(lineInt)
		for fuel > 0 {
			sum += fuel
			fuel = massCalc(fuel)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func main() {
	p1 := part1()
	fmt.Println(p1)
	p2 := part2()
	fmt.Println(p2)
}
