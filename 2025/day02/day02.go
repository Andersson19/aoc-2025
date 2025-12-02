package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func PartOne(lines []string, extras ...any) any {
	// input is 1 line
	line := lines[0]
	id_ranges := strings.Split(line, ",")

	var ranges_candidates []string
	for _, id_range := range id_ranges {
		start, end, _ := strings.Cut(id_range, "-")
		
		// cannot repeat digits
		if len(start) % 2 == 1 && len(end) % 2 == 1 {
			fmt.Println("No repeating digits for range", start, "-", end)
		} else {
			ranges_candidates = append(ranges_candidates, id_range)
		}
	}

	max := 0
	for _, id_range := range ranges_candidates {
		start, end, _ := strings.Cut(id_range, "-")
		s, _ := strconv.Atoi(start)
		e, _ := strconv.Atoi(end)

		if e - s > max {
			max = e - s
		}
	}

	fmt.Println("Max range size: ", max)
	return 0
}

func PartTwo(lines []string, extras ...any) any {
	_ = lines

	return 0
}
