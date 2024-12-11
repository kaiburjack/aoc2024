package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var numToDigitCount = [19][2]uint64{
	{1000000000000000000, 19},
	{100000000000000000, 18},
	{10000000000000000, 17},
	{1000000000000000, 16},
	{100000000000000, 15},
	{10000000000000, 14},
	{1000000000000, 13},
	{100000000000, 12},
	{10000000000, 11},
	{1000000000, 10},
	{100000000, 9},
	{10000000, 8},
	{1000000, 7},
	{100000, 6},
	{10000, 5},
	{1000, 4},
	{100, 3},
	{10, 2},
}

var digitCountToNumEven = [19]uint64{
	0,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
	10000000000,
	100000000000,
	1000000000000,
	10000000000000,
	100000000000000,
	1000000000000000,
	10000000000000000,
	100000000000000000,
}

func numDigitsBase10(n uint64) uint64 {
	if n < 10 {
		return 1
	} else if n < 100 {
		return 2
	} else if n < 1000 {
		return 3
	} else if n < 10000 {
		return 4
	} else if n < 100000 {
		return 5
	} else if n < 1000000 {
		return 6
	} else if n < 10000000 {
		return 7
	} else if n < 100000000 {
		return 8
	}
	for _, pair := range numToDigitCount {
		if n >= pair[0] {
			return pair[1]
		}
	}
	return 1
}

func divide(n uint64, numDigits uint64) (uint64, uint64) {
	u := digitCountToNumEven[numDigits>>1]
	return n / u, n % u
}

func SimulateOneStoneNSteps(stone uint64, n uint64, cache map[[2]uint64]uint64) uint64 {
	if v, ok := cache[[2]uint64{stone, n}]; ok {
		return v
	} else if n == 0 {
		return 1
	}
	if stone == 0 {
		count := SimulateOneStoneNSteps(1, n-1, cache)
		return count
	} else {
		numDigits := numDigitsBase10(stone)
		if numDigits&1 == 0 {
			stone0, stone1 := divide(stone, numDigits)
			count := SimulateOneStoneNSteps(stone0, n-1, cache) + SimulateOneStoneNSteps(stone1, n-1, cache)
			cache[[2]uint64{stone, n}] = count
			return count
		} else {
			count := SimulateOneStoneNSteps(stone*2024, n-1, cache)
			return count
		}
	}
}

func main() {
	bytes, _ := os.ReadFile("real.txt")
	numbers := strings.Fields(string(bytes))
	stones := make([]uint64, len(numbers))
	for i, number := range numbers {
		n, _ := strconv.ParseUint(number, 10, 64)
		stones[i] = n
	}
	cache := make(map[[2]uint64]uint64)
	// Part 1
	{
		totalStones := uint64(0)
		time1 := time.Now()
		for _, stone := range stones {
			totalStones += SimulateOneStoneNSteps(stone, 25, cache)
		}
		time2 := time.Now()
		fmt.Print("Part1:\n")
		fmt.Printf("  Total stones: %d\n", totalStones)
		fmt.Printf("  Time: %v\n", time2.Sub(time1))
	}
	// Part 2
	{
		totalStones := uint64(0)
		time1 := time.Now()
		for _, stone := range stones {
			totalStones += SimulateOneStoneNSteps(stone, 75, cache)
		}
		time2 := time.Now()
		fmt.Print("Part2:\n")
		fmt.Printf("  Total stones: %d\n", totalStones)
		fmt.Printf("  Time: %v\n", time2.Sub(time1))
	}
}
