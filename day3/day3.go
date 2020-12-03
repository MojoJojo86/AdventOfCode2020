package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, err := os.Open("day3/day3input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	var mapGrid []string

	for scanner.Scan() {
		mapGrid = append(mapGrid, scanner.Text())
	}

	rightMove := 3
	downMove := 1
	posX, posY := 0, 0
	treeCount := 0
	localMapWidth := len(mapGrid[0])

	for posY < len(mapGrid) {
		currentValue := mapGrid[posY][posX]
		if currentValue == '#' {
			treeCount++
		}

		posX = (posX + rightMove) % localMapWidth
		posY += downMove
	}

	fmt.Println(treeCount)
}
