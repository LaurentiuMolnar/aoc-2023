package solutions

import (
	utils "aoc-2023/lib"
	"fmt"
)

func getStartCoords(maze []string) (int, int) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if (maze[i][j]) == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

func Day10Part1() {
	sample := `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`
	maze := utils.GetInputOrSampleLines(10, sample)

	startX, startY := getStartCoords(maze)
	fmt.Println("Start at", startX, startY)
}
