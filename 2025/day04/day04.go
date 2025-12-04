package day04

func PartOne(lines []string, extras ...any) any {
	grid := createToiletGrid(lines)
	availableSpots, _ := getRollsToRemove(grid)

	return availableSpots
}

func createToiletGrid(lines []string) [][]rune {
	newLines := make([][]rune, len(lines))

	for i, line := range lines {
		newLines[i] = make([]rune, len(line))
		for j, char := range line {
			newLines[i][j] = char
		}
	}

	return newLines
}

type ToiletRollCords struct {
	x int
	y int
}

func getRollsToRemove(grid [][]rune) (int, []ToiletRollCords) {
	toiletRollsToRemove := 0
	var toiletRollCords []ToiletRollCords
	for y, line := range grid {
		for x, char := range line {
			count := 0

			if char == '.' {
				continue
			}

			if y > 0 {
				if x > 0 && grid[y-1][x-1] == '@' {
					count += 1
				}
				if grid[y-1][x] == '@' {
					count += 1
				}
				if x < len(line)-1 && grid[y-1][x+1] == '@' {
					count += 1
				}
			}

			if y < len(grid)-1 {
				if x > 0 && grid[y+1][x-1] == '@' {
					count += 1
				}
				if grid[y+1][x] == '@' {
					count += 1
				}
				if x < len(line)-1 && grid[y+1][x+1] == '@' {
					count += 1
				}
			}

			if x > 0 && grid[y][x-1] == '@' {
				count += 1
			}
			if x < len(line)-1 && grid[y][x+1] == '@' {
				count += 1
			}

			if count < 4 {
				toiletRollCords = append(toiletRollCords, ToiletRollCords{x: x, y: y})
				toiletRollsToRemove += 1
			}
		}
	}
	return toiletRollsToRemove, toiletRollCords
}

func PartTwo(lines []string, extras ...any) any {
	grid := createToiletGrid(lines)

	sum := 0
	availableToiletRolls := 1
	var cords []ToiletRollCords
	for availableToiletRolls > 0 {
		availableToiletRolls, cords = getRollsToRemove(grid)

		for _, cord := range cords {
			grid[cord.y][cord.x] = '.'
			sum += 1
		}
	}

	return sum
}
