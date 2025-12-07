package day06

import (
	"fmt"
	"strings"

	"github.com/Andersson19/aoc-2025/internal/util"
)

type MathProblem struct {
	numbers   []int
	operation string
}

func PartOne(lines []string, extras ...any) any {

	firstLine := lines[0]
	amountOfProblems := len(strings.Split(strings.ReplaceAll(strings.Trim(firstLine, " "), "  ", " "), " "))
	mathProblems := make([]MathProblem, amountOfProblems)

	for _, mathProblemObj := range mathProblems {
		mathProblemObj.numbers = make([]int, len(lines))
	}

	// just fill this with regex
	mathProblems = fillMathProblems(lines, mathProblems)

	fmt.Println(mathProblems[0].numbers, mathProblems[0].operation)

	return 0
}

func fillMathProblems(lines []string, mathProblems []MathProblem) []MathProblem {
	for _, line := range lines {
		line = strings.ReplaceAll(strings.Trim(line, " "), "  ", " ")

		// instead of this, just regex the line for digits
		// and or '*' and '+'
		elements := strings.Split(line, " ")
		fmt.Println(elements)

		if elements[0] == "*" || elements[0] == "+" {
			for i, operation := range elements {
				mathProblems[i].operation = operation
				fmt.Println(i, mathProblems[i].numbers, mathProblems[i].operation)
			}
		} else {
			for i, number := range elements {
				mathProblems[i].numbers = append(mathProblems[i].numbers, util.Atoi(number))

			}
		}
	}
	return mathProblems

}

func PartTwo(lines []string, extras ...any) any {
	_ = lines

	return 0
}
