package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput() // []string

	position := map[string]int{}
	for i := range input {
		s := strings.Split(input[i], " ")

		direction := s[0]
		amount, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			position["horizontal"] += amount
			if _, ok := position["aim"]; ok && position["aim"] != 0 {
				position["depth"] += position["aim"]*amount
			}
		case "down":
			position["aim"] += amount
		case "up":
			position["aim"] -= amount
		}
	}
	fmt.Printf("%d\n", position["horizontal"]*position["depth"])
}

func readInput() []string{
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return input
}
