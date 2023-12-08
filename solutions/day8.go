package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"regexp"
)

type HauntedNode struct {
	left  string
	right string
}

var lineRegex = regexp.MustCompile(`^([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)$`)

func countEnds(nodes []string) int {
	var count int = 0
	for _, node := range nodes {
		if node[2] == 'Z' {
			count++
		}
	}
	return count
}

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
	// fmt.Println(startNodes)

	var count = 0
	var destinations []string = make([]string, len(startNodes))
	var ends = 0

	for ends != len(startNodes) {
		for _, dir := range directions {
			fmt.Println(startNodes, dir)
			for i, node := range startNodes {
				// fmt.Println(node)
				if dir == 'L' {
					destinations[i] = nodes[node].left
				} else {
					destinations[i] = nodes[node].right
				}
			}
			// fmt.Println(destinations)
			copy(startNodes, destinations)
			count++
			ends = countEnds(startNodes)
			if ends == len(startNodes) {
				break
			}
		}
	}

	fmt.Println(count)
}
