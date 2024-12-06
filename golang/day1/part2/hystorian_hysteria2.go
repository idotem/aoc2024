package main

import (
	"log"
	"os"
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

	totalSum := 0
	for _, val := range firstNums {
		count := 0
		for _, val2 := range secondNums {
			if val == val2 {
				count++
			}
		}
		totalSum += (val * count)
	}

	print(totalSum)
}
