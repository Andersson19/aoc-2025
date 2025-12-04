package day03

import (
	"strings"

	"github.com/Andersson19/aoc-2025/internal/util"
)

func PartOne(lines []string, extras ...any) any {
	sum := 0
	for _, line := range lines {
		index, largestDigit := findNextLargestDigit(line, 2)
		_, secondLargestDigit := findNextLargestDigit(line[index + 1:], 1)

		var strongestBatterySetup strings.Builder
		strongestBatterySetup.WriteRune(largestDigit)
		strongestBatterySetup.WriteRune(secondLargestDigit)

		sum += util.Atoi(strongestBatterySetup.String())
	}
	return sum
}

func PartTwo(lines []string, extras ...any) any {
	sum := 0

	for _, line := range lines {
		index, largestDigit := findNextLargestDigit(line, 12)

		var strongestBatterySetup strings.Builder
		strongestBatterySetup.WriteRune(largestDigit)
		
		remainingBatteries := line[index + 1:]
		for strongestBatterySetup.Len() < 12 {
			index, nextLargestDigit := findNextLargestDigit(remainingBatteries, 12 - strongestBatterySetup.Len())
			strongestBatterySetup.WriteRune(nextLargestDigit)

			remainingBatteries = remainingBatteries[index + 1:]
		}
		
		sum += util.Atoi(strongestBatterySetup.String())
	}

	return sum
}

func findNextLargestDigit(s string, batterySpotsLeft int) (largestIndex int, largest rune) {
	for i, char := range s {
		if char > largest && i <= (len(s) - batterySpotsLeft) {
			largest = char
			largestIndex = i
		}
	}

	return largestIndex, largest
}
