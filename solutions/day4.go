package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"strconv"
	"strings"
)

func Day4Part1() {
	sample := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	input := utils.GetInputOrSampleLines(4, sample)
	lineCount := len(input)
	var winners [][]int = make([][]int, lineCount)
	var attempts [][]int = make([][]int, lineCount)
	var scores []int = make([]int, lineCount)

	var parts []string
	var matches = 0
	var n int
	for i, line := range input {
		matches = 0
		line = strings.TrimRight(strings.Split(line, ": ")[1], "\n")

		// if i <= 99 {
		// 	fmt.Println(line)
		// }

		parts = strings.Split(line, " | ")

		for _, num := range strings.Split(parts[0], " ") {
			n, _ = strconv.Atoi(num)
			if n != 0 {
				winners[i] = append(winners[i], n)
			}
		}

		for _, num := range strings.Split(parts[1], " ") {
			n, _ = strconv.Atoi(num)
			if n != 0 {
				attempts[i] = append(attempts[i], n)
			}
		}

		for j := 0; j < len(winners[i]); j++ {
			for k := 0; k < len(attempts[i]); k++ {
				if attempts[i][k] == winners[i][j] {
					matches++
				}
			}
		}

		if matches == 0 {
			scores[i] = matches
		} else {
			scores[i] = 1 << (matches - 1)
		}
		fmt.Println(matches)
	}
	fmt.Println(utils.Sum(scores))
}

func Day4Part2() {
	sample := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	input := utils.GetInputOrSampleLines(4, sample)
	lineCount := len(input)
	var winners [][]int = make([][]int, lineCount)
	var attempts [][]int = make([][]int, lineCount)
	var copies []int = make([]int, lineCount)

	var parts []string
	var matches = 0
	var n int
	var bound int
	for i, line := range input {
		matches = 0
		line = strings.TrimRight(strings.Split(line, ": ")[1], "\n")
		copies[i] += 1

		parts = strings.Split(line, " | ")

		for _, num := range strings.Split(parts[0], " ") {
			n, _ = strconv.Atoi(num)
			if n != 0 {
				winners[i] = append(winners[i], n)
			}
		}

		for _, num := range strings.Split(parts[1], " ") {
			n, _ = strconv.Atoi(num)
			if n != 0 {
				attempts[i] = append(attempts[i], n)
			}
		}

		for j := 0; j < len(winners[i]); j++ {
			for k := 0; k < len(attempts[i]); k++ {
				if attempts[i][k] == winners[i][j] {
					matches++
				}
			}
		}

		if matches == 0 {
			continue
		}

		if i+matches > lineCount {
			bound = lineCount
		} else {
			bound = i + matches
		}

		fmt.Printf(`Card %d has %d matches. We get copies of cards`, i, matches)
		for q := i + 1; q < bound; q++ {
			fmt.Printf("%v ", q)
		}
		fmt.Println()
	}
}
