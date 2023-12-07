package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	highCard uint = iota
	onePair
	twoPairs
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
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

func computeScoreWithJokers(hand string) uint {
	pair, _ := IsPair(hand)
	jokers := countJokers(hand)

	if IsFive(hand) {
		return fiveOfAKind
	}

	if IsFour(hand) {
		if jokers == 1 || jokers == 4 {
			return fiveOfAKind
		}
		return fourOfAKind
	}

	if IsFull(hand) {
		if jokers == 2 || jokers == 3 { // only possible numbers of jokers in a full
			return fiveOfAKind // if we have a full containing jokers, it's the same as having five identical cards
		}

		return fullHouse
	}

	if IsThree(hand) {
		if jokers == 1 || jokers == 3 {
			// 1 joker + three of a kind => four of a kind
			return fourOfAKind
		}

		// 2 jokers + three of a kind is a full-house, so it's handled by previous if
		return threeOfAKind
	}

	if IsTwoPair(hand) {
		if jokers == 1 { // two pairs + 1 joker => full-house
			return fullHouse
		}

		if jokers == 2 {
			// 1 pair + 1 pair of jokers => four of a kind
			return fourOfAKind
		}

		return twoPairs
	}

	if pair {
		if jokers == 1 || jokers == 2 {
			return threeOfAKind
		}

		// 3 jokers + 1 pair means full-house
		return onePair
	}

	// the hand is high-card (no special combinations)
	if jokers == 1 {
		return onePair
	}
	// other numbers of jokers without other special combinations result in pairs or three/four/five of a kind, which are handled already

	return highCard
}

func buildHands(input []string, scoringFunc func(hand string) uint) *[]Hand {
	handCount := len(input)
	var hands []Hand = make([]Hand, handCount)
	var bidVal uint64
	var cards []string = make([]string, 5)
	var sorted string

	for i, line := range input {
		split := strings.Split(line, " ")
		bidVal, _ = strconv.ParseUint(split[1], 10, 0)
		cards = strings.Split(split[0], "")
		sort.Strings(cards)
		sorted = strings.Join(cards, "")
		hands[i] = Hand{
			cards:  split[0],
			bid:    uint(bidVal),
			sorted: sorted,
			index:  i,
			value:  scoringFunc(sorted),
		}
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
	return &hands
}

func countJokers(hand string) uint {
	return uint(strings.Count(hand, "J"))
}

func computeScore(hand string) uint {
	pair, _ := IsPair(hand)
	switch true {
	case IsFive(hand):
		return fiveOfAKind
	case IsFour(hand):
		return fourOfAKind
	case IsFull(hand):
		return fullHouse
	case IsThree(hand):
		return threeOfAKind
	case IsTwoPair(hand):
		return twoPairs
	case pair:
		return onePair
	default:
		return highCard
	}
}

func Day7Part1() {
	sample := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	input := utils.GetInputOrSampleLines(7, sample)
	hands := buildHands(input, computeScore)
	handCount := len(input)

	var total uint
	for i := 0; i < handCount; i++ {
		total += (*hands)[i].bid * (uint(handCount - i))
	}
	fmt.Println("Total:", total)
}

func Day7Part2() {
	sample := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	input := utils.GetInputOrSampleLines(7, sample)
	cardValues['J'] = 0 // we won't run part1 after part2, so it's no problem reassigning J's value in part 2

	hands := buildHands(input, computeScoreWithJokers)
	handCount := len(input)

	var total uint
	for i := 0; i < handCount; i++ {
		total += (*hands)[i].bid * (uint(handCount - i))
	}
	fmt.Println("Total:", total)
}
