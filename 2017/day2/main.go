package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

// readFile returns every row a csv separated by tabs as slice in slice
func readFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	return reader.ReadAll()
}

func sToI(s []string) (il []int) {
	for i := range s {
		n, err := strconv.Atoi(s[i])
		if err != nil {
			log.Println(err.Error())
		}
		il = append(il, n)
	}
	return
}

// rowDiff returns the difference between the smallest and the largest number in each row.
func rowDiff(row []string) int {
	r := sToI(row)
	var smallest = r[0]
	var largest int
	for _, d := range r {
		if d < smallest {
			smallest = d
		}
		if d > largest {
			largest = d
		}
	}
	return (largest - smallest)
}

func day1() (checksum int) {
	f, err := readFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	for i := range f {
		r := rowDiff(f[i])
		checksum += r
	}
	return
}
