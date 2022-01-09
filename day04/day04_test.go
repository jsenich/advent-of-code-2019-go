package main

import "testing"

func Test_isValidPassword(t *testing.T) {
	tests := []struct {
		password      string
		excludeGroups bool
		want          bool
	}{
		{
			password:      "111111",
			excludeGroups: false,
			want:          true,
		},
		{
			password:      "223450",
			excludeGroups: false,
			want:          false,
		},
		{
			password:      "123789",
			excludeGroups: false,
			want:          false,
		},
		{
			password:      "112233",
			excludeGroups: true,
			want:          true,
		},
		{
			password:      "123444",
			excludeGroups: true,
			want:          false,
		},
		{
			password:      "111122",
			excludeGroups: true,
			want:          true,
		},
		{
			password:      "1111222334442",
			excludeGroups: true,
			want:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			if got := isValidPassword(tt.password, tt.excludeGroups); got != tt.want {
				t.Errorf("isValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
