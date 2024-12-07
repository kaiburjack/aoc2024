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

func findValidCombinationRec(alphabet string, numbers []number, desiredResult int64, operators []byte, index int, acc int64) int64 {
	if index == len(operators) {
		return acc
	}
	for i := 0; i < len(alphabet); i++ {
		operators[index] = alphabet[i]
		var res int64
		switch operators[index] {
		case '+':
			res = acc + numbers[index+1].value
		case '*':
			res = acc * numbers[index+1].value
		case '|':
			res = acc*int64(math.Pow10(numbers[index+1].len)) + numbers[index+1].value
		}
		if res > desiredResult {
			continue
		}
		if result := findValidCombinationRec(alphabet, numbers, desiredResult, operators, index+1, res); result == desiredResult {
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
