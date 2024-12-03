package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	regex := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	file, _ := os.Open("real.txt")
	scanner := bufio.NewScanner(file)
	sum1, sum2 := int64(0), int64(0)
	enabled := true
	for scanner.Scan() {
		matches := regex.FindAllString(scanner.Text(), -1)
		for _, match := range matches {
			product := compute(match, &enabled)
			sum1 += product
			if enabled {
				sum2 += product
			}
		}
	}
	println(sum1)
	println(sum2)
}

func compute(match string, enabled *bool) int64 {
	if match == "do()" {
		*enabled = true
	} else if match == "don't()" {
		*enabled = false
	} else {
		numbers := strings.Split(match, ",")
		num1, _ := strconv.ParseInt(numbers[0][4:], 10, 64)
		num2, _ := strconv.ParseInt(numbers[1][:len(numbers[1])-1], 10, 64)
		return num1 * num2
	}
	return 0
}
