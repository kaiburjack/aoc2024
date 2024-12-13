package main

import (
	"bufio"
	"math/big"
	"os"
	"regexp"
	"strconv"
)

func solve2x2(a00, a01, a10, a11, b0, b1 int64) (*big.Rat, *big.Rat) {
	rat := func(i int64) *big.Rat {
		return big.NewRat(i, 1)
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
	parseInt64 := func(s string) int64 {
		i, _ := strconv.ParseInt(s, 10, 64)
		return i
	}
	var totalTokens [2]int64
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		a := lineRegex.FindStringSubmatch(scanner.Text())
		aX, aY := parseInt64(a[1]), parseInt64(a[2])
		scanner.Scan()
		b := lineRegex.FindStringSubmatch(scanner.Text())
		bX, bY := parseInt64(b[1]), parseInt64(b[2])
		scanner.Scan()
		priceLine := priceRegex.FindStringSubmatch(scanner.Text())
		priceX, priceY := parseInt64(priceLine[1]), parseInt64(priceLine[2])
		for i, add := 0, int64(0); i < 2; i, add = i+1, int64(10000000000000) {
			x, y := solve2x2(aX, bX, aY, bY, priceX+add, priceY+add)
			if x.Denom().Cmp(big.NewInt(1)) == 0 && y.Denom().Cmp(big.NewInt(1)) == 0 {
				totalTokens[i] += 3*x.Num().Int64() + y.Num().Int64()
			}
		}
	}
	println(totalTokens[0])
	println(totalTokens[1])
}
