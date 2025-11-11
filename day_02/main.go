package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id   int
	sets []*Set
}

type Set struct {
	blue  int
	green int
	red   int
}

func NewGame(id int, sets []*Set) *Game {
	return &Game{id, sets}
}

func (g *Game) isValid(red, blue, green int) bool {
	for _, set := range g.sets {

		if set.red > red || set.blue > blue || set.green > green {
			return false
		}
	}

	return true
}

func SetFromString(str string) *Set {
	str = strings.Trim(str, " ")
	cubes := strings.Split(str, ", ")
	set := Set{}

	for _, cube := range cubes {
		d := strings.Split(cube, " ")

		count, err := strconv.Atoi(d[0])
		if err != nil {
			panic("Something went wrong with the parser")
		}

		color := d[1]

		if color == "red" {
			set.red = count
		} else if color == "green" {
			set.green = count
		} else {
			set.blue = count
		}
	}

	return &set
}

func (g *Game) MinimumSet() *Set {
	red := 0
	blue := 0
	green := 0

	for _, set := range g.sets {
		if set.red > red {
			red = set.red
		}

		if set.green > green {
			green = set.green
		}

		if set.blue > blue {
			blue = set.blue
		}
	}

	return &Set{red, blue, green}
}

func (s *Set) Power() int {
	return s.blue * s.green * s.red
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go input_file.txt")
	}

	filePath := os.Args[1]

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")

	// for each line
	// split by : and trim
	// first element contains the game and second element contains the set data
	// split the second element by ; and trim
	// every element is a set
	// for every set
	// split by , and trim
	// each element represents the number of cubes of a specific color drawn in that set
	// the first index of the string reprensets the number
	var games []*Game

	for _, line := range lines {
		if line == "" {
			continue
		}

		data := strings.Split(line, ":")

		gameData := data[0]
		setData := data[1]

		rawSets := strings.Split(setData, ";")
		var sets []*Set

		for _, rawSet := range rawSets {
			set := SetFromString(rawSet)
			sets = append(sets, set)
		}

		gameNo, err := strconv.Atoi(strings.Replace(gameData, "Game ", "", -1))
		if err != nil {
			panic("Something went wrong with the parser")
		}

		game := NewGame(gameNo, sets)
		games = append(games, game)
	}

	sumOfIds := 0
	sumOfPowers := 0

	for _, game := range games {
		if game.isValid(12, 14, 13) {
			sumOfIds += game.id
		}

		minimumSet := game.MinimumSet()
		sumOfPowers += minimumSet.Power()
	}

	fmt.Printf("Sum of ids of valid games: %v\n", sumOfIds)
	fmt.Printf("Sum of powers of minimum set of each game: %v\n", sumOfPowers)
}
