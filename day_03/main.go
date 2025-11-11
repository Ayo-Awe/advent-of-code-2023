package main

import (
	"fmt"
	"log"
	"maps"
	"strconv"

	"github.com/ayo-awe/advent-of-code-2023/aoc"
)

const (
	X = 0
	Y = 1
)

func ParseInput(lines []string) [][]rune {
	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}

func main() {
	lines, err := aoc.ReadInputLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := ParseInput(lines)
	parts, symbols := GetPartsAndSymbols(grid)
	fmt.Println("Part One: ", PartOne(parts))
	fmt.Println("Part Two: ", PartTwo(grid, symbols))
}

func PartOne(parts []int) int {
	sum := 0
	for _, part := range parts {
		sum += part
	}
	return sum
}

func PartTwo(grid [][]rune, symbols map[[2]int][]int) int {
	gearRatio := 0

	for pos, nums := range symbols {
		symbol := grid[pos[Y]][pos[X]]
		if symbol != '*' {
			continue
		}

		if len(nums) != 2 {
			continue
		}

		gearRatio += nums[0] * nums[1]
	}

	return gearRatio
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func adjacentSymbols(grid [][]rune, pos [2]int) map[[2]int]struct{} {
	symbols := make(map[[2]int]struct{})

	neighbourCells := [][2]int{
		{0, 1},   // Down
		{1, 1},   // Bottom Right
		{-1, 1},  // Bottom Left
		{0, -1},  // Up
		{1, -1},  // Top Right
		{-1, -1}, // Top Left
		{1, 0},   // Right
		{-1, 0},  // Left
	}

	for _, cell := range neighbourCells {
		x := cell[X] + pos[X]
		y := cell[Y] + pos[Y]

		// bounds check
		if x > len(grid[0])-1 || x < 0 || y > len(grid)-1 || y < 0 {
			continue
		}

		if grid[y][x] != '.' && !isDigit(grid[y][x]) {
			symbols[newPos(x, y)] = struct{}{}
		}
	}

	return symbols
}

func newPos(x, y int) [2]int {
	return [2]int{x, y}
}

func GetPartsAndSymbols(grid [][]rune) ([]int, map[[2]int][]int) {
	parts := make([]int, 0)
	symbols := make(map[[2]int][]int, 0)
	for y := range grid {
		var inNum bool
		var numStr []rune
		partSymbols := map[[2]int]struct{}{}

		for x := range grid[y] {
			cell := grid[y][x]

			if isDigit(cell) {
				inNum = true
				numStr = append(numStr, cell)
				maps.Copy(partSymbols, adjacentSymbols(grid, newPos(x, y)))
			}

			// Not a digit but isNum is true or at last digit on the row
			// we've reached the end of a number
			if (!isDigit(cell) || x == len(grid[y])-1) && inNum {
				if len(partSymbols) > 0 {
					num, err := strconv.Atoi(string(numStr))
					if err != nil {
						panic(err)
					}

					parts = append(parts, num)
					for symbol := range partSymbols {
						symbols[symbol] = append(symbols[symbol], num)
					}
				}

				inNum = false
				numStr = nil
				clear(partSymbols)
			}
		}
	}

	return parts, symbols
}
