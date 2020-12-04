package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	data, _ := ioutil.ReadFile("day4/day4input.txt")
	correctPassportCount := 0

out:
	for _, line := range strings.Split(string(data), "\n\n") {
		for _, field := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
			if ok := strings.Contains(line, field); !ok {
				continue out
			}
		}
		correctPassportCount++
	}

	fmt.Println(correctPassportCount)
}
