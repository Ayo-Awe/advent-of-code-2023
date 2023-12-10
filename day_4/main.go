package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strings"

	"github.com/ayo-awe/advent-of-code-2023/util"
)

func main() {
	lines, err := util.GetInput()
	if err != nil {
		log.Fatal(err)
	}

	sumOfCardPoints := 0

	for _, line := range lines {

		nums := strings.Split(strings.Split(line, ": ")[1], " | ")

		myNums := strings.Split(nums[0], " ")
		winningNums := strings.Split(nums[1], " ")

		matches := 0
		for _, num := range myNums {
			if slices.Contains(winningNums, num) && num != "" {
				matches++
			}
		}

		cardPoints := 0

		if matches > 0 {
			cardPoints = int(math.Pow(2, float64(matches-1)))
		}

		sumOfCardPoints += cardPoints
	}

	fmt.Printf("Sum of scratchcard points: %v\n", sumOfCardPoints)
}
