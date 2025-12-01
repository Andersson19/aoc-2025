package day01

import (
	"strconv"
)

func PartOne(lines []string, extras ...any) any {
	curr := 50
	res := 0
	
	for _, line := range lines {
		direction := line[0]
		i := line[1:]
		
		num, _ := strconv.Atoi(i)
		num = num % 100

		// turn wheel
		if direction == 'L' {
			curr -= num

			if curr < 0 {
				curr = 99 - ((curr * -1) - 1)
			}
		} else {
			curr += num

			if curr > 99 {
				curr = (curr - 100)
			}
		}

		if curr == 0 {
			res += 1
		}
	}

	return res
}

func PartTwo(lines []string, extras ...any) any {
	curr := 50
	res := 0
	
	for _, line := range lines {
		prev := curr
		direction := line[0]
		i := line[1:]
		
		num, _ := strconv.Atoi(i)
		
		if num > 100 {
			whole_wheel_turns := num / 100
			res += whole_wheel_turns
		}

		num = num % 100
		
		passedZero := false
		curr, passedZero = turnWheel(curr, prev, num, rune(direction))
		
		if passedZero {
			res += 1
		}

		// if we land at zero we still count
		if curr == 0 {
			res += 1
		}
	}

	return res
}

// Turns wheel and checks if the turn passed 0 or not
func turnWheel(curr int, prev int, num int, direction rune) (int,bool) {
	if direction == 'L' {
		curr -= num
		
		// If negative, the left turn has passed 0
		if curr < 0 {
			curr = 99 - ((curr * -1) - 1)

			// just make sure we did not start the turn at 0
			return curr, prev != 0
		}

		return curr, false
	} else {
		curr += num
		
		// if over 99, we will also potentially have passed 0
		if curr > 99 {
			curr -= 100

			// only passes zero if we didn't start or end at 0
			return curr, curr != 0 && prev != 0
		}

		return curr, false
	}
}