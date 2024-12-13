package main

import (
	"bufio"
	"math/big"
	"os"
	"regexp"
	"strconv"
)

func solve2x2(a00, a01, a10, a11, b0, b1 uint64) (*big.Rat, *big.Rat) {
	rat := func(i uint64) *big.Rat {
		return big.NewRat(int64(i), 1)
	}
	a, b, c, d, e, f := rat(a00), rat(a01), rat(a10), rat(a11), rat(b0), rat(b1)
	m := new(big.Rat).Quo(c, a)
	c.Sub(c, new(big.Rat).Mul(a, m))
	d.Sub(d, new(big.Rat).Mul(b, m))
	f.Sub(f, new(big.Rat).Mul(e, m))
	y := new(big.Rat).Quo(f, d)
	x := new(big.Rat).Quo(new(big.Rat).Sub(e, new(big.Rat).Mul(b, y)), a)
	return x, y
}

func main() {
	file, _ := os.Open("real.txt")
	scanner := bufio.NewScanner(file)
	lineRegex := regexp.MustCompile(`Button [AB]: X\+(\d+), Y\+(\d+)`)
	priceRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	parseUint64 := func(s string) uint64 {
		i, _ := strconv.ParseUint(s, 10, 64)
		return i
	}
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
