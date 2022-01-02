package main

import "testing"

func Test_calculate_fuel_requirement(t *testing.T) {
	type args struct {
		mass                    int
		include_additional_fuel bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "12-false", args: args{12, false}, want: 2},
		{name: "14-false", args: args{14, false}, want: 2},
		{name: "1969-false", args: args{1969, false}, want: 654},
		{name: "100756-false", args: args{100756, false}, want: 33583},
		{name: "12-true", args: args{12, true}, want: 2},
		{name: "1969-true", args: args{1969, true}, want: 966},
		{name: "100756-true", args: args{100756, true}, want: 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate_fuel_requirement(tt.args.mass, tt.args.include_additional_fuel); got != tt.want {
				t.Errorf("calculate_fuel_requirement() = %v, want %v", got, tt.want)
			}
		})
	}
}
