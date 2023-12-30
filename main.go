package main

import (
	"AOC2019-Go/days/day01"
	"AOC2019-Go/days/day02"
	"AOC2019-Go/days/day03"
	"AOC2019-Go/days/day04"
	"AOC2019-Go/days/day05"
	"AOC2019-Go/days/day06"
	"AOC2019-Go/days/day07"
	"AOC2019-Go/days/day08"
	"AOC2019-Go/days/day09"
	"AOC2019-Go/days/day10"
	"AOC2019-Go/days/day11"
	"AOC2019-Go/days/day12"
	"AOC2019-Go/days/day13"
	"AOC2019-Go/days/day14"
	"AOC2019-Go/days/day15"
	"AOC2019-Go/days/day16"
	"AOC2019-Go/days/day17"
	"AOC2019-Go/days/day18"
	"AOC2019-Go/days/day19"
	"AOC2019-Go/days/day20"
	"AOC2019-Go/days/day21"
	"AOC2019-Go/days/day22"
	"AOC2019-Go/days/day23"
	"AOC2019-Go/days/day24"
	"AOC2019-Go/days/day25"
	"fmt"
	"os"
	"time"
)

func main() {
	arg := os.Args

	if len(arg) <= 1 {
		fmt.Println("Please provide a day to solve")
		os.Exit(1)
	}

	var f func() (interface{}, interface{})

	for _, day := range arg[1:] {
		switch day {
		case "1":
			f = day01.Solve
		case "2":
			f = day02.Solve
		case "3":
			f = day03.Solve
		case "4":
			f = day04.Solve
		case "5":
			f = day05.Solve
		case "6":
			f = day06.Solve
		case "7":
			f = day07.Solve
		case "8":
			f = day08.Solve
		case "9":
			f = day09.Solve
		case "10":
			f = day10.Solve
		case "11":
			f = day11.Solve
		case "12":
			f = day12.Solve
		case "13":
			f = day13.Solve
		case "14":
			f = day14.Solve
		case "15":
			f = day15.Solve
		case "16":
			f = day16.Solve
		case "17":
			f = day17.Solve
		case "18":
			f = day18.Solve
		case "19":
			f = day19.Solve
		case "20":
			f = day20.Solve
		case "21":
			f = day21.Solve
		case "22":
			f = day22.Solve
		case "23":
			f = day23.Solve
		case "24":
			f = day24.Solve
		case "25":
			f = day25.Solve
		default:
			fmt.Println("Please provide a valid day to solve")
			os.Exit(1)
		}
		start := time.Now()
		sol1, sol2 := f()
		end := time.Since(start)
		printSolution(sol1, sol2, end)
	}

}

func printSolution(sol1, sol2 interface{}, time time.Duration) {
	fmt.Printf("Solution 1: %v\n", sol1)
	fmt.Printf("Solution 2: %v\n", sol2)
	fmt.Printf("Time: %v\n", time)
}
