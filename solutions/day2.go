package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	red   []uint64
	green []uint64
	blue  []uint64
}

var gameRegex = regexp.MustCompile(`^Game (\d+)$`)
var setRegex = regexp.MustCompile(`(\d+) (red|blue|green)`)

func parseInput(lines []string) *[]game {
	var result []game = make([]game, 0)

	var matches []string
	var setMatches [][]string
	var id uint64 = 0

	for _, line := range lines { // input lines
		gameAndCubes := strings.Split(line, ": ")
		matches = gameRegex.FindStringSubmatch(gameAndCubes[0])

		id, _ = strconv.ParseUint(matches[1], 10, 0)

		sets := strings.Split(gameAndCubes[1], "; ")
		result = append(result, game{red: make([]uint64, len(sets)), blue: make([]uint64, len(sets)), green: make([]uint64, len(sets))})

		for i, set := range sets { // all numbers and colors before a semicolon
			setMatches = setRegex.FindAllStringSubmatch(set, -1)

			for _, match := range setMatches {
				switch match[2] {
				case "red":
					result[id-1].red[i], _ = strconv.ParseUint(match[1], 10, 0)
				case "blue":
					result[id-1].blue[i], _ = strconv.ParseUint(match[1], 10, 0)
				case "green":
					result[id-1].green[i], _ = strconv.ParseUint(match[1], 10, 0)
				}
			}
		}
	}
	return &result
}

func filterIds(input *[]game, red, green, blue uint64) []int {
	var ids []int
	var ok bool
	for id, g := range *input {
		ok = true
		for i := 0; i < len(g.red); i++ {
			if g.red[i] > red || g.green[i] > green || g.blue[i] > blue {
				ok = false
			}
		}
		if ok {
			ids = append(ids, id+1)
		}
	}
	return ids
}

func getPowers(input *[]game) []uint64 {
	var result = make([]uint64, len(*input))
	var red, green, blue uint64

	for j, g := range *input {
		red = 0
		green = 0
		blue = 0
		for i := 0; i < len(g.red); i++ {
			if g.red[i] > red {
				red = g.red[i]
			}
			if g.green[i] > green {
				green = g.green[i]
			}
			if g.blue[i] > blue {
				blue = g.blue[i]
			}
		}

		result[j] = red * green * blue
	}

	return result
}

func Day2Part1() {
	sample := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	lines := utils.GetInputOrSampleLines(2, sample)
	input := parseInput(lines)
	ids := filterIds(input, 12, 13, 14)

	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum:", sum)
}

func Day2Part2() {
	sample := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	lines := utils.GetInputOrSampleLines(2, sample)
	input := parseInput(lines)
	powers := getPowers(input)

	var sum uint64 = 0
	for _, p := range powers {
		sum += p
	}

	fmt.Println("Sum:", sum)
}
