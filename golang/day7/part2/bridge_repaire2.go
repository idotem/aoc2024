package main

import (
	"aoc/util"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

var OPERANDS = [3]string{"+", "*", "||"}

func main() {
	lines := util.ReadFile("../input.txt")
	total := int64(0)
	for _, line := range lines {
		split := strings.Split(line, ":")
		result := parseInt(strings.Trim(split[0], ""))
		nums := strings.Fields(split[1])
		total += buildOperandsAndCalc(nums, result)
	}
	fmt.Println(total)
}

func parseInt(s string) int64 {
	result, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic("RESULT FAILED TO CONVERT FROM STRING TO INT")
	}
	return result
}

func buildOperandsAndCalc(nums []string, result int64) int64 {
	combsLen := int64(len(nums) - 1)
	numOfComb := math.Pow(float64(len(OPERANDS)), float64(combsLen))
	var operands = make([][]string, int64(numOfComb))
	for i := int64(0); i < int64(numOfComb); i++ {
		operands[i] = getCombination(i, combsLen, int64(numOfComb-1))
		calc := getCombinationCalc(nums, operands[i])
		if calc == result {
			return result
		}
	}
	return 0
}

func getCombinationCalc(nums []string, combs []string) int64 {
	j := 0
	prevRes := calculate(parseInt(nums[0]), parseInt(nums[1]), combs[j])
	j++
	for i := 2; i < len(nums); i++ {
		prevRes = calculate(prevRes, parseInt(nums[i]), combs[j])
		j++
	}
	return prevRes
}

func calculate(first int64, second int64, op string) int64 {
	switch op {
	case OPERANDS[0]:
		return first + second
	case OPERANDS[1]:
		return first * second
	case OPERANDS[2]:
		return parseInt(fmt.Sprintf("%d%d", first, second))
	default:
		log.Fatal("NOT * NOR + op: ", op, first, second)
	}
	return 0
}

func getCombination(i, combsLen, numOfComb int64) []string {
	combs := make([]string, combsLen+1)
	combsInBinary := getCombinationValuesInBase3(i, numOfComb)
	for j, c := range combsInBinary {
		combs[j] = OPERANDS[parseInt(c)]
	}
	return combs
}

func getCombinationValuesInBase3(n, maxNumber int64) []string {
	maxBits := len(strconv.FormatInt(int64(maxNumber), len(OPERANDS)))

	base3 := strconv.FormatInt(int64(n), len(OPERANDS))

	paddedBinary := fmt.Sprintf("%0*s", maxBits, base3)

	return strings.Split(paddedBinary, "")
}
