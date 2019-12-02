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

func totalMassCalc(i int) (sum int) {
	mass := massCalc(i)
	for mass > 0 {
		sum += mass
		mass = massCalc(mass)
	}
	return sum
}

func main() {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		fmt.Printf("%v", err)
	}
	var fuelNeeded int
	var totalFuelNeeded int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("%v", err)
		}
		fuelNeeded += massCalc(lineInt)
		totalFuelNeeded += totalMassCalc(lineInt)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fuelNeeded)
	fmt.Println(totalFuelNeeded)
}
