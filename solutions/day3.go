package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile(`(\d+)`)

func isSymbol(char byte) bool {
	return !utils.IsDigit(char) && char != '.'
}

func map2dCoordTo1d(i, j, lineLen int) int {
	return i*(lineLen+3) + j
}

func Day3Part1() {
	sample := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	input := utils.GetInputOrSample(3, sample)

	lines := strings.Split(input, "\n")
	lineLen := len(lines[0])

	input = utils.PadInput(input, '.')
	totalLines := len(lines) + 2
	matches := regex.FindAllString(input, -1)
	indexes := regex.FindAllStringIndex(input, -1)

	var index int
	var nums []int
	var num int
	for i := 1; i < totalLines-1; i++ { // all lines but first and last (they are all dots)
		for j := 1; j <= lineLen; j++ { // all chars on line, except first (.) and last two (. and \n)
			index = map2dCoordTo1d(i, j, lineLen)

			if utils.IsDigit(input[index]) {
				if isSymbol(input[map2dCoordTo1d(i-1, j-1, lineLen)]) || // top-left
					isSymbol(input[map2dCoordTo1d(i-1, j, lineLen)]) || // top
					isSymbol(input[map2dCoordTo1d(i-1, j+1, lineLen)]) || // top-right
					isSymbol(input[map2dCoordTo1d(i, j-1, lineLen)]) || // left
					isSymbol(input[map2dCoordTo1d(i, j+1, lineLen)]) || // right
					isSymbol(input[map2dCoordTo1d(i+1, j-1, lineLen)]) || // bottom-left
					isSymbol(input[map2dCoordTo1d(i+1, j, lineLen)]) || // bottom
					isSymbol(input[map2dCoordTo1d(i+1, j+1, lineLen)]) { // bottom-right

					for k, match := range matches {
						if indexes[k][0] <= index && index < indexes[k][1] { // found a char in a number that we want
							j = (indexes[k][1])%(i*(lineLen+3)) - 1
							num, _ = strconv.Atoi(match)
							nums = append(nums, num)
							break
						}
					}
				}
			}
		}
	}

	fmt.Println(utils.Sum(nums))
}

func findMatchIndexesByPosition(matches [][]int, pos int) []int {
	for _, match := range matches {
		if match[0] <= pos && pos < match[1] {
			return match
		}
	}
	return nil
}

func Day3Part2() {
	sample := `.467.114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*69..
...12.34..
...$.*....
.664.598..
`
	input := utils.GetInputOrSample(3, sample)

	lines := utils.PadInputAndSplit(input, '.')

	var nums []int
	var num int
	var ratios []int

	for i, line := range lines {
		for j, char := range line {
			nums = make([]int, 0)
			num = 0
			if char != '*' {
				continue
			}

			// check if char above is a digit
			if utils.IsDigit(lines[i-1][j]) {
				num = findNumberAtCoords(lines, i-1, j)
				if num == -1 {
					continue
				}

				nums = append(nums, num)
			} else {
				// char above is not a digit, we check if chars at top-left and top-right are digits
				if utils.IsDigit(lines[i-1][j-1]) {
					num = findNumberAtCoords(lines, i-1, j-1)
					if num == -1 {
						continue
					}

					nums = append(nums, num)
				}

				if utils.IsDigit(lines[i-1][j+1]) {
					num = findNumberAtCoords(lines, i-1, j+1)
					if num == -1 {
						continue
					}

					nums = append(nums, num)
				}
			}

			// check if char on the left is digit
			if utils.IsDigit(lines[i][j-1]) {
				num = findNumberAtCoords(lines, i, j-1)
				if num == -1 {
					continue
				}

				nums = append(nums, num)
			}

			// check if char on the right is digit
			if utils.IsDigit(lines[i][j+1]) {
				num = findNumberAtCoords(lines, i, j+1)
				if num == -1 {
					continue
				}

				nums = append(nums, num)
			}

			// check if char below is a digit
			if utils.IsDigit(lines[i+1][j]) {
				num = findNumberAtCoords(lines, i+1, j)
				if num == -1 {
					continue
				}

				nums = append(nums, num)
			} else {
				// char below is not a digit, we check if chars at bottom-left and bottom-right are digits
				if utils.IsDigit(lines[i+1][j-1]) {
					num = findNumberAtCoords(lines, i+1, j-1)
					if num == -1 {
						continue
					}

					nums = append(nums, num)
				}

				if utils.IsDigit(lines[i+1][j+1]) {
					num = findNumberAtCoords(lines, i+1, j+1)
					if num == -1 {
						continue
					}

					nums = append(nums, num)
				}
			}
			if len(nums) == 2 {
				ratios = append(ratios, nums[0]*nums[1])
			}
		}
	}
	sum := utils.Sum(ratios)
	fmt.Println(sum)
}

func findNumberAtCoords(lines []string, x, y int) int {
	matches := regex.FindAllStringIndex(lines[x], -1)
	goodMatch := findMatchIndexesByPosition(matches, y)

	if goodMatch == nil {
		fmt.Printf("No match found in %+v for position %v\n", matches, y)
		return -1
	}

	num, err := strconv.Atoi(lines[x][goodMatch[0]:goodMatch[1]])
	if err != nil {
		fmt.Printf("Shit happened converting %s to int", lines[x][goodMatch[0]:goodMatch[1]])
		return -1
	}
	return num
}
