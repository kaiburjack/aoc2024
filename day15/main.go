package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input() ([][]byte, int, int, []byte) {
	file, err := os.Open("real.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var grid [][]byte
	rx, ry := 0, 0
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			break
		}
		if idx := strings.IndexByte(row, '@'); idx != -1 {
			rx, ry = idx, len(grid)
		}
		grid = append(grid, []byte(row))
	}
	var insns []byte
	for scanner.Scan() {
		insns = append(insns, scanner.Bytes()...)
	}
	return grid, rx, ry, insns
}

func moveIfPossiblePart1(grid [][]byte, x, y, dx, dy int) bool {
	nx, ny := x+dx, y+dy
	at := grid[ny][nx]
	if at == '#' {
		return false
	}
	switch at {
	case '.':
		grid[ny][nx] = grid[y][x]
		return true
	default:
		if moveIfPossiblePart1(grid, nx, ny, dx, dy) {
			grid[ny][nx] = grid[y][x]
			return true
		}
	}
	return false
}

var charToDxDy = [256][2]int{
	'<': {-1, 0},
	'>': {1, 0},
	'^': {0, -1},
	'v': {0, 1},
}

func part1(grid [][]byte, rx, ry int, insns []byte) {
	for _, insn := range insns {
		dx, dy := charToDxDy[insn][0], charToDxDy[insn][1]
		if moveIfPossiblePart1(grid, rx, ry, dx, dy) {
			grid[ry][rx] = '.'
			rx, ry = rx+dx, ry+dy
		}
	}
}

func sumOfGpsCoords(grid [][]byte) int64 {
	var sum int64
	for y, row := range grid {
		for x, c := range row {
			if c == 'O' {
				sum += int64(x) + 100*int64(y)
			}
		}
	}
	return sum
}

func main() {
	grid, rx, ry, insns := input()
	part1(grid, rx, ry, insns)
	fmt.Printf("Part1: %d\n", sumOfGpsCoords(grid))
}
