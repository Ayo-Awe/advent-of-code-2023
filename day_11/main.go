package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ayo-awe/advent-of-code-2023/aoc"
)

const (
	X, Y = 0, 1
)

func ParseInput(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for y := range lines {
		grid[y] = []rune(lines[y])
	}
	return grid
}

func main() {
	filename := flag.String("file", "input.txt", "input file name")
	flag.Parse()

	lines, err := aoc.ReadInputLineByLine(*filename)
	if err != nil {
		log.Fatal(err)
	}

	grid := ParseInput(lines)

	fmt.Println("solution to part one: ", PartOne(grid))
	fmt.Println("solution to part two: ", PartTwo(grid))
}

func PartOne(grid [][]rune) int {
	rExp, cExp := expansions(grid, 2)

	// find galaxies
	var galaxies [][2]int
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '#' {
				galaxies = append(galaxies, [2]int{c, r})
			}
		}
	}

	// permutate galaxies
	var sum int
	for i := range len(galaxies) {
		for j := i + 1; j < len(galaxies); j++ {
			galA, galB := galaxies[i], galaxies[j]
			sum += steps(galA, galB, rExp, cExp)
		}
	}

	return sum
}

// returns a map of non-expanded rows and columns
func expansions(grid [][]rune, magnitude int) (r, c []int) {
	// stores a map of expanded columns and rows
	r = make([]int, len(grid))
	c = make([]int, len(grid[0]))

	for y := range grid {
		// prefill row with magnitude
		r[y] = magnitude

		for x, cell := range grid[y] {
			// prefill columns with magnitude
			if c[x] == 0 {
				c[x] = magnitude
			}

			if cell == '#' {
				// mark row and column as non-expanded
				r[y] = 1
				c[x] = 1
			}
		}
	}

	return r, c
}

func PartTwo(grid [][]rune) int {
	rExp, cExp := expansions(grid, 1_000_000)

	// find galaxies
	var galaxies [][2]int
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '#' {
				galaxies = append(galaxies, [2]int{c, r})
			}
		}
	}

	// permutate galaxies
	var sum int
	for i := range len(galaxies) {
		for j := i + 1; j < len(galaxies); j++ {
			galA, galB := galaxies[i], galaxies[j]
			sum += steps(galA, galB, rExp, cExp)
		}
	}

	return sum
}

func steps(galA, galB [2]int, rExp, cExp []int) int {
	dx := sum(galA[X], galB[X], cExp)
	dy := sum(galA[Y], galB[Y], rExp)
	return dx + dy
}

// sum values between a and b
// exlusive of the lower index and inclusive of the higher index
// a or b could be the higher/lower index doesn't matter
func sum(a, b int, s []int) int {
	if a < 0 || a >= len(s) || b < 0 || b >= len(s) {
		return 0
	}

	start, end := min(a, b), max(a, b)

	var sum int
	for i := start + 1; i <= end; i++ {
		sum += s[i]
	}

	return sum
}
