package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ayo-awe/advent-of-code-2023/util"
)

type Number struct {
	value      int
	neighbours []string
}

func (n *Number) IsPartNumber() bool {
	// A number is a part number if one of it's neighbours is a symbol
	for _, neigbour := range n.neighbours {
		if !util.IsDigit(neigbour) && neigbour != "." {
			return true
		}
	}

	return false
}

func getNeighbourCells(i, j, height, width int, grid []string) []string {
	var neighbours []string

	for a := i - 1; a <= i+1; a++ {
		// Indexes out of bounds
		if a < 0 || a > height-1 {
			continue
		}

		for b := j - 1; b <= j+1; b++ {
			// Indexes out of bounds
			if b < 0 || b > width-1 {
				continue
			}

			// Ignore cell under question
			if a == i && b == j {
				continue
			}

			neighbours = append(neighbours, string(grid[a][b]))
		}
	}

	return neighbours
}

func main() {
	lines, err := util.GetInput()
	if err != nil {
		log.Fatal(err)
	}

	//  Grid size
	height := len(lines)
	width := len(lines[0])

	var numbers []Number

	// Treat input as an m * n grid
	for i, line := range lines {
		inNumber := false
		numberString := ""
		var neighbours []string

		for j, char := range line {
			if util.IsDigit(string(char)) && !inNumber {
				numberString += string(char)
				neighbourCells := getNeighbourCells(i, j, height, width, lines)
				neighbours = append(neighbours, neighbourCells...)
				inNumber = true
			} else if util.IsDigit(string(char)) && inNumber {
				numberString += string(char)
				neighbourCells := getNeighbourCells(i, j, height, width, lines)
				neighbours = append(neighbours, neighbourCells...)
			} else if !util.IsDigit(string(char)) && inNumber {
				inNumber = false
				value, err := strconv.Atoi(numberString)
				if err != nil {
					log.Panic("Something is wrong with the parser")
				}

				number := Number{value, neighbours}
				numbers = append(numbers, number)

				numberString = ""
				neighbours = nil
			}

		}

		if inNumber {
			value, err := strconv.Atoi(numberString)
			if err != nil {
				log.Panic("Something is wrong with the parser")
			}

			number := Number{value, neighbours}
			numbers = append(numbers, number)
		}

	}

	partNumberSum := 0
	for _, number := range numbers {
		if number.IsPartNumber() {
			partNumberSum += number.value
		}

	}

	fmt.Printf("Part number sum is: %v\n", partNumberSum)
}
