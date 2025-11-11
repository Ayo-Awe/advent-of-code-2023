package main

import (
	"flag"
	"fmt"
	"log"
	"slices"
	"strings"

	aoc "github.com/ayo-awe/advent-of-code-2023/util"
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

	Wildcard = 'J'
)

type Hand struct {
	cards      string
	bid        int
	handType   HandType
	cardCounts map[rune]int
}

func (h Hand) String() string {
	return fmt.Sprintf("Hand(%q %d %d)", h.cards, h.bid, h.handType)
}

func (hand Hand) PromotedType() HandType {
	// Doesn't have a wildcard, can't be promoted
	if hand.cardCounts[Wildcard] == 0 {
		return hand.handType
	}

	switch hand.handType {
	case FiveOfAKind, FourOfAKind, FullHouse:
		return FiveOfAKind
	case ThreeOfAKind:
		return FourOfAKind
	case TwoPair:
		if hand.cardCounts[Wildcard] == 1 {
			return FullHouse
		}
		return FourOfAKind
	case OnePair:
		if hand.cardCounts[Wildcard] == 1 {
			return TwoPair
		}
		return ThreeOfAKind
	case HighCard:
		return OnePair
	default:
		panic("unknown hand type")
	}
}

func NewHand(cards string, bid int) Hand {
	cardCounts := GetCardCounts(cards)
	handType := GetHandType(cardCounts)
	return Hand{cards, bid, handType, cardCounts}
}

func GetCardCounts(cards string) map[rune]int {
	cardCounts := make(map[rune]int)
	for _, card := range cards {
		cardCounts[card] += 1
	}
	return cardCounts
}

func GetHandType(cardCounts map[rune]int) HandType {
	var dist []int
	for _, value := range cardCounts {
		dist = append(dist, value)
	}

	handTypeDistribution := map[HandType][]int{
		FiveOfAKind:  {5},
		FourOfAKind:  {1, 4},
		FullHouse:    {2, 3},
		ThreeOfAKind: {1, 1, 3},
		TwoPair:      {1, 2, 2},
		OnePair:      {1, 1, 1, 2},
		HighCard:     {1, 1, 1, 1, 1},
	}

	slices.Sort(dist)
	for handType, value := range handTypeDistribution {
		if slices.Equal(dist, value) {
			return handType
		}
	}

	panic("invalid hand")
}

func main() {
	filename := flag.String("file", "input.txt", "input file name")
	flag.Parse()

	lines, err := aoc.ReadInputLineByLine(*filename)
	if err != nil {
		log.Fatal(err)
	}

	hands, err := ParseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part one: ", PartOne(hands))
	fmt.Println("Part two: ", PartTwo(hands))
}

func ParseInput(lines []string) ([]Hand, error) {
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		d := strings.Split(line, " ")
		cards := d[0]
		bid := aoc.MustToInt(d[1])
		hand := NewHand(cards, bid)
		hands[i] = hand
	}

	return hands, nil
}

func PartOne(hands []Hand) int {
	slices.SortFunc(hands, Part1HandCmpFunc)

	totalWinnings := 0
	for idx, hand := range hands {
		rank := idx + 1
		totalWinnings += rank * hand.bid
	}

	return totalWinnings
}

func PartTwo(hands []Hand) int {
	slices.SortFunc(hands, Part2HandCmpFunc)

	totalWinnings := 0
	for idx, hand := range hands {
		rank := idx + 1
		totalWinnings += rank * hand.bid
	}

	return totalWinnings
}

// Comp functions
func Part2HandCmpFunc(a, b Hand) int {
	cards := "AKQT98765432J"

	cardRanks := map[rune]int{}
	for i, card := range cards {
		cardRanks[card] = len(cards) - i
	}

	if a.PromotedType() != b.PromotedType() {
		return int(a.PromotedType() - b.PromotedType())
	}

	// assume cards i and j have equal length
	for i := 0; i < len(a.cards); i++ {
		cardA, cardB := rune(a.cards[i]), rune(b.cards[i])

		// Compare cards
		if cardRanks[cardA] != cardRanks[cardB] {
			return cardRanks[cardA] - cardRanks[cardB]
		}
	}

	return 0
}

func Part1HandCmpFunc(a, b Hand) int {
	cards := "AKQJT98765432"
	cardRanks := map[rune]int{}
	for i, card := range cards {
		cardRanks[card] = len(cards) - i
	}

	if a.handType != b.handType {
		return int(a.handType - b.handType)
	}

	// assume cards i and j have equal length
	for i := 0; i < len(a.cards); i++ {
		cardA, cardB := rune(a.cards[i]), rune(b.cards[i])

		// Compare cards
		if cardRanks[cardA] != cardRanks[cardB] {
			return cardRanks[cardA] - cardRanks[cardB]
		}
	}

	return 0
}
