package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PartOne(puzzleInput string) int {
	passwordRange := strings.Split(puzzleInput, "-")
	from, _ := strconv.Atoi(passwordRange[0])
	to, _ := strconv.Atoi(passwordRange[1])

	var validCount int
	for password := from; password <= to; password++ {
		if isValidPassword(strconv.Itoa(password), false) {
			validCount += 1
		}
	}

	return validCount
}

func PartTwo(puzzleInput string) int {
	passwordRange := strings.Split(puzzleInput, "-")
	from, _ := strconv.Atoi(passwordRange[0])
	to, _ := strconv.Atoi(passwordRange[1])

	var validCount int
	for password := from; password <= to; password++ {
		if isValidPassword(strconv.Itoa(password), true) {
			validCount += 1
		}
	}

	return validCount
}

func isValidPassword(password string, excludeGroups bool) bool {
	containsDouble := false
	alwaysIncreases := true
	for i := 0; i < len(password); i++ {
		if i != len(password)-1 && password[i+1] < password[i] {
			alwaysIncreases = false
			break
		}

		if excludeGroups {
			current := string(password[i])
			occurrences := 1
			lastOccurrence := 0
			for j := i + 1; j < len(password); j++ {
				if string(password[j]) == current {
					occurrences += 1
					lastOccurrence = j
				} else {
					break
				}
			}
			if occurrences >= 2 {
				i = lastOccurrence - 1
				if occurrences == 2 {
					containsDouble = true
				}
			}
		} else {
			if i != len(password)-1 && string(password[i]) == string(password[i+1]) {
				containsDouble = true
			}
		}
	}

	return containsDouble && alwaysIncreases
}

func main() {
	puzzleInput := "357253-892942"
	fmt.Printf("Part One: %d\n", PartOne(puzzleInput)) // 530
	fmt.Printf("Part Two: %d\n", PartTwo(puzzleInput)) // 324
}
