package day11_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/Andersson19/aoc-2025/2025/day11"
	"github.com/Andersson19/aoc-2025/internal/test"
)

//go:embed input.txt
var input string

func TestDay(t *testing.T) {
  exampleInput := strings.Split(
    `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`,
    "\n",
  )

  exampleInputPartTwo := strings.Split(
	`svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`,
	"\n",
  )

	realInput := strings.Split(input[0:len(input)-1], "\n")

	tests := []test.Test{
		{
			Name:     "part 1 example",
			DayFunc:  day.PartOne,
			Input:    exampleInput,
			Expected: 5,
		},
		{
			Name:     "part 1 real",
			DayFunc:  day.PartOne,
			Input:    realInput,
			Expected: 643,
		},
		{
			Name:     "part 2 example",
			DayFunc:  day.PartTwo,
			Input:    exampleInputPartTwo,
			Expected: 2,
		},
		{
			Name:     "part 2 real",
			DayFunc:  day.PartTwo,
			Input:    realInput,
			Expected: 417190406827152,
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
