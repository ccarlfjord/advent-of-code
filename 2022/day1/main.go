package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	sum := summarize(file)
	largest, _ := largest(sum)
	fmt.Println(largest)
}

// summarize summarizes numbers separated as a new line
// empty lines are treated as a new sum
func summarize(r io.Reader) []int {
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

// largest returns the largest number in a slice of ints and its postition
func largest(s []int) (n int, pos int) {
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

// largestThree returns a sum of the three largest numbers in a slice.
func largestThree(s []int) []int {
	fmt.Println(len(s))
	var largestThree []int
	for len(largestThree) < 3 {
		n, pos := largest(s)
		largestThree = append(largestThree, n)
		s = append(s[:pos], s[pos+1:]...)
		fmt.Println(len(s))
	}
	return largestThree
}
