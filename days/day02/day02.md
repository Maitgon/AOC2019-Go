# Day 01 : The Tyranny of the Rocket Equation

* Difficulty: ⭐

## Description

This day's problem was to implement an IntCode computer. Each step of the IntCode computer consist in an OpCode, and three other integers that depending on the OpCode they mean different things. For the first part we need to run the IntCode program once and check the value at position 0. For the second part we need to find the combination of word and verb that produces the output 19690720.

## Problem solution

This is again another easy problem to solve. I knwo in advance that there are quite a lot of IntCode problems in this year so I will keep the IntCode computer logic in a separated package and keep updating it for each problem. For now, I defined the IntCode computer as a struct with a slice of integers and a pointer to the current position.

```go
type IntCode struct {
    Program []int
    Idx     int
}
```

Then I defined a few methods to be able to run the program and to identify the OpCode and the parameters. The methods are:

```go
func Copy(ic *IntCode) *IntCode {
    Program := make([]int, len(ic.Program))
    copy(Program, ic.Program)
    return &IntCode{
        Program: Program,
        Idx:     ic.Idx,
    }
}

func (ic *IntCode) Step() bool {
    opcode := ic.Program[ic.Idx]
    switch opcode {
    case 1:
        ic.Program[ic.Program[ic.Idx+3]] = ic.Program[ic.Program[ic.Idx+1]] + ic.Program[ic.Program[ic.Idx+2]]
        ic.Idx += 4
    case 2:
        ic.Program[ic.Program[ic.Idx+3]] = ic.Program[ic.Program[ic.Idx+1]] * ic.Program[ic.Program[ic.Idx+2]]
        ic.Idx += 4
    case 99:
        return false
    default:
        panic("Unknown opcode")
    }
    return true
}

func (ic *IntCode) Run() {
    for ic.Step() {
    }
}
```

In the first part we simply need to run the program once after replacing the values at position 1 and 2 with 12 and 2 respectively. Then we just need to check the value at position 0.

```go
func part1(computer *intcode.IntCode) int {
    copy := intcode.Copy(computer)
    copy.Program[1] = 12
    copy.Program[2] = 2
    copy.Run()
    return copy.Program[0]
}
```

For the second part we need to find the combination of noun and verb that produces the output 19690720. I just iterated over all the possible noun and verb combinations as they were only 10000.

```go
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
```

## Opinion

The IntCode computer is a very interesting concept. I'm looking forward to see how it will be used in the next problems. For now, I think it was really cool.
