package main

import (
	"bufio"
	"os"
)

func extrude(minM int, maxM int, f [2]int, dx int, dy int, field [][]byte, uniquePoints map[[2]int]struct{}) {
	for m, x0, y0 := minM, f[0]+minM*dx, f[1]+minM*dy; m <= maxM && x0 >= 0 && x0 < len(field[0]) && y0 >= 0 && y0 < len(field); m, x0, y0 = m+1, x0+dx, y0+dy {
		uniquePoints[[2]int{x0, y0}] = struct{}{}
	}
}

func computeUniqueAntinodes(freqs map[byte][][2]int, field [][]byte, minM, maxM int) map[[2]int]struct{} {
	uniquePoints := make(map[[2]int]struct{})
	for _, fp := range freqs {
		for fi := 0; fi < len(fp); fi++ {
			for fj := fi + 1; fj < len(fp); fj++ {
				extrude(minM, maxM, fp[fi], fp[fi][0]-fp[fj][0], fp[fi][1]-fp[fj][1], field, uniquePoints)
				extrude(minM, maxM, fp[fj], fp[fj][0]-fp[fi][0], fp[fj][1]-fp[fi][1], field, uniquePoints)
			}
		}
	}
	return uniquePoints
}

func main() {
	file, _ := os.Open("real.txt")
	scanner, field, freqs := bufio.NewScanner(file), make([][]byte, 0), make(map[byte][][2]int)
	for scanner.Scan() {
		for i, c := range scanner.Bytes() {
			if c == '.' {
				continue
			}
			freqs[c] = append(freqs[c], [2]int{i, len(field)})
		}
		field = append(field, scanner.Bytes())
	}
	println(len(computeUniqueAntinodes(freqs, field, 1, 1)))
	println(len(computeUniqueAntinodes(freqs, field, 0, 100000)))
}
