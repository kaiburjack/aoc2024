package main

import (
	"bufio"
	"os"
)

func walk(x, y int, v byte, grid [][]byte, reachable map[[2]int]struct{}) int {
	distinct := 0
	if v == '9' {
		reachable[[2]int{x, y}] = struct{}{}
		return 1
	}
	if x > 0 && grid[y][x-1] == v+1 {
		distinct += walk(x-1, y, v+1, grid, reachable)
	}
	if x < len(grid[0])-1 && grid[y][x+1] == v+1 {
		distinct += walk(x+1, y, v+1, grid, reachable)
	}
	if y > 0 && grid[y-1][x] == v+1 {
		distinct += walk(x, y-1, v+1, grid, reachable)
	}
	if y < len(grid)-1 && grid[y+1][x] == v+1 {
		distinct += walk(x, y+1, v+1, grid, reachable)
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
	for _, sp := range startingPoints {
		reachable := make(map[[2]int]struct{})
		sumPart2 += walk(sp[0], sp[1], '0', grid, reachable)
		sumPart1 += len(reachable)
	}

	println(sumPart1)
	println(sumPart2)
}
