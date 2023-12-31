package day02

import (
	"AOC2019-Go/intcode"
	"os"
	"strconv"
	"strings"
)

type Input = []int

func Solve() (interface{}, interface{}) {
	file, err := os.ReadFile("inputs/day02.txt")
	if err != nil {
		panic(err)
	}

	input := parseInput(file)
	computer := intcode.NewIntCode(input)

	sol1 := part1(computer)
	sol2 := part2(computer)

	return sol1, sol2
}

func part1(computer *intcode.IntCode) int {
	copy := intcode.Copy(computer)
	copy.Program[1] = 12
	copy.Program[2] = 2
	copy.Run()
	return copy.Program[0]
}

func part2(computer *intcode.IntCode) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy := intcode.Copy(computer)
			copy.Program[1] = noun
			copy.Program[2] = verb
			copy.Run()
			if copy.Program[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	panic("No solution found")
}

func parseInput(file []byte) Input {
	inputAux := strings.Split(strings.TrimSpace(string(file)), ",")

	input := make(Input, len(inputAux))

	for i, v := range inputAux {
		input[i], _ = strconv.Atoi(v)
	}

	return input
}
