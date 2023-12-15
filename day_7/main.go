package main

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strings"

	"github.com/ayo-awe/advent-of-code-2023/util"
)

type HandType int

// higher the index, stronger the card
const (
	HighCard     HandType = iota + 1 // index = 1
	OnePair                          // index = 2
	TwoPair                          // index = 3
	ThreeOfAKind                     // index = 4
	FullHouse                        // index = 5
	FourOfAKind                      // index = 6
	FiveOfAKind                      // index = 7
)

type Hand struct {
	cards    string
	bid      int
	handType HandType
}

func NewHand(cards string, bid int) *Hand {
	handType := GetHandType(cards)
	return &Hand{cards, bid, handType}
}

func GetHandType(cards string) HandType {
	cardDistribution := make(map[string]int)

	handTypeDistribution := map[HandType][]int{
		FiveOfAKind:  {5},
		FourOfAKind:  {1, 4},
		FullHouse:    {2, 3},
		ThreeOfAKind: {1, 1, 3},
		TwoPair:      {1, 2, 2},
		OnePair:      {1, 1, 1, 2},
		HighCard:     {1, 1, 1, 1, 1},
	}

	for _, card := range cards {
		cardDistribution[string(card)] += 1
	}

	var dist []int
	for _, value := range cardDistribution {
		dist = append(dist, value)
	}

	slices.Sort(dist)
	var handType HandType
	for key, value := range handTypeDistribution {
		if slices.Compare(dist, value) == 0 {
			handType = key
			break
		}
	}

	return handType
}

type Hands []*Hand

func (h Hands) Len() int {
	return len(h)
}

func (h Hands) Less(i, j int) bool {
	cardRanks := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}

	if h[i].handType < h[j].handType {
		return true
	} else if h[i].handType == h[j].handType {

		// assume cards i and j have equal length
		for idx := 0; idx < len(h[i].cards); idx++ {

			cardI, cardJ := string(h[i].cards[idx]), string(h[j].cards[idx])

			// Compare cards
			if cardRanks[cardI] != cardRanks[cardJ] {
				return cardRanks[cardI] < cardRanks[cardJ]
			}

		}
	}

	return false
}

func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	hands := Initialize()
	puzzleOne(hands)
}

func Initialize() Hands {
	lines, err := util.GetInput()
	if err != nil {
		log.Fatal(err)
	}

	var hands Hands

	for _, line := range lines {
		d := strings.Split(line, " ")
		cards := d[0]
		bid := util.MustToInt(d[1])
		hand := NewHand(cards, bid)
		hands = append(hands, hand)
	}

	return hands
}

func puzzleOne(hands Hands) {
	sort.Sort(hands)

	totalWinnings := 0
	for idx, hand := range hands {
		rank := idx + 1
		totalWinnings += rank * hand.bid
	}

	fmt.Printf("Puzzle one solution: %v\n", totalWinnings)
}
