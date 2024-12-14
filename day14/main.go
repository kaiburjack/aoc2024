package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

const w, h = 101, 103

var sevenOnes = bytes.Repeat([]byte{1}, 7)

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
		rx, ry := wrap(r.x+r.vx*seconds, r.y+r.vy*seconds)
		if rx != w>>1 && ry != h>>1 {
			quadrants[rx/((w+1)>>1)+ry/((h+1)>>1)<<1]++
		}
	}
	safetyFactor := 1
	for i := 0; i < 4; i++ {
		safetyFactor *= quadrants[i]
	}
	println(safetyFactor)
}

func robotsInARow(robots []*robot) bool {
	var grid [h][w]byte
	for _, r := range robots {
		grid[r.y][r.x] = 1
	}
	for y := 0; y < h; y++ {
		if bytes.Contains(grid[y][:], sevenOnes) {
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
		if robotsInARow(robots) {
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
	in := input()
	time1 := time.Now()
	part1(in)
	part2(in)
	time2 := time.Now()
	fmt.Printf("Execution time: %v\n", time2.Sub(time1))
}
