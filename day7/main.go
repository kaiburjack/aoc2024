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

func findValidCombinationRec(alphabet string, numbers []number, desiredResult int64, operators []byte, index int) int64 {
	if index == len(operators) {
		if calculate(numbers, operators) == desiredResult {
			return desiredResult
		}
		return 0
	}
	for i := 0; i < len(alphabet); i++ {
		operators[index] = alphabet[i]
		if findValidCombinationRec(alphabet, numbers, desiredResult, operators, index+1) == desiredResult {
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
		operators := make([]byte, len(numbers)-1)
		sumPart1 += findValidCombinationRec("+*", numbers, leftAsInt, operators, 0)
		sumPart2 += findValidCombinationRec("+*|", numbers, leftAsInt, operators, 0)
	}
	time2 := time.Now()
	println(sumPart1)
	println(sumPart2)
	fmt.Printf("Execution time: %v\n", time2.Sub(time1))
}
