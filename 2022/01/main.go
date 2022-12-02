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
	largest := getLargest(sum)
	fmt.Println(largest)
}

// summarize summariezes numbers separated as a new line
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

func getLargest(s []int) int {
	var largest int
	for i := range s {
		j := i + 1
		if j >= len(s) {
			return largest
		}
		if s[i] > s[j] && s[i] > largest {
			largest = s[i]
		}
	}
	return largest
}
