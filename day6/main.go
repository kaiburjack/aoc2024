package main

import (
	"bufio"
	"os"
)

func part2() {
	m, sx, gx, gy := read()
	v := make([]int, len(m)*sx)
	numPossibleObstacles := 0
	for oy := 0; oy < len(m); oy++ {
		for ox := 0; ox < sx; ox++ {
			if ox == gx && oy == gy || m[oy][ox] == '#' {
				continue
			}
			m[oy][ox] = '#'
			loopDetected := false
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
					loopDetected = true
					break
				}
			}
			if loopDetected {
				numPossibleObstacles++
			}
			m[oy][ox] = '.'
		}
	}
	println(numPossibleObstacles)
}

func part1() {
	m, sx, gx, gy := read()
	v := make([]bool, len(m)*sx)
	x, y := gx, gy
	dx, dy := 0, -1
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
	for i := 0; i < len(v); i++ {
		if v[i] {
			numVisited++
		}
	}
	println(numVisited)
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
	part1()
	part2()
}
