package day01

import (
	"os"
	"strconv"
	"strings"
)

type Input = []int

func Solve() (interface{}, interface{}) {
	// Reading the input
	file, err := os.ReadFile("inputs/day01.txt")
	if err != nil {
		panic(err)
	}

	// Parsing the input
	input := parseInput(file)

	sol1, sol2 := parts(input)

	return sol1, sol2
}

func parts(input Input) (int, int) {
	sol1 := 0
	sol2 := 0
	for _, v := range input {
		sol1 += v/3 - 2
		sol2 += getFuel(v)
	}
	return sol1, sol2
}

func getFuel(mass int) int {
	fuel := mass/3 - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + getFuel(fuel)
}

func parseInput(file []byte) Input {
	inputAux := strings.Split(string(file), "\n")

	input := make(Input, len(inputAux))

	for i, v := range inputAux {
		input[i], _ = strconv.Atoi(v)
	}

	return input
}
