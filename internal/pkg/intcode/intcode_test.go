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
		args    []int
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
			c.ExecuteProgram(tt.args...)

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
		args                 []int
		want                 int
		checkPosition        int
		returnDiagnosticCode bool
	}{
		{
			name:    "returns diagnostic code output equal to the input",
			program: "3,0,4,0,99", args: []int{55}, want: 55, returnDiagnosticCode: true,
		},
		{
			name:    "parameter modes with multiply operation sets value at expected postition (33 * 3 = 99)",
			program: "1002,4,3,4,33", want: 99, checkPosition: 4,
		},
		{
			name:    "parameter modes with add operation sets value at expected postition (100 + -1 = 99)",
			program: "1101,100,-1,4,0", want: 99, checkPosition: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComputer([]byte(tt.program))
			c.ExecuteProgram(tt.args...)
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
