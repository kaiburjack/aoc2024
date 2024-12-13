package main

import (
	"bufio"
	"math/big"
	"os"
	"regexp"
	"strconv"
)

func solve2x2(a00, a01, a10, a11, b0, b1 uint64) (*big.Rat, *big.Rat) {
	r := func(i uint64) *big.Rat { return big.NewRat(int64(i), 1) }
	n := func() *big.Rat { return new(big.Rat) }
	a, b, c, d, e, f := r(a00), r(a01), r(a10), r(a11), r(b0), r(b1)
	m := n().Quo(c, a)
	c.Sub(c, n().Mul(a, m))
	d.Sub(d, n().Mul(b, m))
	f.Sub(f, n().Mul(e, m))
	y := n().Quo(f, d)
	x := n().Quo(n().Sub(e, n().Mul(b, y)), a)
	return x, y
}

func main() {
	file, _ := os.Open("real.txt")
	scanner := bufio.NewScanner(file)
	lineRegex := regexp.MustCompile(`Button [AB]: X\+(\d+), Y\+(\d+)`)
	priceRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	parseUint64 := func(s string) uint64 { i, _ := strconv.ParseUint(s, 10, 64); return i }
	var totalTokens [2]uint64
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		a := lineRegex.FindStringSubmatch(scanner.Text())
		aX, aY := parseUint64(a[1]), parseUint64(a[2])
		scanner.Scan()
		b := lineRegex.FindStringSubmatch(scanner.Text())
		bX, bY := parseUint64(b[1]), parseUint64(b[2])
		scanner.Scan()
		priceLine := priceRegex.FindStringSubmatch(scanner.Text())
		priceX, priceY := parseUint64(priceLine[1]), parseUint64(priceLine[2])
		for i, add := 0, uint64(0); i < 2; i, add = i+1, uint64(10000000000000) {
			x, y := solve2x2(aX, bX, aY, bY, priceX+add, priceY+add)
			if x.Denom().Cmp(big.NewInt(1)) == 0 && y.Denom().Cmp(big.NewInt(1)) == 0 {
				totalTokens[i] += 3*x.Num().Uint64() + y.Num().Uint64()
			}
		}
	}
	println(totalTokens[0])
	println(totalTokens[1])
}
