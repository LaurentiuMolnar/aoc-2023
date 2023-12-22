package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
	"strings"
)

type ModuleType int

const (
	BROADCASTER ModuleType = iota
	FLIP_FLOP
	CONJUNCTION
)

type Module = struct {
	destinations []string
	moduleType   ModuleType
}

func parseModules(input []string) map[string]Module {
	var modules map[string]Module = make(map[string]Module)

	for _, line := range input {
		parts := strings.Split(line, " -> ")

		source := parts[0]
		dest := parts[1]

		destinations := strings.Split(dest, ", ")

		moduleType := BROADCASTER
		if source == "broadcaster" {
			moduleType = BROADCASTER
		} else if source[0] == '%' {
			moduleType = FLIP_FLOP
		} else {
			moduleType = CONJUNCTION
		}

		source = strings.Trim(source, "%&")

		modules[source] = Module{
			moduleType:   moduleType,
			destinations: destinations,
		}
	}

	return modules
}

func Day20Part1() {
	sample := `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`
	input := utils.GetInputOrSampleLines(20, sample)
	modules := parseModules(input)

	fmt.Println(modules)
}
