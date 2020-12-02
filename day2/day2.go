package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := os.Open("day2input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)

	var stage int = 0
	var letterChoice string
	var minRange int
	var maxRange int
	var valid int

	for scanner.Scan() {
		if stage == 0 {
			// Figure out the ranges in this bit
			var countRange []string = strings.Split(scanner.Text(), "-")
			minRange, _ = strconv.Atoi(countRange[0])
			maxRange, _ = strconv.Atoi(countRange[1])
			stage++
		} else if stage == 1 {
			// Pull out the letter needed for password here
			var values []string = strings.Split(scanner.Text(), "")
			letterChoice = values[0]
			// fmt.Println(letterChoice)
			stage++
		} else if stage == 2 {
			// Check the password here matches the range and letter requirements.
			var password []string = strings.Split(scanner.Text(), "")
			if minRange < len(password) || maxRange < len(password) {
				if password[minRange-1] == letterChoice || password[maxRange-1] == letterChoice {
					if password[minRange-1] != password[maxRange-1] {
						valid++
					}
				}
			}
			stage = 0
		}
	}
	fmt.Println(valid)
}
