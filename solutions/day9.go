package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"slices"
	"strings"
)

func coef(index, differenceIndex int) int {
	if index == 0 {
		return 1
	}
	return coef(index-1, differenceIndex) * (differenceIndex - index + 1) / index
}

func Day9Part1() {
	sample := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	input := utils.GetInputOrSampleLines(9, sample)
	var seqs [][]int = make([][]int, len(input))

	for i, line := range input {
		seqs[i] = utils.MapStringsToInts(strings.Split(line, " "))
	}

	var total = 0
	var val int

	for _, seq := range seqs {
		total += seq[len(seq)-1]
		allZeroes := false
		k := 0

		for !allZeroes {
			val = 0
			k++
			allZeroes = true
			for i := 0; i < len(seq)-k; i++ {
				val = 0
				for j := 0; j <= k; j++ {
					if j%2 != 0 {
						val += (-1) * coef(j, k) * seq[i+k-j]
					} else {
						val += coef(j, k) * seq[i+k-j]
					}
				}
				if val != 0 {
					allZeroes = false
				}
			}
			total += val
		}
	}
	fmt.Println("Part 1:", total)
}

func Day9Part2() {
	sample := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	input := utils.GetInputOrSampleLines(9, sample)
	var seqs [][]int = make([][]int, len(input))

	for i, line := range input {
		seqs[i] = utils.MapStringsToInts(strings.Split(line, " "))
		slices.Reverse(seqs[i])
	}

	var total = 0
	var val int

	for _, seq := range seqs {
		total += seq[len(seq)-1]
		allZeroes := false
		k := 0

		for !allZeroes {
			val = 0
			k++
			allZeroes = true
			for i := 0; i < len(seq)-k; i++ {
				val = 0
				for j := 0; j <= k; j++ {
					if j%2 != 0 {
						val += (-1) * coef(j, k) * seq[i+k-j]
					} else {
						val += coef(j, k) * seq[i+k-j]
					}
				}
				if val != 0 {
					allZeroes = false
				}
			}
			total += val
		}
	}
	fmt.Println("Part 2:", total)
}
