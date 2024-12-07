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

func findValidCombinations(alphabet string, numbers []number, leftAsInt int64) int64 {
	operators := make([]byte, len(numbers)-1)
	n := int64(len(alphabet))
	np := powInt(n, len(numbers)-1)
	for i := int64(0); i < np; i++ {
		for j, den := 0, int64(1); j < len(operators); j, den = j+1, den*n {
			operators[j] = alphabet[(i / den % n)]
		}
		result := calculate(numbers, operators)
		if result == leftAsInt {
			return leftAsInt
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
		line := scanner.Text()
		splitted := strings.Split(line, ": ")
		leftAsInt, _ := strconv.ParseInt(splitted[0], 10, 64)
		rightSplitted := strings.Split(splitted[1], " ")
		numbers := make([]number, 0)
		for _, n := range rightSplitted {
			num, _ := strconv.ParseInt(n, 10, 64)
			numbers = append(numbers, number{num, len(n)})
		}
		sumPart1 += findValidCombinations("+*", numbers, leftAsInt)
		sumPart2 += findValidCombinations("+*|", numbers, leftAsInt)
	}
	time2 := time.Now()
	println(sumPart1)
	println(sumPart2)
	fmt.Printf("Execution time: %v\n", time2.Sub(time1))
}
