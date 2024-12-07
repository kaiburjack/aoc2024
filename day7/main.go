package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type number struct {
	value int64
	len   int
}

func calculate(numbers []number, operators []byte) int64 {
	result := numbers[0].value
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case '+':
			result += numbers[i+1].value
		case '*':
			result *= numbers[i+1].value
		case '|':
			result = result*int64(math.Pow10(numbers[i+1].len)) + numbers[i+1].value
		}
	}
	return result
}

func powInt(base int64, exp int) int64 {
	var result int64 = 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func findValidCombination(alphabet string, numbers []number, desiredResult int64) int64 {
	operators := make([]byte, len(numbers)-1)
	n := int64(len(alphabet))
	np := powInt(n, len(numbers)-1)
	for i := int64(0); i < np; i++ {
		for j, den := 0, int64(1); j < len(operators); j, den = j+1, den*n {
			operators[j] = alphabet[(i / den % n)]
		}
		result := calculate(numbers, operators)
		if result == desiredResult {
			return desiredResult
		}
	}
	return 0
}

func main() {
	file, _ := os.Open("real.txt")
	scanner := bufio.NewScanner(file)
	var sumPart1, sumPart2 int64
	time1 := time.Now()
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), ": ")
		leftAsInt, _ := strconv.ParseInt(splitted[0], 10, 64)
		numbers := make([]number, 0)
		for _, n := range strings.Split(splitted[1], " ") {
			num, _ := strconv.ParseInt(n, 10, 64)
			numbers = append(numbers, number{num, len(n)})
		}
		sumPart1 += findValidCombination("+*", numbers, leftAsInt)
		sumPart2 += findValidCombination("+*|", numbers, leftAsInt)
	}
	time2 := time.Now()
	println(sumPart1)
	println(sumPart2)
	fmt.Printf("Execution time: %v\n", time2.Sub(time1))
}
