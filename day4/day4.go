package day4

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Each line of the reader represents a different scratch card
// The inputs on the left side of the seperator | are the winnign number
// the inputs on the right side of the seperator are the actual ones.
// Each match doubles the amount of poitns recieviedc
// For example i reieve 1 then 2 then 4 then 8 points for each match
func Day4(puzzle io.Reader) int {
	scanner := bufio.NewScanner(puzzle)
	sum := 0
	for scanner.Scan() {
		currentMultiplier := 0
		line := scanner.Text()
		_, line, found := strings.Cut(line, ":")
		if !found {
			panic("COULD NOT FIND : IN LINE")
		}
		before, after, found := strings.Cut(line, "|")
		if !found {
			panic("COULD NOT FIND : IN LINE")
		}
		fmt.Printf("%s, %s\n", before, after)
		winners := strings.Split(strings.TrimSpace(before), " ")
		actuals := strings.Split(strings.TrimSpace(after), " ")
		fmt.Println()
		for _, w := range winners {
			for _, a := range actuals {
				wTrimmed := strings.TrimSpace(w)
				aTrimmmed := strings.TrimSpace(a)
				if wTrimmed == aTrimmmed && wTrimmed != "" {
					fmt.Printf("match %s %s  ", w, a)
					// double the multiplier
					if currentMultiplier == 0 {
						currentMultiplier = 1
					} else {
						currentMultiplier *= 2
					}

				}
			}
		}

		fmt.Printf("\nend of script %d \n  ", currentMultiplier)
		sum += currentMultiplier
	}
	return sum
}
