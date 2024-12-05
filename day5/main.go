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
	before := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		if v, ok := before[parts[0]]; ok {
			before[parts[0]] = append(v, parts[1])
		} else {
			before[parts[0]] = []string{parts[1]}
		}
	}

	sumOfMiddlesPart1 := 0
	sumOfMiddlesPart2 := 0
	for scanner.Scan() {
		numbers := make([]string, 0)
		splitted := strings.Split(scanner.Text(), ",")
		for i := 0; i < len(splitted); i++ {
			numbers = append(numbers, splitted[i])
		}
		numbersCopy := make([]string, len(numbers))
		copy(numbersCopy, numbers)
		slices.SortStableFunc(numbers, func(a, b string) int {
			if oa, ok := before[a]; ok && slices.Contains(oa, b) {
				return -1
			} else if ob, ok := before[b]; ok && slices.Contains(ob, a) {
				return 1
			}
			return 0
		})

		middle := len(numbers) / 2
		middleAsInt, _ := strconv.Atoi(numbers[middle])
		if slices.Equal(numbers, numbersCopy) {
			sumOfMiddlesPart1 += middleAsInt
		} else {
			sumOfMiddlesPart2 += middleAsInt
		}
	}
	println(sumOfMiddlesPart1)
	println(sumOfMiddlesPart2)
}
