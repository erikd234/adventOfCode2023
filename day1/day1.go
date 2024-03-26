package day1

import (
	"fmt"
	"strconv"
	"unicode"
)

// Find the first and last number in a line and adds them together.
// if only one number is found then repeat the first number
func Day1(challenge string) int {
	firstNum, lastNum := -1, -1
	sum := 0
	lookingForFirst := true
	for _, char := range challenge {
		if unicode.IsDigit(char) {
			var err error
			if lookingForFirst {
				firstNum, err = strconv.Atoi(string(char))
				if err != nil {
					panic(0)
				}
				lookingForFirst = false
				// reset the last num to -1, if ew get to end
				// and its -1, we will set the value of lastNum ot be first num
				lastNum = -1
			} else {
				lastNum, err = strconv.Atoi(string(char))
				if err != nil {
					panic(0)
				}
			}

		} else if char == '\n' && !lookingForFirst {
			lookingForFirst = true
			if lastNum == -1 {
				lastNum = firstNum
			}
			s := fmt.Sprintf("%d%d", firstNum, lastNum)
			fmt.Println("Calibration number: " + s)
			val, err := strconv.Atoi(s)
			if err != nil {
				panic(0)
			}
			sum += val
		}
	}
	return sum
}
