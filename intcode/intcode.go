package intcode

type IntCode struct {
	Program []int
	Idx     int
}

func NewIntCode(Program []int) *IntCode {
	return &IntCode{
		Program: Program,
		Idx:     0,
	}
}

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
