package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, t := count(readInput())

	fmt.Printf("Increase: %d\n", i)
	fmt.Printf("Window Increase: %d\n", t)
}

func readInput() []int {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		input = append(input, i)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return input
}

func count(input []int) (int, int){
	var numIncrease int
	var numWindowIncrease int

	var sum int
	var prevSum int
	for i := range input {
		if i != len(input)-1 {
			if input[i] < input[i+1] {
				numIncrease++
			}
		}

		if i != len(input)-3 {
			sum = input[i] + input[i+1] + input[i+2]
			if sum > prevSum {
				numWindowIncrease++
			}
		} else {
			break
		}

		prevSum = sum
	}
	return numIncrease, numWindowIncrease
}
