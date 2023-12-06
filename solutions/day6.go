package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var numRegex = regexp.MustCompile(`\d+`)

func Day6Part1() {
	sample := `Time:      7  15   30
Distance:  9  40  200`
	input := utils.GetInputOrSample(6, sample)

	lines := strings.Split(input, "\n")
	timeMatches := numRegex.FindAllString(lines[0], -1)
	distanceMatches := numRegex.FindAllString(lines[1], -1)
	races := len(timeMatches)
	fmt.Println(timeMatches, distanceMatches)

	var time, dist int
	var winningScenarios int
	var possibilities int = 1
	for i := 0; i < races; i++ {
		winningScenarios = 0
		time, _ = strconv.Atoi(timeMatches[i])
		dist, _ = strconv.Atoi(distanceMatches[i])

		// fmt.Println("Race ", i+1)
		for j := 1; j <= time/2; j++ {
			// fmt.Printf("Hold for %vms, travel at %vmm/ms for %vms => total %v traveled\n", j, j, time-j, j*(time-j))
			if j*(time-j) > dist {
				if time%2 == 0 && j == time/2 {
					winningScenarios += 1
				} else {
					winningScenarios += 2
				}
			}
		}
		possibilities *= winningScenarios
		// fmt.Printf("We win race %v in %v scenarios\n", i+1, winningScenarios)
	}
	fmt.Println(possibilities)
}

func Day6Part2() {
	sample := `Time:      7  15   30
Distance:  9  40  200`
	input := strings.ReplaceAll(utils.GetInputOrSample(6, sample), " ", "")

	lines := strings.Split(input, "\n")
	timeMatches := numRegex.FindAllString(lines[0], -1)
	distanceMatches := numRegex.FindAllString(lines[1], -1)
	races := len(timeMatches)
	fmt.Println(timeMatches, distanceMatches)

	var time, dist int
	var winningScenarios int
	var possibilities int = 1
	for i := 0; i < races; i++ {
		winningScenarios = 0
		time, _ = strconv.Atoi(timeMatches[i])
		dist, _ = strconv.Atoi(distanceMatches[i])

		// fmt.Println("Race ", i+1)
		for j := 1; j <= time/2; j++ {
			// fmt.Printf("Hold for %vms, travel at %vmm/ms for %vms => total %v traveled\n", j, j, time-j, j*(time-j))
			if j*(time-j) > dist {
				if time%2 == 0 && j == time/2 {
					winningScenarios += 1
				} else {
					winningScenarios += 2
				}
			}
		}
		possibilities *= winningScenarios
		// fmt.Printf("We win race %v in %v scenarios\n", i+1, winningScenarios)
	}
	fmt.Println(possibilities)
}
