package intcode

import (
	"strconv"
	"strings"
	"testing"
)

func IntSliceToString(vals []int) string {
	stringSlice := make([]string, len(vals))

	for i, v := range vals {
		stringSlice[i] = strconv.Itoa(v)
	}

	return strings.Join(stringSlice, ",")
}

func TestComputer_ExecuteProgram_Day02(t *testing.T) {
	tests := []struct {
		name    string
		program string
		want    string
	}{
		{
			name:    "add operation returns expected state (1 + 1 = 2)",
			program: "1,0,0,0,99", want: "2,0,0,0,99",
		},
		{
			name:    "multiply operation returns expected state (3 * 2 = 6)",
			program: "2,3,0,3,99", want: "2,3,0,6,99",
		},
		{
			name:    "multiply operation returns expected state (99 * 99 = 9801)",
			program: "2,4,4,5,99,0", want: "2,4,4,5,99,9801",
		},
		{
			name:    "multiple operations return expected state",
			program: "1,1,1,4,99,5,6,0,99", want: "30,1,1,4,2,5,6,0,99",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComputer([]byte(tt.program))
			c.ExecuteProgram()

			got := IntSliceToString(c.Memory)

			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

func TestComputer_ExecuteProgram_Day05(t *testing.T) {
	tests := []struct {
		name                 string
		program              string
		input                int
		want                 int
		checkPosition        int
		returnDiagnosticCode bool
	}{
		{
			name:    "returns diagnostic code output equal to the input",
			program: "3,0,4,0,99", input: 55, want: 55, returnDiagnosticCode: true,
		},
		{
			name:    "parameter modes with multiply operation sets value at expected postition (33 * 3 = 99)",
			program: "1002,4,3,4,33", want: 99, checkPosition: 4,
		},
		{
			name:    "parameter modes with add operation sets value at expected postition (100 + -1 = 99)",
			program: "1101,100,-1,4,0", want: 99, checkPosition: 4,
		},
		{
			name:    "equals op postion mode equals expected output",
			program: "3,9,8,9,10,9,4,9,99,-1,8", input: 8, want: 1, returnDiagnosticCode: true,
		},
		{
			name:    "equals op postion mode not equals expected output",
			program: "3,9,8,9,10,9,4,9,99,-1,8", input: 7, want: 0, returnDiagnosticCode: true,
		},
		{
			name:    "less than op postion mode less than expected output",
			program: "3,9,7,9,10,9,4,9,99,-1,8", input: 7, want: 1, returnDiagnosticCode: true,
		},
		{
			name:    "less than op postion mode greater than expected output",
			program: "3,9,7,9,10,9,4,9,99,-1,8", input: 9, want: 0, returnDiagnosticCode: true,
		},
		{
			name:    "equals op postion mode equals expected output",
			program: "3,3,1108,-1,8,3,4,3,99", input: 8, want: 1, returnDiagnosticCode: true,
		},
		{
			name:    "equals op postion mode not equals expected output",
			program: "3,3,1108,-1,8,3,4,3,99", input: 7, want: 0, returnDiagnosticCode: true,
		},
		{
			name:    "less than op immediate mode less than expected output",
			program: "3,3,1107,-1,8,3,4,3,99", input: 7, want: 1, returnDiagnosticCode: true,
		},
		{
			name:    "less than op immediate mode greater than expected output",
			program: "3,3,1107,-1,8,3,4,3,99", input: 9, want: 0, returnDiagnosticCode: true,
		},
		{
			name:    "jump op position mode non-zero input expected output",
			program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", input: 5, want: 1, returnDiagnosticCode: true,
		},
		{
			name:    "jump op position mode zero input expected output",
			program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", input: 0, want: 0, returnDiagnosticCode: true,
		},
		{
			name:    "jump op immediate mode non-zero input expected output",
			program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", input: 5, want: 1, returnDiagnosticCode: true,
		},
		{
			name:    "jump op immediate mode zero input expected output",
			program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", input: 0, want: 0, returnDiagnosticCode: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComputer([]byte(tt.program))
			c.ExecuteProgram(tt.input)
			var got int
			if tt.returnDiagnosticCode {
				got = c.GetDiagnosticCode().(int)
			} else {
				got = c.Memory[tt.checkPosition]
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestComputer_ExecuteProgram_Day05_LargerExample(t *testing.T) {
	program := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"

	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "input less than 8",
			input: 7,
			want:  999,
		},
		{
			name:  "input equal to 8",
			input: 8,
			want:  1000,
		},
		{
			name:  "input greater than 8",
			input: 9,
			want:  1001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComputer([]byte(program))
			c.ExecuteProgram(tt.input)
			got := c.GetDiagnosticCode().(int)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestComputer_ExecuteProgram_Day09_InputOutputsCopyOfItself(t *testing.T) {
	program := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"

	c := NewComputer([]byte(program))
	c.ExecuteProgram()

	outputs := c.GetDiagnosticOutputs()
	strSlice := make([]string, len(outputs))
	for i, o := range outputs {
		strSlice[i] = strconv.Itoa(o)
	}

	got := strings.Join(strSlice, ",")

	if got != program {
		t.Errorf("output does not match input. \ngot:\n%s\nwant:\n%s", got, program)
	}

}

func TestComputer_ExecuteProgram_Day09_InputProduces16DigitNumber(t *testing.T) {
	program := "1102,34915192,34915192,7,4,7,99,0"

	c := NewComputer([]byte(program))
	c.ExecuteProgram()

	out := c.GetDiagnosticCode().(int)
	got := len(strconv.Itoa(out))

	if got != 16 {
		t.Errorf("output: %d is not 16 digits", out)
	}
}

func TestComputer_ExecuteProgram_Day09_OutputsLargeNumberFromInput(t *testing.T) {
	program := "104,1125899906842624,99"
	want := 1125899906842624

	c := NewComputer([]byte(program))
	c.ExecuteProgram()

	got := c.GetDiagnosticCode().(int)

	if got != want {
		t.Errorf("unexpected output. got: %d, want: %d", got, want)
	}
}
