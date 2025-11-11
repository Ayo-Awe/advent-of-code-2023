package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strings"

	"github.com/ayo-awe/advent-of-code-2023/util"
)

func countMatches(input string) int {
	nums := strings.Split(strings.Split(input, ": ")[1], " | ")

	myNums := strings.Split(nums[0], " ")
	winningNums := strings.Split(nums[1], " ")

	matches := 0
	for _, num := range myNums {
		if slices.Contains(winningNums, num) && num != "" {
			matches++
		}
	}

	return matches
}

func main() {
	lines, err := util.GetInput()
	if err != nil {
		log.Fatal(err)
	}

	numberOfCards := len(lines)
	sumOfCardPoints := 0
	cardCopies := make([]int, numberOfCards)
	totalCardInstances := numberOfCards

	for idx, line := range lines {

		matches := countMatches(line)

		// distribute cardInstances
		cardInstances := cardCopies[idx] + 1

		for i := idx + 1; i < idx+1+matches; i++ {
			// index out of range
			if i > len(cardCopies)-1 {
				break
			}

			cardCopies[i] += cardInstances
			totalCardInstances += cardInstances
		}

		cardPoints := 0

		if matches > 0 {
			cardPoints = int(math.Pow(2, float64(matches-1)))
		}

		sumOfCardPoints += cardPoints
	}

	fmt.Printf("Sum of scratchcard points: %v\n", sumOfCardPoints)
	fmt.Printf("Total card instances: %v\n", totalCardInstances)
}
