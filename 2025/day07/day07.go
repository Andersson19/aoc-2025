package day07

import "fmt"

type Node struct {
	x int
	y int
}

func PartOne(lines []string, extras ...any) any {
	var startY, startX int
	for y, line := range lines {
		for x := range len(line) {
			if string(lines[y][x]) == "S" {
				startY, startX = y, x
			}
		}
	}
	var visitedNodes []Node

	_, splits := calculateBeamSplits(lines, startY, startX, visitedNodes, 0)
	return splits
}

func IsVisited(visited []Node, y int, x int) bool {
	for _, node := range visited {
		if node.x == x && node.y == y {
			return true
		}
	}
	return false
}

func calculateBeamSplits(lines []string, currY int, currX int, visited []Node, splits int) ([]Node, int) {

	// beam has hit end of the road
	if currY == len(lines)-1 {
		if !IsVisited(visited, currY, currX) {
			visited = append(visited, Node{
				x: currX, y: currY,
			})
		}
		return visited, splits
	}

	// beam hits splitter
	if lines[currY][currX] == '^' {
		if currX == 0 {
			// beam can only go right from here
			return calculateBeamSplits(lines, currY, currX+1, visited, splits+1)
		} else if currX == len(lines[0])-1 {
			// beam can only go left from here
			return calculateBeamSplits(lines, currY, currX-1, visited, splits+1)
		} else {
			splits += 1
			visited, splits = calculateBeamSplits(lines, currY, currX+1, visited, splits)
			visited, splits = calculateBeamSplits(lines, currY, currX-1, visited, splits)
			return visited, splits
		}

	}

	// beam neither hit splitter or is done, needs to keep going
	if !IsVisited(visited, currY, currX) {
		visited = append(visited, Node{
			x: currX, y: currY,
		})
		return calculateBeamSplits(lines, currY+1, currX, visited, splits)
	} else {
		// Node has already been visited
		return visited, splits
	}

}

type NewNode struct {
	x int
	y int
	visitedCounter int
}

func PartTwo(lines []string, extras ...any) any {
	var startY, startX int
	Outer:
		for y, line := range lines {
			for x := range len(line) {
				if string(lines[y][x]) == "S" {
					startY, startX = y, x
					break Outer
				}
			}
		}
	var visitedNodes []NewNode

	visited := calculateTimelines(lines, startY, startX, visitedNodes)

	sum := 0
	for _, node := range visited {
		if node.y == len(lines)-1 {
			sum += node.visitedCounter
		}
	}
	return sum
}

func NewIsVisited(visited []NewNode, y int, x int) bool {
	for _, node := range visited {
		if node.x == x && node.y == y {
			return true
		}
	}
	return false
}

func PrintGrid(lines []string, currY int, currX int, visited []Node) {
	for y, line := range lines {
		for x := range len(line) {
			if currY == y && currX == x {
				fmt.Print("#")
			} else if IsVisited(visited, y, x) {
				fmt.Print("|")
			} else {
				fmt.Print(string(lines[y][x]))
			}
		}
		fmt.Println()
	}
}

func calculateTimelines(lines []string, currY int, currX int, visited []NewNode) ([]NewNode) {
	// PrintGrid(lines, currY, currX, visited)

	// check if this path has already been run
	if !NewIsVisited(visited, currY, currX) {
		visited = append(visited, NewNode{
			y: currY,
			x: currX,
			visitedCounter: 1,
		})
	} else {
		for _, node := range visited {
			if node.x == currX && node.y == currY {
				node.visitedCounter += 1
				break
			}
		}
	}
	

	// beam has hit end of the road
	if currY == len(lines)-1 {
		return visited
	}

	// beam hits splitter
	if lines[currY][currX] == '^' {
		visited = calculateTimelines(lines, currY, currX+1, visited)
		visited = calculateTimelines(lines, currY, currX-1, visited)
		return visited

	}

	// keep going down otherwise
	return calculateTimelines(lines, currY+1, currX, visited)
}
