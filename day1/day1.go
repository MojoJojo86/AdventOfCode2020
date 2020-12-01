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

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			firstNum, _ := strconv.Atoi(string(values[i]))
			secondNum, _ := strconv.Atoi(string(values[j]))
			if firstNum+secondNum == 2020 {
				fmt.Printf("first: %d second: %d, overall: %d", firstNum, secondNum, firstNum*secondNum)
			}
		}
	}
}
