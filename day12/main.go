package main

import (
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
	"time"
)

type edge struct {
	x, y, d int
}

type point struct {
	x, y int
}

func areaPerimeterAndNumberEdges(x, y int, v byte, field [][]byte, seen map[point]struct{}) (area int64, perimeter int64, nedges int64) {
	seen[point{x, y}] = struct{}{}
	var edges []edge
	for q := []point{{x, y}}; len(q) > 0; area++ {
		p := q[len(q)-1]
		q = q[:len(q)-1]
		for di, d := range []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nx, ny := p.x+d.x, p.y+d.y
			if nx < 0 || ny < 0 || ny >= len(field) || nx >= len(field[ny]) || field[ny][nx] != v {
				perimeter++
				edges = append(edges, edge{p.x, p.y, di})
			} else if _, ok := seen[point{nx, ny}]; !ok {
				seen[point{nx, ny}] = struct{}{}
				q = append(q, point{nx, ny})
			}
		}
	}
	slices.SortFunc(edges, func(a, b edge) int {
		if a.d == 2 || a.d == 3 {
			return cmp.Or(cmp.Compare(a.d, b.d), cmp.Compare(a.x, b.x), cmp.Compare(a.y, b.y))
		} else if a.d == 0 || a.d == 1 {
			return cmp.Or(cmp.Compare(a.d, b.d), cmp.Compare(a.y, b.y), cmp.Compare(a.x, b.x))
		}
		return 0
	})
	lastD, lastX, lastY := -1, -1, -1
	for _, e := range edges {
		if e.d != lastD ||
			(e.d == 0 || e.d == 1) && (e.y != lastY || e.x != lastX+1) ||
			(e.d == 2 || e.d == 3) && (e.x != lastX || e.y != lastY+1) {
			nedges++
		}
		lastX, lastY, lastD = e.x, e.y, e.d
	}
	return area, perimeter, nedges
}

func main() {
	b, _ := os.ReadFile("real.txt")
	lines := bytes.Split(b, []byte("\n"))
	seen := make(map[point]struct{})
	totalPrice1, totalPrice2 := int64(0), int64(0)
	time1 := time.Now()
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if _, ok := seen[point{x, y}]; ok {
				continue
			}
			area, perimeter, edges := areaPerimeterAndNumberEdges(x, y, lines[y][x], lines, seen)
			totalPrice1, totalPrice2 = totalPrice1+area*perimeter, totalPrice2+area*edges
		}
	}
	time2 := time.Now()
	fmt.Printf("Total time (1+2): %v\n", time2.Sub(time1))
	println(totalPrice1)
	println(totalPrice2)
}
