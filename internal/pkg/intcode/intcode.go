package intcode

import (
	"log"
	"strconv"
	"strings"
)

type Opcode int

const (
	_ Opcode = iota
	Add
	Multiply
	Input
	Output
	JumpIfTrue
	JumpIfFalse
	LessThan
	Equals
	Halt = Opcode(99)
)

var opcodeParameterCounts map[Opcode]int = map[Opcode]int{
	Add:         3,
	Multiply:    3,
	Input:       1,
	Output:      1,
	JumpIfTrue:  2,
	JumpIfFalse: 2,
	LessThan:    3,
	Equals:      3,
}

func opcodeLength(n Opcode) (length int) {
	for n > 0 {
		n = n / 10
		length++
	}

	return length
}

type Computer struct {
	program            []int
	instructionPointer int
	Memory             []int
	idInput            int
	diagnosticCode     interface{}
}

func (c *Computer) Reset() {
	c.Memory = make([]int, len(c.program))
	copy(c.Memory, c.program)
	c.instructionPointer = 0
}

func (c *Computer) GetDiagnosticCode() interface{} {
	return c.diagnosticCode
}

func (c *Computer) evaluateParam(mode int, offset int) int {
	pos := c.instructionPointer + offset
	if mode == 0 {
		return c.Memory[pos]
	} else {
		return pos
	}
}

func (c *Computer) ExecuteProgram(args ...int) {
	if len(args) == 2 {
		c.Memory[1] = args[0] // noun
		c.Memory[2] = args[1] // verb
	} else if len(args) == 1 {
		c.idInput = args[0]
	}

out:
	for {
		opcode := Opcode(c.Memory[c.instructionPointer])
		if opcode == Halt {
			break
		}

		modes := make([]int, 3)
		if opcodeLength(opcode) > 1 {
			strOpcode := strconv.Itoa(int(opcode))
			oc, _ := strconv.Atoi(string(strOpcode[len(strOpcode)-1]))
			opcode = Opcode(oc)
			var modePos int = 0
			for i := len(strOpcode) - 3; i >= 0; i-- {
				mode, _ := strconv.Atoi(string(strOpcode[i]))
				modes[modePos] = mode
				modePos++
			}
		}

		numParams := opcodeParameterCounts[opcode]
		parameters := make([]int, numParams)
		for i := 0; i < numParams; i++ {
			parameters[i] = c.evaluateParam(modes[i], i+1)
		}

		switch opcode {
		case Add:
			c.Memory[parameters[2]] = c.Memory[parameters[0]] + c.Memory[parameters[1]]
			c.instructionPointer += opcodeParameterCounts[opcode] + 1
		case Multiply:
			c.Memory[parameters[2]] = c.Memory[parameters[0]] * c.Memory[parameters[1]]
			c.instructionPointer += opcodeParameterCounts[opcode] + 1
		case Input:
			c.Memory[parameters[0]] = c.idInput
			c.instructionPointer += opcodeParameterCounts[opcode] + 1
		case Output:
			c.diagnosticCode = c.Memory[parameters[0]]
			c.instructionPointer += opcodeParameterCounts[opcode] + 1
		case JumpIfTrue:
			if c.Memory[parameters[0]] != 0 {
				c.instructionPointer = c.Memory[parameters[1]]
			} else {
				c.instructionPointer += opcodeParameterCounts[opcode] + 1
			}
		case JumpIfFalse:
			if c.Memory[parameters[0]] == 0 {
				c.instructionPointer = c.Memory[parameters[1]]
			} else {
				c.instructionPointer += opcodeParameterCounts[opcode] + 1
			}
		case LessThan:
			if c.Memory[parameters[0]] < c.Memory[parameters[1]] {
				c.Memory[parameters[2]] = 1
			} else {
				c.Memory[parameters[2]] = 0
			}
			c.instructionPointer += opcodeParameterCounts[opcode] + 1
		case Equals:
			if c.Memory[parameters[0]] == c.Memory[parameters[1]] {
				c.Memory[parameters[2]] = 1
			} else {
				c.Memory[parameters[2]] = 0
			}
			c.instructionPointer += opcodeParameterCounts[opcode] + 1
		default:
			log.Printf("unexpected opcode: %v", opcode)
			break out
		}
	}
}

func NewComputer(program []byte) *Computer {
	programStrs := strings.Split(string(program), ",")
	programInts := make([]int, len(programStrs))

	for i, s := range programStrs {
		programInts[i], _ = strconv.Atoi(s)
	}
	c := new(Computer)
	c.program = programInts
	c.idInput = -1
	c.Reset()

	return c
}
