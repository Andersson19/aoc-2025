package day10

import (
	"fmt"
	"strings"

	"github.com/Andersson19/aoc-2025/internal/util"
)

type Machine struct {
	diagram []bool
	buttons [][]int
	joltage []int
}

type Queue [][]bool

func PartOne(lines []string, extras ...any) any {
	machines := createMachines(lines)

	fewestPresses := 0
	// for _, machine := range machines {
	// 	fmt.Println("=================================")
	// 	fmt.Println("Wanted diagram:", printDiagram(machine.diagram))
	// 	fmt.Println("buttons:", machine.buttons)
	// 	fmt.Println("joltage:", machine.joltage)
	// 	fewestPresses += findFewestPresses(machine)
	// }

	fmt.Println("=================================")
	fmt.Println("Wanted diagram:", printDiagram(machines[0].diagram))
	fmt.Println("buttons:", machines[0].buttons)
	fmt.Println("joltage:", machines[0].joltage)
	fewestPresses += findFewestPresses(machines[0])

	return fewestPresses
}

func PartTwo(lines []string, extras ...any) any {
	_ = lines

	return 0
}

func createMachines(lines []string) []Machine {
	var machines []Machine
	for _, line := range lines {
		split := strings.Split(line, " ")
		diagram := strings.Trim(split[0], "[]")
		buttons := split[1:len(split)-1]
		joltage := strings.Split(strings.Trim(split[len(split)-1], "{}"), ",")

		diagramObj := make([]bool, len(diagram))
		for i, d := range diagram {
			if d == '#' {
				diagramObj[i] = true
			} else {
				diagramObj[i] = false
			}
		}

		buttonsObj := make([][]int, len(buttons))
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

func findFewestPresses(machine Machine) int {
	startDiagram := createStartNode(len(machine.diagram))
	var queue Queue
	var nextDiagram []bool

	//init queue
	queue = append(queue, startDiagram)

	pressCounter := 0
	for {
		var nextQueue Queue
		isAdded := false
		fmt.Println()
		for _, diagram := range queue {
			if util.AreEqual(diagram, machine.diagram) {
				return pressCounter
			}

			for _, button := range machine.buttons {
				nextDiagram = pressButton(diagram, button)

				if len(nextQueue) > 0 {
					for _, q := range nextQueue {
						if util.AreEqual(nextDiagram, q) {
							isAdded = true
							break
						}
					}
				}
				if !isAdded {
					nextQueue = append(nextQueue, nextDiagram)
					isAdded = false
				}

			}
		}
		queue = nextQueue
		pressCounter += 1

		if pressCounter > 3 {
			return pressCounter
		}
	}
}

func createStartNode(l int) []bool {
	list := make([]bool, l)
	for i := range list {
		list[i] = false
	}
	return list
}

func pressButton(diagram []bool, buttons []int) []bool {
	fmt.Println(">", printDiagram(diagram))
	fmt.Println(">", buttons)
	
	d := make([]bool, len(diagram))
	copy(diagram, d)


	fmt.Println("-- before:", printDiagram(d))
	for _, lightIndex := range buttons {
		d[lightIndex] = !d[lightIndex]
	}
	fmt.Println("-- after:", printDiagram(d))
	return d
}

func printDiagram(diagram []bool) string {
	var sb strings.Builder
	sb.WriteString("[")
	for _, d := range diagram {
		if d {
			sb.WriteString("#")
		} else {
			sb.WriteString(".")
		}
	}
	sb.WriteString("]")
	return sb.String()
}