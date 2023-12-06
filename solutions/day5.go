package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type RuleMap struct {
	srcStart  uint64
	destStart uint64
	size      uint64
}

var mapLineRegex = regexp.MustCompile(`(\d+)`)
var part2Regex = regexp.MustCompile(`(\d+) (\d+)`)

func buildMap(inputLines string) *[]RuleMap {
	var srcStart uint64
	var destStart uint64
	var size uint64
	var matches []string
	var result []RuleMap
	for _, line := range strings.Split(strings.Trim(inputLines, "\n"), "\n") {
		matches = mapLineRegex.FindAllString(line, -1)
		destStart, _ = strconv.ParseUint(matches[0], 10, 64)
		srcStart, _ = strconv.ParseUint(matches[1], 10, 64)
		size, _ = strconv.ParseUint(matches[2], 10, 64)

		result = append(result, RuleMap{destStart: destStart, srcStart: srcStart, size: size})
	}

	return &result
}

func getDestinationValue(srcValue uint64, rules *[]RuleMap) uint64 {
	for _, rule := range *rules {
		if rule.srcStart <= srcValue && srcValue < rule.srcStart+rule.size {
			// found a range containing our srcValue
			return rule.destStart + srcValue - rule.srcStart
		}
	}
	return srcValue
}

func getLocationBySeed(seed uint64, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHum, humToLoc *[]RuleMap) uint64 {
	soil := getDestinationValue(seed, seedToSoil)
	fertilizer := getDestinationValue(soil, soilToFert)
	water := getDestinationValue(fertilizer, fertToWater)
	light := getDestinationValue(water, waterToLight)
	temperature := getDestinationValue(light, lightToTemp)
	humidity := getDestinationValue(temperature, tempToHum)
	return getDestinationValue(humidity, humToLoc)
}

func solve(seeds []uint64, input string) uint64 {
	rest, hToL, _ := strings.Cut(input, "\nhumidity-to-location map:\n")
	rest, tToH, _ := strings.Cut(rest, "\ntemperature-to-humidity map:\n")
	rest, lToT, _ := strings.Cut(rest, "\nlight-to-temperature map:\n")
	rest, wToL, _ := strings.Cut(rest, "\nwater-to-light map:\n")
	rest, fToW, _ := strings.Cut(rest, "\nfertilizer-to-water map:\n")
	rest, soilToF, _ := strings.Cut(rest, "\nsoil-to-fertilizer map:\n")
	_, sdToSl, _ := strings.Cut(rest, "\nseed-to-soil map:\n")

	humToLoc := buildMap(hToL)
	tempToHum := buildMap(tToH)
	lightToTemp := buildMap(lToT)
	waterToLight := buildMap(wToL)
	fertToWater := buildMap(fToW)
	soilToFert := buildMap(soilToF)
	seedToSoil := buildMap(sdToSl)

	var locations []uint64
	for _, seed := range seeds {
		locations = append(locations, getLocationBySeed(seed, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHum, humToLoc))
	}

	return utils.Min(locations)
}

func Day5Part1() {
	sample := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`
	input := utils.GetInputOrSample(5, sample)
	seedString, _ := strings.CutPrefix(input, "seeds: ")
	seedString, _, _ = strings.Cut(seedString, "\n")

	var seed uint64
	var seeds []uint64
	for _, seedStr := range mapLineRegex.FindAllString(seedString, -1) {
		seed, _ = strconv.ParseUint(seedStr, 10, 64)
		seeds = append(seeds, seed)
	}

	solution := solve(seeds, input)
	fmt.Println(solution)
}

func Day5Part2() {
	sample := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`
	input := utils.GetInputOrSample(5, sample)
	seedString, _ := strings.CutPrefix(input, "seeds: ")
	seedString, _, _ = strings.Cut(seedString, "\n")

	rest, hToL, _ := strings.Cut(input, "\nhumidity-to-location map:\n")
	rest, tToH, _ := strings.Cut(rest, "\ntemperature-to-humidity map:\n")
	rest, lToT, _ := strings.Cut(rest, "\nlight-to-temperature map:\n")
	rest, wToL, _ := strings.Cut(rest, "\nwater-to-light map:\n")
	rest, fToW, _ := strings.Cut(rest, "\nfertilizer-to-water map:\n")
	rest, soilToF, _ := strings.Cut(rest, "\nsoil-to-fertilizer map:\n")
	_, sdToSl, _ := strings.Cut(rest, "\nseed-to-soil map:\n")

	humToLoc := buildMap(hToL)
	tempToHum := buildMap(tToH)
	lightToTemp := buildMap(lToT)
	waterToLight := buildMap(wToL)
	fertToWater := buildMap(fToW)
	soilToFert := buildMap(soilToF)
	seedToSoil := buildMap(sdToSl)

	var location uint64 = math.MaxUint64
	var loc uint64
	var seed uint64
	var size uint64
	for _, seedStr := range part2Regex.FindAllStringSubmatch(seedString, -1) {
		seed, _ = strconv.ParseUint(seedStr[1], 10, 64)
		size, _ = strconv.ParseUint(seedStr[2], 10, 64)

		for i := seed; i < seed+size; i++ {
			loc = getLocationBySeed(i, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHum, humToLoc)
			if loc < location {
				location = loc
			}
		}
	}

	fmt.Println(location)
}
