package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/ayo-awe/advent-of-code-2023/util"
)

type Race struct {
	time     int
	distance int
}

func NewRace(time, distance int) *Race {
	return &Race{time: time, distance: distance}
}

func (r *Race) WaysToWin() int {
	var minWinSpeed int

	for i := 1; i < r.time; i++ {
		speed := i
		time := r.time - speed

		dist := speed * time

		// first speed to win a race is the minimum speed required to win
		if dist > r.distance {
			minWinSpeed = i
			break
		}
	}

	return r.time - 2*minWinSpeed + 1
}

func main() {
	lines, err := util.GetInput()
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\d+`)
	times := re.FindAllString(lines[0], -1)
	distances := re.FindAllString(lines[1], -1)

	// Prep puzzle one input
	var races []Race
	for i := 0; i < len(times); i++ {
		time := util.MustToInt(times[i])
		distance := util.MustToInt(distances[i])
		races = append(races, Race{time, distance})
	}

	puzzleOne(races)

	// Prep puzzle two input
	time := strings.Join(times, "")
	distance := strings.Join(distances, "")
	race := &Race{util.MustToInt(time), util.MustToInt(distance)}

	puzzleTwo(race)
}

func puzzleOne(races []Race) {
	marginOfError := 1

	for _, race := range races {
		nWins := race.WaysToWin()
		marginOfError *= nWins
	}

	fmt.Println("Puzzle one solution:", marginOfError)
}

func puzzleTwo(r *Race) {
	nWins := r.WaysToWin()

	fmt.Println("Puzzle two solution:", nWins)
}
