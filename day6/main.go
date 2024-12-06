package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func part2(m [][]byte, sx int, gx int, gy int, allVisited [][]int) {
	v := make([]int, len(m)*sx)
	numPossibleObstacles := 0
	for _, p := range allVisited {
		ox, oy := p[0], p[1]
		if ox == gx && oy == gy || m[oy][ox] == '#' {
			continue
		}
		m[oy][ox] = '#'
		for i := 0; i < len(v); i++ {
			v[i] = 0
		}
		x, y := gx, gy
		dx, dy := 0, -1
		for {
			if y+dy >= 0 && y+dy < len(m) && x+dx >= 0 && x+dx < sx && m[y+dy][x+dx] == '#' {
				dx, dy = -dy, dx
				continue
			}
			x, y = x+dx, y+dy
			if y < 0 || y >= len(m) || x < 0 || x >= sx {
				break
			}
			if v[y*sx+x] == 0 {
				v[y*sx+x] = (dx+1)<<2 + (dy + 1)
			} else if v[y*sx+x] == (dx+1)<<2+(dy+1) {
				numPossibleObstacles++
				break
			}
		}
		m[oy][ox] = '.'
	}
	println(numPossibleObstacles)
}

func part1(m [][]byte, sx int, gx int, gy int) [][]int {
	x, y := gx, gy
	dx, dy := 0, -1
	v := make([]bool, len(m)*sx)
	for {
		if y+dy >= 0 && y+dy < len(m) && x+dx >= 0 && x+dx < sx && m[y+dy][x+dx] == '#' {
			dx, dy = -dy, dx
			continue
		}
		v[y*sx+x] = true
		x, y = x+dx, y+dy
		if y < 0 || y >= len(m) || x < 0 || x >= sx {
			break
		}
	}
	numVisited := 0
	allVisited := make([][]int, 0)
	for i := 0; i < len(v); i++ {
		if v[i] {
			numVisited++
			allVisited = append(allVisited, []int{i % sx, i / sx})
		}
	}
	println(numVisited)
	return allVisited
}

func read() ([][]byte, int, int, int) {
	file, _ := os.Open("real.txt")
	scanner := bufio.NewScanner(file)
	var m [][]byte
	gx, gy := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		l := make([]byte, 0)
		for i := 0; i < len(line); i++ {
			l = append(l, line[i])
			if line[i] == '^' {
				gx, gy = i, len(m)
			}
		}
		m = append(m, l)
	}
	return m, len(m[0]), gx, gy
}

func main() {
	m, sx, gx, gy := read()
	allVisited := part1(m, sx, gx, gy)
	time1 := time.Now()
	part2(m, sx, gx, gy, allVisited)
	time2 := time.Now()
	fmt.Printf("Execution time Part 2: %v ms.\n", time2.Sub(time1).Milliseconds())
}