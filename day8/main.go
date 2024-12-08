package main

import (
	"bufio"
	"os"
)

type coord struct {
	x, y int
}

func extrude(minM int, maxM int, f coord, dx int, dy int, field [][]byte, uniquePoints map[coord]struct{}) {
	for m, x0, y0 := minM, f.x+minM*dx, f.y+minM*dy; m <= maxM && x0 >= 0 && x0 < len(field[0]) && y0 >= 0 && y0 < len(field); m, x0, y0 = m+1, x0+dx, y0+dy {
		uniquePoints[coord{x0, y0}] = struct{}{}
	}
}

func computeUniqueAntinodes(freqs map[byte][]coord, field [][]byte, minM, maxM int) map[coord]struct{} {
	uniquePoints := make(map[coord]struct{})
	for _, fp := range freqs {
		for fi := 0; fi < len(fp); fi++ {
			for fj := fi + 1; fj < len(fp); fj++ {
				extrude(minM, maxM, fp[fi], fp[fi].x-fp[fj].x, fp[fi].y-fp[fj].y, field, uniquePoints)
				extrude(minM, maxM, fp[fj], fp[fj].x-fp[fi].x, fp[fj].y-fp[fi].y, field, uniquePoints)
			}
		}
	}
	return uniquePoints
}

func main() {
	file, _ := os.Open("real.txt")
	scanner := bufio.NewScanner(file)
	field := make([][]byte, 0)
	freqs := make(map[byte][]coord)
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c == '.' {
				continue
			}
			freqs[byte(c)] = append(freqs[byte(c)], coord{i, len(field)})
		}
		field = append(field, []byte(line))
	}
	println(len(computeUniqueAntinodes(freqs, field, 1, 1)))
	println(len(computeUniqueAntinodes(freqs, field, 0, 100000)))
}
