package day2

import (
	"fmt"
	"strings"
	"testing"
)

var testRounds = `A Y
B X
C Z
`

func TestPart1(t *testing.T) {
	r := strings.NewReader(testRounds)

	score := Score(r)
	fmt.Println(score)

}

func TestPart2(t *testing.T) {

}
