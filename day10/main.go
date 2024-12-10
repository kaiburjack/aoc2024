package main

import (
	"bufio"
	"os"
)

func walk(x, y int, v byte, grid [][]byte, reachable map[[2]int]struct{}) int {
	if v == '9' {
		reachable[[2]int{x, y}] = struct{}{}
		return 1
	}
	distinct := 0
	dirs := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, d := range dirs {
		if x+d[0] >= 0 && x+d[0] < len(grid[0]) && y+d[1] >= 0 && y+d[1] < len(grid) && grid[y+d[1]][x+d[0]] == v+1 {
			distinct += walk(x+d[0], y+d[1], v+1, grid, reachable)
		}
	}
	return distinct
}

func main() {
	file, _ := os.Open("real.txt")
	r := bufio.NewScanner(file)
	var grid [][]byte
	var startingPoints [][2]int
	for r.Scan() {
		row := r.Bytes()
		for x, c := range row {
			if c == '0' {
				startingPoints = append(startingPoints, [2]int{x, len(grid)})
			}
		}
		grid = append(grid, row)
	}

	sumPart1, sumPart2 := 0, 0
	for _, sp := range startingPoints {
		reachable := make(map[[2]int]struct{})
		sumPart2 += walk(sp[0], sp[1], '0', grid, reachable)
		sumPart1 += len(reachable)
	}

	println(sumPart1)
	println(sumPart2)
}
