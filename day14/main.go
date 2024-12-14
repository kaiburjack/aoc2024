package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

const w, h = 101, 103

type robot struct {
	x, y, vx, vy int64
}

func printGrid(robots []*robot) {
	var grid [h][w]byte
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			grid[y][x] = '.'
		}
	}
	for _, r := range robots {
		grid[r.y][r.x] = '#'
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			print(string(grid[y][x]))
		}
		println()
	}
}

func wrap(a, b int64) (int64, int64) {
	return (a%w + w) % w, (b%h + h) % h
}

func part1(robots []*robot) {
	var quadrants [4]int
	const seconds = 100
	for _, r := range robots {
		r.x, r.y = wrap(r.x+r.vx*seconds, r.y+r.vy*seconds)
		if r.x == w>>1 || r.y == h>>1 {
			continue
		}
		quadrants[r.x/((w+1)>>1)+r.y/((h+1)>>1)<<1]++
	}
	safetyFactor := 1
	for i := 0; i < 4; i++ {
		safetyFactor *= quadrants[i]
	}
	println(safetyFactor)
}

func robotsInARow(robots []*robot, n int) bool {
	var grid [h][w]bool
	for _, r := range robots {
		grid[r.y][r.x] = true
	}
	for y := 0; y < h-n; y++ {
	inner:
		for x := 0; x < w-n; x++ {
			if !grid[y][x] {
				continue
			}
			for i := 1; i < n; i++ {
				if grid[y+i][x] != grid[y][x] {
					continue inner
				}
			}
			return true
		}
	}
	return false
}

func part2(robots []*robot) {
	var numSeconds int64
	for {
		for _, r := range robots {
			r.x, r.y = wrap(r.x+r.vx, r.y+r.vy)
		}
		numSeconds++
		if robotsInARow(robots, 8) {
			break
		}
	}
	//printGrid(robots)
	println(numSeconds)
}

func input() []*robot {
	file, _ := os.Open("real.txt")
	scanner := bufio.NewScanner(file)
	lineRegex := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	var robots []*robot
	for scanner.Scan() {
		strToInt64 := func(s string) int64 {
			i, _ := strconv.ParseInt(s, 10, 64)
			return i
		}
		line := scanner.Text()
		matches := lineRegex.FindStringSubmatch(line)
		x, y := strToInt64(matches[1]), strToInt64(matches[2])
		vx, vy := strToInt64(matches[3]), strToInt64(matches[4])
		robots = append(robots, &robot{x, y, vx, vy})
	}
	return robots
}

func main() {
	part1(input())
	part2(input())
}