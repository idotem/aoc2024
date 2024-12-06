package main

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	stream, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(stream), "\n")
	var secondNums []int = []int{}
	var firstNums []int = []int{}
	for _, line := range lines {
		num := strings.Fields(line)
		first, err := strconv.Atoi(num[0])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(num[1])
		if err != nil {
			panic(err)
		}
		secondNums = append(secondNums, second)
		firstNums = append(firstNums, first)
	}

	sort.Slice(firstNums, func(i, j int) bool {
		return firstNums[i] < firstNums[j]
	})
	sort.Slice(secondNums, func(i, j int) bool {
		return secondNums[i] < secondNums[j]
	})

	totalVal := 0
	for i, val := range firstNums {
		println(val, secondNums[i])
		totalVal += (int(math.Abs(float64(val - secondNums[i]))))
	}

	print(totalVal)
}
