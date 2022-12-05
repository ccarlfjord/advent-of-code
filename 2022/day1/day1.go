package day1

import (
	"bufio"
	"io"
	"strconv"
)

// Parse returns a slice of summarized values separated as a new line from input
func Parse(r io.Reader) []int {
	scanner := bufio.NewScanner(r)
	var sum int
	var sumSlice []int
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			sumSlice = append(sumSlice, sum)
			sum = 0
		}
		if text != "" {
			curr, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			sum += curr
		}
	}
	for scanner.Scan() {
		if scanner.Text() == "" {
			sumSlice = append(sumSlice, sum)
			sum = 0
		} else {
			snack, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			sum += snack
		}
	}
	return sumSlice
}

// largest returns the largest int in a slice and its postition
func Largest(s []int) (n int, pos int) {
	for i := range s {
		j := i + 1
		if j >= len(s) {
			return n, pos
		}
		if s[i] > s[j] && s[i] > n {
			n = s[i]
			pos = i
		}
	}
	return n, pos
}

// largestThree returns a slice of the three largest ints from s.
func ThreeLargest(s []int) []int {
	var largestThree []int
	for len(largestThree) < 3 {
		// Get largest value
		n, pos := Largest(s)
		// Add largest value to new slice
		largestThree = append(largestThree, n)
		// Reassign slice without largest value
		s = append(s[:pos], s[pos+1:]...)
	}
	return largestThree
}

func Sum(s []int) int {
	var sum int
	for _, i := range s {
		sum += i
	}
	return sum
}
