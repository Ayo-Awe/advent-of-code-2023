package main

import (
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/ayo-awe/advent-of-code-2023/aoc"
)

var (
	cards = map[rune][2]int{
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
	fmt.Println("solution to part two: ", PartTwo())
}

type node struct {
	pos   [2]int
	dir   rune
	steps int
}

func PartOne(grid [][]rune, start [2]int) int {
	var queue []node

	for dir := range cards {
		queue = append(queue, node{
			pos:   start,
			dir:   dir,
			steps: 0,
		})
	}

	seen := map[[2]int]int{}
	max := math.MinInt

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if prevSteps := seen[curr.pos]; prevSteps > 0 && prevSteps < curr.steps {
			continue
		}

		seen[curr.pos] = curr.steps
		if curr.steps > max {
			max = curr.steps
		}

		d := cards[curr.dir]
		nx, ny := curr.pos[X]+d[X], curr.pos[Y]+d[Y]

		// bounds check
		if nx < 0 || nx >= len(grid[0]) || ny < 0 || ny >= len(grid) {
			continue
		}

		newDir := advance(curr.dir, grid[ny][nx])

		// not possible to advance
		if newDir == InvalidDir {
			continue
		}

		// add new node to the queue
		queue = append(queue, node{
			pos:   [2]int{nx, ny},
			dir:   newDir,
			steps: curr.steps + 1,
		})
	}

	return max
}

func PartTwo() int {
	return 0
}
