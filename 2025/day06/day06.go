package day06

import (
	"strings"

	"github.com/Andersson19/aoc-2025/internal/util"
)

type MathProblem struct {
	numbers   []int
	operation string
}

func PartOne(lines []string, extras ...any) any {
	firstLine := strings.Fields(lines[0])

	// allocate for slice of MathProblems
	amountOfProblems := len(firstLine)
	mathProblems := make([]MathProblem, amountOfProblems)

	// fill from input
	mathProblems = fillMathProblems(lines, mathProblems)

	// calc sum for problem
	sum := 0
	for _, mathProblem := range mathProblems {
		if mathProblem.operation == "+" {
			res := 0
			for _, num := range mathProblem.numbers {
				res += num
			}
			sum += res
		} else {
			res := 1
			for _, num := range mathProblem.numbers {
				res *= num
			}
			sum += res
		}	
	}
	return sum
}

func fillMathProblems(lines []string, mathProblems []MathProblem) []MathProblem {
	for _, line := range lines {
		elements := strings.Fields(line)

		if elements[0] == "*" || elements[0] == "+" {
			for i, operation := range elements {
				mathProblems[i].operation = operation
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
	return calculateCephalopodMath(lines)
}

func calculateCephalopodMath(lines []string) int {
	sum := 0

	lineLength := len(lines[0])

	var mathProblemObj MathProblem
	
	numberOfSpaces := 0
	var numStr strings.Builder
	for x := range lineLength {
		
		numberOfSpaces = 0
		for y := range len(lines) {
			elem := string(lines[y][x])
			if elem == "+" || elem == "*" {
				mathProblemObj.operation = elem
			} else if elem != " " {
				numStr.WriteString(elem)
			} else {
				numberOfSpaces++
			}
		}
		if numStr.Len() != 0 {
			mathProblemObj.numbers = append(mathProblemObj.numbers, util.Atoi(numStr.String()))
			numStr.Reset()
		}

		if numberOfSpaces == len(lines) || x == lineLength - 1 {
			switch mathProblemObj.operation {
				case "+":
					res := 0
					for _, num := range mathProblemObj.numbers {
						res += num
					}
					sum += res
				case "*":
					res := 1
					for _, num := range mathProblemObj.numbers {
						res *= num
					}
					sum += res
			default:
				panic("Unknown operation")
			}

			// reset
			numberOfSpaces = 0
			mathProblemObj.numbers = nil
			mathProblemObj.operation = ""
		}
	}
	return sum
}
