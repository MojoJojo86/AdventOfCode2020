package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	data, err := os.Open("day1input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	// Part 1 Solution
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			firstNum, _ := strconv.Atoi(string(values[i]))
			secondNum, _ := strconv.Atoi(string(values[j]))
			if firstNum+secondNum == 2020 {
				fmt.Printf("first: %d second: %d, overall: %d", firstNum, secondNum, firstNum*secondNum)
			}
		}
	}

	// Part 2 Solution 
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			for k := j + 1; k < len(values); k++ {
				firstNum, _ := strconv.Atoi(string(values[i]))
				secondNum, _ := strconv.Atoi(string(values[j]))
				thirdNum, _ := strconv.Atoi(string(values[k]))
				if firstNum+secondNum+thirdNum == 2020 {
					fmt.Printf("first: %d second: %d, third: %d, overall: %d", firstNum, secondNum, thirdNum, firstNum*secondNum*thirdNum)
				}
			}
		}
	}
}
