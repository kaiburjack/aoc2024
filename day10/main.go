package main

import (
	"bufio"
	"os"
)

func walk(p [2]int, v byte, grid [][]byte, reachable map[[2]int]struct{}) int {
	distinct := 0
	if v == '9' {
		reachable[p] = struct{}{}
		return 1
	}
	if p[0] > 0 && grid[p[1]][p[0]-1] == v+1 {
		distinct += walk([2]int{p[0] - 1, p[1]}, v+1, grid, reachable)
	}
	if p[0] < len(grid[0])-1 && grid[p[1]][p[0]+1] == v+1 {
		distinct += walk([2]int{p[0] + 1, p[1]}, v+1, grid, reachable)
	}
	if p[1] > 0 && grid[p[1]-1][p[0]] == v+1 {
		distinct += walk([2]int{p[0], p[1] - 1}, v+1, grid, reachable)
	}
	if p[1] < len(grid)-1 && grid[p[1]+1][p[0]] == v+1 {
		distinct += walk([2]int{p[0], p[1] + 1}, v+1, grid, reachable)
	}
	return distinct
}

func main() {
	file, _ := os.Open("real.txt")
	r := bufio.NewScanner(file)
	var grid [][]byte
	var startingPoints [][2]int
	for r.Scan() {
		var row []byte
		for _, c := range r.Bytes() {
			row = append(row, c)
			if c == '0' {
				startingPoints = append(startingPoints, [2]int{len(row) - 1, len(grid)})
			}
		}
		grid = append(grid, row)
	}

	sumPart2, sumPart1 := 0, 0
	for sp := range startingPoints {
		reachable := make(map[[2]int]struct{})
		sumPart2 += walk(startingPoints[sp], '0', grid, reachable)
		sumPart1 += len(reachable)
	}

	println(sumPart1)
	println(sumPart2)
}
