package day10

/*
	Thanks to:
		- Reddit user: tenthmascot (https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory)
		- Github user: pemoreau (https://github.com/pemoreau/advent-of-code/)
	for the inspiration to this solution
*/

import (
	"math"
	"strings"

	"github.com/Andersson19/aoc-2025/internal/util"
)

type Machine struct {
	diagram Diagram
	buttons []Button
	joltage Joltage
}

type ButtonCombination struct {
	joltage Joltage
	buttonPresses int
}

type Button []int
type Diagram []int
type Joltage []int

func PartOne(lines []string, extras ...any) any {
	machines := createMachines(lines)
	var btnCombs []ButtonCombination

	fewestPresses := 0
	for _, machine := range machines {
		btnCombs = findAllButtonCombinations(machine.buttons, len(machine.diagram))
		fewestPresses += machine.diagram.findFewestPresses(btnCombs)
	}
	return fewestPresses
}

func PartTwo(lines []string, extras ...any) any {
	machines := createMachines(lines)
	fewestPresses := 0

	for _, machine := range machines {
		var fewest int
		btnCombs := findAllButtonCombinations(machine.buttons, len(machine.joltage))
		fewest, _ = machine.joltage.findFewestPresses(btnCombs)
		fewestPresses += fewest
	}

	return fewestPresses
}

func createMachines(lines []string) []Machine {
	var machines []Machine
	for _, line := range lines {
		split := strings.Split(line, " ")
		diagram := strings.Trim(split[0], "[]")
		buttons := split[1:len(split)-1]
		joltage := strings.Split(strings.Trim(split[len(split)-1], "{}"), ",")

		diagramObj := make([]int, len(diagram))
		for i, d := range diagram {
			if d == '#' {
				diagramObj[i] = 1
			} else {
				diagramObj[i] = 0
			}
		}

		buttonsObj := make([]Button, len(buttons))
		for i, b := range buttons {
			bSplit := strings.Split(strings.Trim(b, "()"), ",")

			btnElem := make([]int, len(bSplit))
			for j, n := range bSplit {
				btnElem[j] = util.Atoi(n)
			}
			buttonsObj[i] = btnElem
		}
		
		joltageObj := make([]int, len(joltage))
		for i, j := range joltage {
			joltageObj[i] = util.Atoi(j)
		}

		machines = append(machines, Machine{
			diagram: diagramObj,
			buttons: buttonsObj,
			joltage: joltageObj,
		})
	}
	return machines
}

func findAllButtonCombinations(buttons []Button, diagramSize int) []ButtonCombination {
	possibleCombinations := 1 << len(buttons)
	var joltage []int
	var btnPresses int

	btnCombs := make([]ButtonCombination, 0, possibleCombinations)
	for i := range possibleCombinations {
		joltage = make([]int, diagramSize)
		btnPresses = 0
		
		// press buttons
		for j := range buttons {
			if (i & (1 << j)) != 0 {
				btnPresses++
				for _, idx := range buttons[j] {
					joltage[idx]++
				}
			}
		}
		
		btnCombs = append(btnCombs, ButtonCombination{
			joltage: joltage,
			buttonPresses: btnPresses,
		})
	}
	return btnCombs
}

func (diagram Diagram) findFewestPresses(btnCombs []ButtonCombination) int {
	minBtnPresses := -1
	Inner:
	for _, comb := range btnCombs {
		for i, value := range comb.joltage {
			if diagram[i] != value % 2 {
				continue Inner
			}
		}

		if minBtnPresses == -1 {
			minBtnPresses = comb.buttonPresses
		} else {
			minBtnPresses = min(minBtnPresses, comb.buttonPresses)
		}
	}
	return minBtnPresses
}

func (joltage Joltage) findFewestPresses(btnCombs []ButtonCombination) (int, bool) {
	if joltage.isZero() {
		return 0, true
	}

	fewestPresses := math.MaxInt
	for _, comb := range btnCombs {
		if !comb.joltage.smallerOrEqual(joltage) {
			continue
		}

		if !comb.joltage.equalsModulo2(joltage) {
			continue
		}

		nextJoltage := make(Joltage, len(joltage))
		for i := range nextJoltage {
			nextJoltage[i] = (joltage[i] - comb.joltage[i]) / 2
		}

		presses, ok := nextJoltage.findFewestPresses(btnCombs)
		if !ok {
			continue
		}

		if totalPresses := 2 * presses + comb.buttonPresses; totalPresses < fewestPresses {
			fewestPresses = totalPresses
		}
	}
	if fewestPresses < math.MaxInt {
		return fewestPresses, true
	}
	return 0, false
}

func (joltage Joltage) isZero() bool {
	for _, j := range joltage {
		if j != 0 {
			return false
		}
	}
	return true
}

func (a Joltage) smallerOrEqual(b Joltage) bool {
	for i := range a {
		if a[i] > b[i] {
			return false
		}
	}
	return true
}

func (a Joltage) equalsModulo2(b Joltage) bool {
	for i := range a {
		if a[i] % 2 != b[i] % 2 {
			return false
		} 
	}
	return true
}