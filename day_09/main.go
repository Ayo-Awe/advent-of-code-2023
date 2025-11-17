package main

import (
	"flag"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/ayo-awe/advent-of-code-2023/aoc"
)

func ParseInput(lines []string) ([][]int, error) {
	histories := make([][]int, len(lines))
	for i, line := range lines {
		historyStr := strings.Split(line, " ")
		history := make([]int, len(historyStr))
		for j, val := range historyStr {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			history[j] = intVal
		}
		histories[i] = history
	}
	return histories, nil
}

func main() {
	filename := flag.String("file", "input.txt", "input file name")
	flag.Parse()

	lines, err := aoc.ReadInputLineByLine(*filename)
	if err != nil {
		log.Fatal(err)
	}

	histories, err := ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("solution to part one: ", PartOne(histories))
	fmt.Println("solution to part two: ", PartTwo())
}

func PartOne(histories [][]int) int {
	var sum int
	for i := range histories {
		sum += extrapolate(histories[i])
	}

	return sum
}

func extrapolate(sequences []int) int {
	// if everything is zero return 0
	if slices.Max(sequences) == 0 && slices.Min(sequences) == 0 {
		return 0
	}

	// compute next sequence level and call extrapolate on that
	nextSeqs := make([]int, len(sequences)-1)
	for i := range nextSeqs {
		nextSeqs[i] = sequences[i+1] - sequences[i]
	}

	return extrapolate(nextSeqs) + sequences[len(sequences)-1]
}

func PartTwo() int {
	return 0
}
