package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func input() ([][]byte, int, int, []byte) {
	file, _ := os.Open("real.txt")
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

func isPossiblePart2(grid [][]byte, x, y, dx, dy int, moves *[][2]int) bool {
	nx, ny := x+dx, y+dy
	at := grid[ny][nx]
	if at == '#' {
		return false
	}
	switch at {
	case '.':
		return true
	case '[':
		if dx == 1 && isPossiblePart2(grid, nx+1, ny, dx, dy, moves) ||
			dx != 1 && isPossiblePart2(grid, nx, ny, dx, dy, moves) &&
				isPossiblePart2(grid, nx+1, ny, dx, dy, moves) {
			*moves = append(*moves, [2]int{nx, ny})
			return true
		}
	case ']':
		if dx == -1 && isPossiblePart2(grid, nx-1, ny, dx, dy, moves) ||
			dx != -1 && isPossiblePart2(grid, nx-1, ny, dx, dy, moves) &&
				isPossiblePart2(grid, nx, ny, dx, dy, moves) {
			*moves = append(*moves, [2]int{nx - 1, ny})
			return true
		}
	}
	return false
}

func part2(grid [][]byte, rx, ry int, insns []byte) [][]byte {
	for _, insn := range insns {
		dx, dy := charToDxDy[insn][0], charToDxDy[insn][1]
		var moves [][2]int
		if isPossiblePart2(grid, rx, ry, dx, dy, &moves) {
			for i := 0; i < len(moves); i++ {
				x, y := moves[i][0], moves[i][1]
				grid[y][x], grid[y][x+1] = '.', '.'
			}
			for i := 0; i < len(moves); i++ {
				x, y := moves[i][0], moves[i][1]
				grid[y+dy][x+dx], grid[y+dy][x+dx+1] = '[', ']'
			}
			grid[ry][rx], grid[ry+dy][rx+dx] = '.', '@'
			rx, ry = rx+dx, ry+dy
		}
	}
	return grid
}

func sumOfGpsCoords(grid [][]byte) int64 {
	var sum int64
	for y, row := range grid {
		for x, c := range row {
			if c == 'O' || c == '[' {
				sum += int64(x) + 100*int64(y)
			}
		}
	}
	return sum
}

func extendGrid(grid [][]byte) [][]byte {
	biggerGrid := make([][]byte, len(grid))
	for row := range grid {
		biggerGrid[row] = make([]byte, 0, len(grid[row])<<1)
		for _, c := range grid[row] {
			switch c {
			case '#':
				biggerGrid[row] = append(biggerGrid[row], '#', '#')
			case '.':
				biggerGrid[row] = append(biggerGrid[row], '.', '.')
			case 'O':
				biggerGrid[row] = append(biggerGrid[row], '[', ']')
			case '@':
				biggerGrid[row] = append(biggerGrid[row], '@', '.')
			default:
			}
		}
	}
	return biggerGrid
}

func main() {
	grid, rx, ry, insns := input()
	largerGrid := extendGrid(grid)
	time0 := time.Now()
	part1(grid, rx, ry, insns)
	time1 := time.Now()
	fmt.Printf("Part1: %d (in %v)\n", sumOfGpsCoords(grid), time1.Sub(time0))
	time2 := time.Now()
	g := part2(largerGrid, rx<<1, ry, insns)
	time3 := time.Now()
	fmt.Printf("Part2: %d (in %v)\n", sumOfGpsCoords(g), time3.Sub(time2))
}
