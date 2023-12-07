package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards  string
	sorted string
	bid    uint
	index  int
	value  uint
}

var cardValues = map[rune]uint{
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func IsFive(hand string) bool {
	return hand[0] == hand[4]
}

func IsFour(hand string) bool {
	return (hand[0] == hand[3] || hand[1] == hand[4]) && !IsFive(hand)
}

func IsFull(hand string) bool {
	return (hand[0] == hand[2] && hand[0] != hand[3] && hand[3] == hand[4]) || hand[0] == hand[1] && hand[1] != hand[2] && hand[2] == hand[4]
}

func IsPair(hand string) (bool, int) {
	if len(hand) == 5 && IsFull(hand) {
		return false, -1
	}

	for i := 0; i < len(hand)-1; i++ {
		if hand[i] == hand[i+1] {
			if i != len(hand)-2 {
				if hand[i] != hand[i+2] {
					return true, i
				}
				return false, -1
			}
			return true, i
		}
	}
	return false, -1
}

func IsThree(hand string) bool {
	if IsFull(hand) || IsFour(hand) || IsFive(hand) {
		return false
	}

	return hand[0] == hand[2] || hand[1] == hand[3] || hand[2] == hand[4]
}

func IsTwoPair(hand string) bool {
	ok, firstPair := IsPair(hand)
	if !ok {
		return false
	}

	if firstPair < 3 && hand[firstPair] == hand[firstPair+2] {
		return false
	}

	ok, _ = IsPair(hand[firstPair+2:])
	return ok
}

func computeScore(hand string) uint {
	pair, _ := IsPair(hand)
	switch true {
	case IsFive(hand):
		return 6
	case IsFour(hand):
		return 5
	case IsFull(hand):
		return 4
	case IsThree(hand):
		return 3
	case IsTwoPair(hand):
		return 2
	case pair:
		return 1
	default:
		return 0
	}
}

func Day7Part1() {
	sample := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	input := utils.GetInputOrSampleLines(7, sample)
	handCount := len(input)

	var hands []Hand = make([]Hand, handCount)
	var bidVal uint64
	var cards []string = make([]string, 5)
	var sorted string
	var total uint = 0
	for i, line := range input {
		split := strings.Split(line, " ")
		bidVal, _ = strconv.ParseUint(split[1], 10, 0)
		cards = strings.Split(split[0], "")
		sort.Strings(cards)
		sorted = strings.Join(cards, "")
		hands[i] = Hand{cards: split[0], bid: uint(bidVal), sorted: sorted, index: i, value: computeScore(sorted)}
	}

	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].value == hands[j].value {
			for k := 0; k < 5; k++ {
				if hands[i].cards[k] != hands[j].cards[k] {
					return cardValues[rune(hands[i].cards[k])] > cardValues[rune(hands[j].cards[k])]
				}
			}
		}
		return hands[i].value > hands[j].value
	})

	for i := 0; i < handCount; i++ {
		total += hands[i].bid * (uint(handCount - i))
	}
	fmt.Println("Total:", total)
}
