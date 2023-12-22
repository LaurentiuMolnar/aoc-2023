package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"slices"
	"strings"
)

func getEmptyCols(input []string) []int {
	var result []int
	colCount := len(input[0])
	lineCount := len(input)

	for j := 0; j < colCount; j++ {
		for i := 0; i < lineCount; i++ {
			if input[i][j] != '.' {
				break
			}
			if i == lineCount-1 && !slices.Contains(result, j) {
				result = append(result, j)
			}
		}
	}

	return result
}

func expandInput(input []string) []string {
	var result []string
	for i := 0; i < len(input); i++ {
		result = append(result, input[i])
		if strings.Count(input[i], ".") == len(input[i]) {
			fmt.Printf("Line %v empty\n", i)
			result = append(result, input[i])
		}
	}
	return result
}

func expandInputCols(input []string) []string {
	emptyCols := getEmptyCols(input)
	var result []string = make([]string, len(input))
	// var lineChars []byte
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if slices.Contains(emptyCols, j) {
				result[i] = input[i][0:j+1] + "." + input[i][j+1:]
			} else {
				result[i] = input[i]
			}
		}
	}
	return result
}

func Day11Part1() {
	sample := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	input := utils.GetInputOrSampleLines(11, sample)
	fmt.Println(input)
	input = expandInput(input)
	fmt.Println(input)
	emptyCols := getEmptyCols(input)
	fmt.Println(emptyCols)
	input = expandInputCols(input)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			fmt.Printf("%c", input[i][j])
		}
		fmt.Println()
	}
}
