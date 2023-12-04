package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var numStringsMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var numStrings = []string{
	"one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9",
}

func Day1Part1() {
	sample := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
eightwo
`
	regex, _ := regexp.Compile(`(\d)`)
	input := utils.GetInputOrSample(1, sample)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]

	var sum int = 0

	for _, line := range lines {
		found := regex.FindAllString(line, -1)
		s := found[0] + found[len(found)-1]

		num, _ := strconv.Atoi(s)
		sum += num
	}
	fmt.Println(sum)
}

func Day1Part2() {
	sample := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
eightwo
`
	input := utils.GetInputOrSample(1, sample)
	lines := strings.Split(input, "\n")

	var sum int = 0

	for _, line := range lines {
		firstIndex := len(line)
		first := ""
		lastIndex := -1
		last := ""
		for _, numStr := range numStrings {
			fi := strings.Index(line, numStr)
			if fi < firstIndex && fi >= 0 {
				firstIndex = fi
				first = numStr
			}

			li := strings.LastIndex(line, numStr)
			if li > lastIndex {
				lastIndex = li
				last = numStr
			}
		}

		firstNum, err := strconv.Atoi(first)
		if err != nil {
			firstNum = numStringsMap[first]
		}

		lastNum, err := strconv.Atoi(last)
		if err != nil {
			lastNum = numStringsMap[last]
		}

		num := 10*firstNum + lastNum
		sum += num
	}
	fmt.Println(sum)
}
