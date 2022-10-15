package intcode

import (
	"testing"
)

func TestComputer_ExecuteProgram(t *testing.T) {
	tests := []struct {
		program              string
		args                 []int
		want                 int
		checkPosition        int
		returnDiagnosticCode bool
	}{
		{program: "1,0,0,0,99", want: 2, checkPosition: 0},
		{program: "2,3,0,3,99", want: 2, checkPosition: 0},
		{program: "2,4,4,5,99,0", want: 2, checkPosition: 0},
		{program: "1,1,1,4,99,5,6,0,99", want: 30},
		{program: "3,0,4,0,99", args: []int{55}, want: 55, returnDiagnosticCode: true},
		{program: "3,0,4,0,99", args: []int{2}, want: 2, returnDiagnosticCode: true},
		{program: "3,0,4,0,99", args: []int{123}, want: 123, returnDiagnosticCode: true},
		{program: "1002,4,3,4,33", want: 99, checkPosition: 4},
		{program: "1101,100,-1,4,0", want: 99, checkPosition: 4},
	}
	for _, tt := range tests {
		t.Run(tt.program, func(t *testing.T) {
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
