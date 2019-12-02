package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputScanner() *bufio.Scanner {
	f, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(f)
	return scanner

}
func p1() {

}
func main() {
	fmt.Println(p1())
}
