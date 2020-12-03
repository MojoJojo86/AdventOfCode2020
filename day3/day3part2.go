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

	localMapWidth := len(mapGrid[0])

	routes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	treesPerRoute := []int{}

	for _, route := range routes {

		posX, posY := 0, 0
		treeCount := 0
		rightMove := route[0]
		downMove := route[1]

		for posY < len(mapGrid) {
			currentValue := mapGrid[posY][posX]
			if currentValue == '#' {
				treeCount++
			}

			posX = (posX + rightMove) % localMapWidth
			posY += downMove
		}
		treesPerRoute = append(treesPerRoute, treeCount)
	}

	result := 1
	for _, value := range treesPerRoute {
		result *= value
	}

	fmt.Println(result)
}
