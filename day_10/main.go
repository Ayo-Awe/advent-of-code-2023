package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ayo-awe/advent-of-code-2023/aoc"
)

var (
	cardinals = map[rune][2]int{
		'^': {0, -1},
		'v': {0, 1},
		'>': {1, 0},
		'<': {-1, 0},
	}

	InvalidDir = '0'

	X = 0
	Y = 1
)

func advance(dir, cell rune) rune {
	if cell == 'S' {
		return dir
	}

	combo := string(dir) + string(cell)
	switch combo {
	case ">J":
		return '^'
	case "vJ":
		return '<'
	case "^F":
		return '>'
	case "<F":
		return 'v'
	case ">7":
		return 'v'
	case "^7":
		return '<'
	case "vL":
		return '>'
	case "<L":
		return '^'
	case "^|", "v|":
		return dir
	case ">-", "<-":
		return dir
	default:
		return InvalidDir
	}
}

func ParseInput(lines []string) ([][]rune, [2]int) {
	var start [2]int
	grid := make([][]rune, len(lines))

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == 'S' {
				start = [2]int{x, y}
			}
		}
		grid[y] = []rune(lines[y])
	}

	return grid, start
}

func main() {
	filename := flag.String("file", "input.txt", "input file name")
	flag.Parse()

	lines, err := aoc.ReadInputLineByLine(*filename)
	if err != nil {
		log.Fatal(err)
	}

	grid, start := ParseInput(lines)

	fmt.Println("solution to part one: ", PartOne(grid, start))
	fmt.Println("solution to part two: ", PartTwo(grid, start))
}

type node struct {
	pos [2]int
	dir rune
}

func (n node) String() string {
	return fmt.Sprintf("x: %d, y: %d, dir: %c", n.pos[X], n.pos[Y], n.dir)
}

func PartOne(grid [][]rune, start [2]int) int {
	loop := findLoop(grid, start)
	loopSize := len(loop)
	// further point is half the loop length rounded up
	return (loopSize / 2) + (loopSize % 2)
}

// 1: find all non-pipe cells
// 2: check if they're inside/outside the loop(polygon) using the raycast algo
func PartTwo(grid [][]rune, start [2]int) int {
	loop := findLoop(grid, start)
	var tilesInLoop int

	for y := range grid {
		for x := range grid[y] {
			pos := [2]int{x, y}
			if loop[pos] {
				continue
			}

			inTile := raycast(grid, loop, pos)
			if inTile {
				tilesInLoop++
			}
		}
	}

	return tilesInLoop
}

// given a point in the grid, returns true if
// the point is within the loop using the raycast algo
func raycast(grid [][]rune, loop map[[2]int]bool, pos [2]int) bool {
	corners := map[rune]bool{
		'J': true,
		'F': true,
		'L': true,
		'7': true,
		'S': true,
	}

	x, y := pos[X], pos[Y]

	// shoot ray to the right and count how many times
	// we cross the polygon/loop
	var crosses int
	var prevCorner rune
	for nx := x + 1; nx < len(grid[0]); nx++ {
		// skip this, point isn't part of the loop/polygon
		if !loop[[2]int{nx, y}] {
			continue
		}

		cell := grid[y][nx]
		if cell == '|' {
			crosses++
			continue
		}

		if !corners[cell] {
			continue
		}

		// no previous corner
		if prevCorner == 0 {
			prevCorner = cell
			continue
		}

		cornerPair := string(prevCorner) + string(cell)
		// we only count a cross if the corners are facing opposite directions
		if cornerPair == "FJ" || cornerPair == "L7" {
			crosses++
		}

		// reset previous corner
		prevCorner = 0
	}

	// num of crosses are odd for points in the loop/polygon and
	// even when the point is outside the loop
	return crosses%2 == 1
}

func findLoop(grid [][]rune, start [2]int) map[[2]int]bool {
	var curr node
	for dir, pos := range cardinals {
		nx, ny := pos[X]+start[X], pos[Y]+start[Y]

		// bounds check
		if nx < 0 || nx >= len(grid[0]) || ny < 0 || ny >= len(grid) {
			continue
		}

		if ndir := advance(dir, grid[ny][nx]); ndir != InvalidDir {
			curr = node{pos: start, dir: ndir}
			break
		}
	}

	seen := map[[2]int]bool{}
	for !seen[curr.pos] {
		d := cardinals[curr.dir]
		nx, ny := curr.pos[X]+d[X], curr.pos[Y]+d[Y]

		// bounds check
		if nx < 0 || nx >= len(grid[0]) || ny < 0 || ny >= len(grid) {
			// shouldn't happen with valid input
			panic("invalid loop")
		}

		newDir := advance(curr.dir, grid[ny][nx])
		if newDir == InvalidDir {
			// shouldn't happen with valid input
			panic("invalid loop")
		}

		seen[curr.pos] = true
		curr = node{pos: [2]int{nx, ny}, dir: newDir}
	}

	return seen
}
