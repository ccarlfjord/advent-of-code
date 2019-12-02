package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calc(res int, s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	res += i
	return res
}

func inputScanner() *bufio.Scanner {
	f, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(f)
	return scanner

}
func p1() (res int) {
	scanner := inputScanner()
	for scanner.Scan() {
		res = calc(res, scanner.Text())
	}
	return
}

func p2() (hit int) {
	res := 0
	m := make(map[int]bool)
	for {
		scanner := inputScanner()
		for scanner.Scan() {
			res = calc(res, scanner.Text())
			if m[res] == true {
				hit = res
				return
			}
			m[res] = true
		}
	}
}
func main() {
	fmt.Println(p1())
	fmt.Println(p2())
}
