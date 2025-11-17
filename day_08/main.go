package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/ayo-awe/advent-of-code-2023/aoc"
)

var (
	networkPattern = regexp.MustCompile(`[0-9a-zA-Z]{3}`)
	dirIndexes     = map[rune]int{
		'L': 0,
		'R': 1,
	}
)

func ParseInput(lines []string) (string, map[string][2]string, error) {
	// line 1 contains the directions
	directions := lines[0]

	// network starts from line 3
	network := make(map[string][2]string)
	for _, line := range lines[2:] {
		matches := networkPattern.FindAllString(line, -1)
		if len(matches) != 3 {
			return "", nil, fmt.Errorf("invalid input")
		}

		node := matches[0]
		network[node] = [2]string(matches[1:])
	}

	return directions, network, nil
}

func main() {
	filename := flag.String("file", "input.txt", "input file name")
	flag.Parse()

	lines, err := aoc.ReadInputLineByLine(*filename)
	if err != nil {
		log.Fatal(err)
	}

	directions, network, err := ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(directions, network))
	fmt.Println("solution to part two: ", PartTwo(directions, network))
}

func PartOne(directions string, network map[string][2]string) int {
	steps := 0
	start := "AAA"
	end := "ZZZ"

	if _, ok := network[start]; !ok {
		return 0
	}

	node := start
	for ; node != end; steps++ {
		dir := directions[steps%len(directions)]

		// map L/R to indexes in the array
		dirIndex := dirIndexes[rune(dir)]
		node = network[node][dirIndex]
	}

	return steps
}

func PartTwo(directions string, network map[string][2]string) int {

	var nodes []string
	for k := range network {
		if strings.HasSuffix(k, "A") {
			nodes = append(nodes, k)
		}
	}

	nodeSteps := make([]int, len(nodes))
	for i, node := range nodes {
		steps := 0
		for current := node; !strings.HasSuffix(current, "Z"); steps++ {
			dir := rune(directions[steps%len(directions)])
			idx := dirIndexes[dir]
			current = network[current][idx]
		}
		nodeSteps[i] = steps
	}

	// solution is the lcm of the steps of all nodes
	lcmVal := 1
	for _, steps := range nodeSteps {
		lcmVal = lcm(lcmVal, steps)
	}

	return lcmVal
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// largest number that divides both a & b
func gcd(a, b int) int {
	for i := min(a, b); i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}

	return 1
}
