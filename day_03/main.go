package main

import (
	"fmt"
	"log"

	"github.com/ayo-awe/advent-of-code-2023/util"
)

type symbol struct {
	value rune
	pos   [2]int
}

type part struct {
	value   int
	symbols []symbol
}

func getAdjacentSymbols(x, y int, grid [][]rune) []symbol {
	var symbols []symbol
	for nx := x - 1; nx <= x+1; nx++ {
		for ny := y - 1; ny <= y+1; ny++ {
			if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(grid[0]) {
				continue
			}

			value := grid[ny][nx]

			// skip main cell
			if nx == x && ny == y {
				continue
			}

			// x,y out of bounds

			isSymbol := !isDigit(value) && value != '.'
			if isSymbol {
				symbols = append(symbols, symbol{
					value: value,
					pos:   [2]int{nx, ny},
				})
			}
		}
	}

	return symbols
}

func parseInput(lines []string) [][]rune {
	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}

func main() {
	lines, err := util.ReadInputLineByLine("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := parseInput(lines)
	parts := partsWithSymbols(grid)
	fmt.Println(len(parts))
}

func partsWithSymbols(grid [][]rune) []part {
	var parts []part

	inNumber := false
	partDigits := ""
	partSymbols := map[symbol]struct{}{}

	for y, row := range grid {
		for x, cell := range row {
			if isDigit(cell) {
				inNumber = true
				partDigits += string(cell)

				adjSymbols := getAdjacentSymbols(x, y, grid)
				for _, symbol := range adjSymbols {
					partSymbols[symbol] = struct{}{}
				}
			} else if (!isDigit(cell) || isLastCell(x, y, grid)) && inNumber {
				partValue := util.MustToInt(partDigits)

				// this number isn't a part because it has no symbols
				if len(partSymbols) == 0 {
					continue
				}

				symbols := make([]symbol, 0, len(partSymbols))
				for s := range partSymbols {
					symbols = append(symbols, s)
				}

				parts = append(parts, part{value: partValue, symbols: symbols})

				// reset
				inNumber = false
				partSymbols = map[symbol]struct{}{}
				partDigits = ""
			}
		}
	}

	if inNumber && len(partSymbols) == 0 {
		partValue := util.MustToInt(partDigits)

		symbols := make([]symbol, 0, len(partSymbols))
		for s := range partSymbols {
			symbols = append(symbols, s)
		}

		parts = append(parts, part{value: partValue, symbols: symbols})
	}

	return parts
}

func isLastCell(x, y int, grid [][]rune) bool {
	return y == len(grid)-1 && x == len(grid[0])-1
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}
