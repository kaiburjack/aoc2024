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

func calculateOne(acc, v int64, vlen int, operator byte) int64 {
	switch operator {
	case '+':
		return acc + v
	case '*':
		return acc * v
	case '|':
		return acc*int64(math.Pow10(vlen)) + v
	}
	return 0
}

func findValidCombinationRec(alphabet string, numbers []number, desiredResult int64, operators []byte, index int, acc int64) int64 {
	if index == len(operators) {
		if acc == desiredResult {
			return desiredResult
		}
		return 0
	}
	for i := 0; i < len(alphabet); i++ {
		operators[index] = alphabet[i]
		newAcc := calculateOne(acc, numbers[index+1].value, numbers[index+1].len, operators[index])
		if newAcc > desiredResult {
			continue
		}
		if result := findValidCombinationRec(alphabet, numbers, desiredResult, operators, index+1, newAcc); result == desiredResult {
			return result
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
		sumPart1 += findValidCombinationRec("+*", numbers, leftAsInt, operators, 0, numbers[0].value)
		sumPart2 += findValidCombinationRec("+*|", numbers, leftAsInt, operators, 0, numbers[0].value)
	}
	time2 := time.Now()
	println(sumPart1)
	println(sumPart2)
	fmt.Printf("Execution time: %v\n", time2.Sub(time1))
}
