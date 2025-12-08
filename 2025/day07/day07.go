package day07

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

func PartTwo(lines []string, extras ...any) any {
	var startY, startX int
	for y, line := range lines {
		for x := range len(line) {
			if string(lines[y][x]) == "S" {
				startY, startX = y, x
			}
		}
	}
	var visitedNodes []Node

	_, timelines := calculateTimelines(lines, startY, startX, visitedNodes, 0)
	return timelines
}

func calculateTimelines(lines []string, currY int, currX int, visited []Node, timelines int) ([]Node, int) {

	// beam has hit end of the road
	if currY == len(lines)-1 {
		timelines += 1
		return visited, timelines
	}

	// beam hits splitter
	if lines[currY][currX] == '^' {
		if currX == 0 {
			// beam can only go right from here
			return calculateTimelines(lines, currY, currX+1, visited, timelines)
		} else if currX == len(lines[0])-1 {
			// beam can only go left from here
			return calculateTimelines(lines, currY, currX-1, visited, timelines)
		} else {
			visited, timelines = calculateTimelines(lines, currY, currX+1, visited, timelines)
			visited, timelines = calculateTimelines(lines, currY, currX-1, visited, timelines)
			return visited, timelines
		}

	}

	return calculateTimelines(lines, currY+1, currX, visited, timelines)
}
