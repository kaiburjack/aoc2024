package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("real.txt")
	scanner := bufio.NewScanner(file)
	before := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		firstInt, _ := strconv.Atoi(parts[0])
		secondInt, _ := strconv.Atoi(parts[1])
		if _, ok := before[firstInt]; ok {
			before[firstInt] = append(before[firstInt], secondInt)
		} else {
			before[firstInt] = []int{secondInt}
		}
	}

	sumOfMiddlesPart1 := 0
	sumOfMiddlesPart2 := 0
	for scanner.Scan() {
		numbers := make([]int, 0)
		splitted := strings.Split(scanner.Text(), ",")
		for i := 0; i < len(splitted); i++ {
			number, _ := strconv.Atoi(splitted[i])
			numbers = append(numbers, number)
		}
		numbersCopy := make([]int, len(numbers))
		copy(numbersCopy, numbers)
		slices.SortStableFunc(numbers, func(a, b int) int {
			if oa, ok := before[a]; ok {
				if slices.Contains(oa, b) {
					return -1
				}
			}
			if ob, ok := before[b]; ok {
				if slices.Contains(ob, a) {
					return 1
				}
			}
			return 0
		})

		middle := len(numbers) / 2
		if slices.Equal(numbers, numbersCopy) {
			sumOfMiddlesPart1 += numbers[middle]
		} else {
			sumOfMiddlesPart2 += numbers[middle]
		}
	}
	println(sumOfMiddlesPart1)
	println(sumOfMiddlesPart2)
}
