package day02

import (
	"strconv"
	"strings"

	"github.com/Andersson19/aoc-2025/internal/util"
)

func PartOne(lines []string, extras ...any) any {
	sum := 0

	//lines = strings.Split(lines[0], ",")
	for _, id_range := range lines {
		id_range = strings.Trim(id_range, "\n")
		start, end, _ := strings.Cut(id_range, "-")

		// if start has odd number of digits and end is of same length, we skip
		if !hasEvenAmountOfDigits(start) && len(end) == len(start) {
			continue
		}

		endValue := util.Atoi(end)

		// if start has odd number of digits, find next value that has even number of
		for !hasEvenAmountOfDigits(start) {
			s := util.Atoi(start)
			s += 1
			start = strconv.Itoa(s)
		}
		startValue := util.Atoi(start)

		// if start value is now greater then end value, we skip
		if startValue > endValue {
			continue
		}

		// take first half of start and find constant needed
		firstHalf := start[:len(start)/2]
		_constant := findConstantFromLength(len(firstHalf))

		// calculate first possible incorrect id
		firstHalfVal := util.Atoi(firstHalf)
		incorrectId := firstHalfVal * _constant

		// check all values within range
		for incorrectId <= endValue {
			if !hasEvenAmountOfDigits(strconv.Itoa(incorrectId)) {
				incorrectId += 1
				continue
			}

			if incorrectId >= startValue {
				sum += incorrectId
			}

			firstHalfVal += 1
			incorrectId = firstHalfVal * _constant
		}
	}

	return sum
}

func hasEvenAmountOfDigits(s string) bool {
	return len(s)%2 == 0
}

func findConstantFromLength(length int) int {
	l := length - 1

	constant := "1"

	for range l {
		constant = constant + "0"
	}

	val := util.Atoi(constant + "1")
	return val
}

func PartTwo(lines []string, extras ...any) any {
	_ = lines

	return 0
}
