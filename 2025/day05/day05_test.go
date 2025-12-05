package day05_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/Andersson19/aoc-2025/2025/day05"
	"github.com/Andersson19/aoc-2025/internal/test"
)

//go:embed input.txt
var input string

func TestDay(t *testing.T) {
  exampleInput := strings.Split(
    `3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
    "\n",
  )

	realInput := strings.Split(input[0:len(input)-1], "\n")

	tests := []test.Test{
		{
			Name:     "part 1 example",
			DayFunc:  day.PartOne,
			Input:    exampleInput,
			Expected: 3,
		},
		{
			Name:     "part 1 real",
			DayFunc:  day.PartOne,
			Input:    realInput,
			Expected: 789,
		},
		{
			Name:     "part 2 example",
			DayFunc:  day.PartTwo,
			Input:    exampleInput,
			Expected: 14,
		},
		{
			Name:     "part 2 real",
			DayFunc:  day.PartTwo,
			Input:    realInput,
			Expected: 343329651880509,
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
