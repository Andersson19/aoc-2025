package day09

import (
	"fmt"
	"strings"

	"github.com/Andersson19/aoc-2025/internal/util"
)

type Coordinate struct {
	x int
	y int
}

func PartOne(lines []string, extras ...any) any {
	var cords []Coordinate 
	for _, line := range lines {
		coordinates := strings.Split(line, ",")
		cords = append(cords, Coordinate{
			x: util.Atoi(coordinates[0]),
			y: util.Atoi(coordinates[1]),
		})
	}

	maxSize := 0
	fmt.Println(len(cords))
	for i := range len(cords) {
		currCord := cords[i]
		for j := range len(cords) {
			if i == j {
				continue
			}
			otherCord := cords[j]
			// fmt.Println("Checking (",currCord.x,",",currCord.y,") with (",otherCord.x,",",otherCord.y,")")


			height := util.Abs(otherCord.y - currCord.y) + 1
			width := util.Abs(otherCord.x - currCord.x) + 1
			size := height * width
			// fmt.Println("Size:",size)
			
			if size > maxSize {
				fmt.Println("Found new max! Height:", height, "Width:", width, "Size:", size)
				maxSize = size
			}


		}
	}

	return maxSize
}

func PartTwo(lines []string, extras ...any) any {
	_ = lines

	return 0
}
