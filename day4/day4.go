package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.Open("day4/day4input.txt")
	//data, err := os.Open("day4/demo.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	completePasswordCount := 0
	byr := false
	iyr := false
	eyr := false
	hgt := false
	hcl := false
	ecl := false
	pid := false
	for scanner.Scan() {
		if scanner.Text() != "" {
			if strings.Contains(scanner.Text(), "byr:") {
				byr = true
			}
			if strings.Contains(scanner.Text(), "iyr:") {
				iyr = true
			}
			if strings.Contains(scanner.Text(), "eyr:") {
				eyr = true
			}
			if strings.Contains(scanner.Text(), "hgt:") {
				hgt = true
			}
			if strings.Contains(scanner.Text(), "hcl:") {
				hcl = true
			}
			if strings.Contains(scanner.Text(), "ecl:") {
				ecl = true
			}
			if strings.Contains(scanner.Text(), "pid:") {
				pid = true
			}
			if byr && iyr && eyr && hgt && hcl && ecl && pid == true {
				completePasswordCount++
			}

		} else {
			byr = false
			iyr = false
			eyr = false
			hgt = false
			hcl = false
			ecl = false
			pid = false
		}
	}
	fmt.Println(completePasswordCount)
}
