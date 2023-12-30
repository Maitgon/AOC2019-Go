# Day 01 : The Tyranny of the Rocket Equation

* Difficulty: ⭐

## Description

First day was very simple. The problem required to read a series of numbers from a file and apply some sort of formula to them. For part 1, the formula was to divide by 3, round down, and subtract 2. For part 2, the formula was to divide by 3, round down, subtract 2, and repeat until the result is less than or equal to 0. The sum of all the results is the answer.

## Problem solution

This is a pretty straightforwards problem. The input is very easy to parse, and the formula is very easy to implement. I just parsed the numbers into a slice of ints.

```go
func parseInput(file []byte) Input {
    inputAux := strings.Split(string(file), "\n")

    input := make(Input, len(inputAux))

    for i, v := range inputAux {
        input[i], _ = strconv.Atoi(v)
    }

    return input
}
```

Then, for each part I just iterated over the slice and apply the corresponding part formula to each number.

```go
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
```

This solution runs in about 40 &mu;s on my machine. I can think of a few ways to improve it, first would be to use multiple threads to calculate the fuel for each number. Another idea is, calculate the sum of the numbers while they are beeing read instead of storing them in a slice and then interating over them again. Lastly, I'm not sure if there is a closed form to calculate the second part getFuel function, but if there is, that would be a good imprevement. But after all, I think 40 &mu;s is good enough and I don't really think that it's worth it to improve it further.

## Opinion

I think this was a good first problem. It was very simple and easy to implement, I can't say much more about it.
