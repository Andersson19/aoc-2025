package day05

import (
	"slices"

	"github.com/Andersson19/aoc-2025/internal/util"
)

type Range struct {
	start int
	end   int
}

func PartOne(lines []string, extras ...any) any {
	count := 0
	rangeMap := make(map[int]Range)
	doneReadingFreshRanges := false
	for _, line := range lines {

		if line == "" {
			doneReadingFreshRanges = true
			continue
		}

		if doneReadingFreshRanges {
			ingredient := util.Atoi(line)
			// fmt.Println("> Ingredient:", ingredient)

			for key, val := range rangeMap {
				if ingredient < key {
					continue
				}

				if ingredient <= val.end {
					// fmt.Println("Fresh! Fits in range: [", val.start, "-", val.end, "]")
					count += 1
					break
				}
			}

		} else {
			keys := make([]int, len(rangeMap))

			i := 0
			for k := range rangeMap {
				keys[i] = k
				i++
			}
			// process ranges
			start, end := util.CutToInt(line, "-")
			
			if slices.Contains(keys, start) {
				currRange := rangeMap[start]

				// update range with new end value
				if end > currRange.end {
					rangeMap[start] = Range{
						start: currRange.start,
						end: end,
					}
				}
			} else {
				rangeMap[start] = Range{start: start, end: end}
			}
		}
	}

	return count
}

func PartTwo(lines []string, extras ...any) any {
	count := 0
	rangeMap := make(map[int]Range)
	for _, line := range lines {

		if line == "" {
			break
		}

		keys := make([]int, len(rangeMap))

		i := 0
		for k := range rangeMap {
			keys[i] = k
			i++
		}
		// process ranges
		start, end := util.CutToInt(line, "-")
		
		if slices.Contains(keys, start) {
			currRange := rangeMap[start]

			// update range with new end value
			if end > currRange.end {
				rangeMap[start] = Range{
					start: currRange.start,
					end: end,
				}
			}
		} else {
			rangeMap[start] = Range{start: start, end: end}
		}
	}

	// handle overlapping ranges
	done := false
	for !done {
		foundOverlappingRange := false
		for key, value := range rangeMap {
			for otherKey, otherVal := range rangeMap {
				if key == otherKey {
					continue
				}

				if value.start < otherVal.start && value.end > otherVal.end {
					delete(rangeMap, otherKey)
					foundOverlappingRange = true
					break
				}

				if value.end >= otherVal.start && value.end <= otherVal.end {
					foundOverlappingRange = true
					rangeMap[key] = Range{start: min(value.start, otherVal.start), end: otherVal.end}
					delete(rangeMap, otherKey)
					break
				}
			}
			if foundOverlappingRange {
					break
			}
		}
		if !foundOverlappingRange {
			done = true
		}
	}

	for _, v := range rangeMap {
		count += v.end - v.start + 1
	}

	return count
}
