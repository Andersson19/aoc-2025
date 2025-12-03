package day02_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/Andersson19/aoc-2025/2025/day02"
	"github.com/Andersson19/aoc-2025/internal/test"
)

//go:embed input.txt
var input string

func TestDay(t *testing.T) {
	exampleInput := strings.Split("11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124", ",")

	realInput := strings.Split(input[0:len(input)-1], ",")

	tests := []test.Test{
		{
			Name:     "part 1 example",
			DayFunc:  day.PartOne,
			Input:    exampleInput,
			Expected: 1227775554,
		},
		{
			Name:     "part 1 real",
			DayFunc:  day.PartOne,
			Input:    realInput,
			Expected: 17077011375,
		},
		{
			Name:     "part 2 example",
			DayFunc:  day.PartTwo,
			Input:    exampleInput,
			Expected: 0,
		},
		{
			Name:     "part 2 real",
			DayFunc:  day.PartTwo,
			Input:    realInput,
			Expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := test.DayFunc(test.Input, test.Extras...)
			if actual != test.Expected {
				t.Errorf("Expected %d, actual %d", test.Expected, actual)
			}
		})
	}
}
