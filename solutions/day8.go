package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"regexp"

	"modernc.org/mathutil" // it was 2 a.m., did not feel like implementing prime factorization myself
)

type HauntedNode struct {
	left  string
	right string
}

var lineRegex = regexp.MustCompile(`^([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)$`)

func Day8Part1() {
	sample := `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`
	input := utils.GetInputOrSampleLines(8, sample)

	directions := input[0]
	input = input[2:]

	var nodes map[string](HauntedNode) = make(map[string]HauntedNode)

	var matches [][]string
	for _, line := range input {
		matches = lineRegex.FindAllStringSubmatch(line, -1)
		rootValue, leftValue, rightValue := matches[0][1], matches[0][2], matches[0][3]
		nodes[rootValue] = HauntedNode{left: leftValue, right: rightValue}
	}

	var node = "AAA"
	var count = 0
	for node != "ZZZ" {
		for _, dir := range directions {
			fmt.Println(node)
			if dir == 'L' {
				node = nodes[node].left
			} else {
				node = nodes[node].right
			}
			count++
		}
	}
	fmt.Println(count)
}

func Day8Part2() {
	sample := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
	input := utils.GetInputOrSampleLines(8, sample)

	directions := input[0]
	input = input[2:]

	var nodes map[string](HauntedNode) = make(map[string]HauntedNode)

	var matches [][]string
	var startNodes []string
	for _, line := range input {
		matches = lineRegex.FindAllStringSubmatch(line, -1)
		rootValue, leftValue, rightValue := matches[0][1], matches[0][2], matches[0][3]
		if rootValue == leftValue && leftValue == rightValue {
			fmt.Println(rootValue)
		}
		nodes[rootValue] = HauntedNode{left: leftValue, right: rightValue}
		if rootValue[2] == 'A' {
			startNodes = append(startNodes, rootValue)
		}
	}

	var count uint32 = 0
	var factors map[uint32]uint32 = make(map[uint32]uint32)

	var n string
	var total uint64 = 1
	for _, node := range startNodes {
		n = node
		count = 0
		for n[2] != 'Z' {
			for _, dir := range directions {
				if dir == 'L' {
					n = nodes[n].left
				} else {
					n = nodes[n].right
				}
				count++
			}
		}
		terms := mathutil.FactorInt(count)

		for _, factor := range terms {
			_, exists := factors[factor.Prime]
			if !exists || (exists && factors[factor.Prime] > factor.Power) {
				factors[factor.Prime] = factor.Power
			}
		}
	}

	var i uint32
	for factor, power := range factors {
		for i = 1; i <= power; i++ {
			total *= uint64(factor)
		}
	}

	fmt.Println(factors)
	fmt.Println("Total:", total)
}
