package main

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	_, err := readFile("input.txt")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestRowCalc(t *testing.T) {
	f, err := readFile("input.txt")
	if err != nil {
		t.Error(err.Error())
	}
	for i := range f {
		r := rowDiff(f[i])
		t.Logf("%#v is the diff off the largest and smallest on loop run %v \n", r, i)
	}
}

func TestDay1(t *testing.T) {
	r := day1()
	t.Logf("%v is the sum of of the diffs", r)
}
