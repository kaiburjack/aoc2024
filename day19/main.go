package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type node struct {
	next     [256]*node
	terminal bool
}

func buildPrefixTree(towels []string) *node {
	root := &node{}
	for _, towel := range towels {
		c := root
		for i := 0; i < len(towel); i++ {
			if c.next[towel[i]] == nil {
				c.next[towel[i]] = &node{}
			}
			c = c.next[towel[i]]
		}
		c.terminal = true
	}
	return root
}

func checkDesign(design string, root *node, cache map[string]uint64) uint64 {
	if v, ok := cache[design]; ok {
		return v
	}
	if len(design) == 0 {
		return 1
	}
	c := root
	numPossible := uint64(0)
	for i := 0; i < len(design); i++ {
		c = c.next[design[i]]
		if c == nil {
			break
		}
		if c.terminal {
			numPossible += checkDesign(design[i+1:], root, cache)
		}
	}
	cache[design] = numPossible
	return numPossible
}

func input() ([]string, []string) {
	file, _ := os.Open("real.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	towels := strings.Split(scanner.Text(), ", ")
	var designs []string
	scanner.Scan()
	for scanner.Scan() {
		designs = append(designs, scanner.Text())
	}
	return towels, designs
}

func main() {
	towels, designs := input()
	root := buildPrefixTree(towels)
	var numPossible, numCombinations uint64
	cache := make(map[string]uint64)
	time1 := time.Now()
	for _, design := range designs {
		if n := checkDesign(design, root, cache); n > 0 {
			numPossible++
			numCombinations += n
		}
	}
	time2 := time.Now()
	fmt.Printf("Number of possible designs: %d\n", numPossible)
	fmt.Printf("Number of possible combinations: %d\n", numCombinations)
	fmt.Printf("Time taken: %v\n", time2.Sub(time1))
}
