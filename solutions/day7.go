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
}

type Score struct {
	value uint
	index int
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

// var pair = regexp.MustCompile(`(.){2}`)
// var three = regexp.MustCompile(`(.){3}`)
// var full = regexp.MustCompile(fmt.Sprintf("^%s%s|%s%s$", pair, three, three, pair))
// var four = regexp.MustCompile(`(.){4}`)
// var five = regexp.MustCompile(`(.)`)

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
				return hand[i] != hand[i+2], i
			}
			return true, -1
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
	ok2, _ := IsPair(hand[firstPair+2:])

	return ok && ok2
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
	var scores []Score = make([]Score, handCount)
	var total uint = 0
	for i, line := range input {
		split := strings.Split(line, " ")
		bidVal, _ = strconv.ParseUint(split[1], 10, 0)
		cards = strings.Split(split[0], "")
		sort.Strings(cards)
		hands[i] = Hand{cards: split[0], bid: uint(bidVal), sorted: strings.Join(cards, "")}
		scores[i] = Score{value: uint(computeScore(hands[i].sorted)), index: i}
	}
	fmt.Printf("%+v\n", hands)
	sort.SliceStable(scores, func(i, j int) bool {
		return scores[i].value < scores[j].value
	})
	fmt.Printf("%+v\n", scores)

	for i := 0; i < handCount-1; i++ {
		fmt.Printf("Comparing %s and %s\n", hands[scores[i].index].cards, hands[scores[i+1].index].cards)
		if scores[i].value == scores[i+1].value {
			for j := 0; j < 5; j++ {
				card1 := rune(hands[scores[i].index].cards[j])
				card2 := rune(hands[scores[i+1].index].cards[j])
				fmt.Printf("\tComparing %c and %c\n", card1, card2)
				if card1 != card2 {
					if cardValues[card1] > cardValues[card2] {
						// ranks[i] = scores[i].index
						scores[i], scores[i+1] = scores[i+1], scores[i]
					}
					// else {
					// 	ranks[i] = scores[i+1].index

					// }
					break
				}
				//  else if j == 4 {
				// 	ranks[i] = scores[i].index
				// }
			}
		}
	}
	fmt.Printf("Ranking: %+v\n", scores)
	for i := 0; i < handCount; i++ {
		fmt.Printf("%s ", hands[scores[i].index].cards)
		total += (uint(i + 1)) * hands[scores[i].index].bid
	}
	fmt.Println()
	fmt.Println("Total:", total)
}
