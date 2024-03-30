// https://adventofcode.com/2023/day/3
package day3

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

// we need to make the string a 2d slice.
// We need to define boundries at 0 and the width
// then we need to define boundries at 0 and the height
// Algo is like this.
//  1. Find number, record start, record end.
//  2. Search for symbols, for starty -1 (dont go before edge) to starty+1 (dont go over edge) :
//     start search at startx - 1 or startedge,
//     end search at startx + 1 or endedge
//
// We will have it as an array of runes since its comone from a string
// ASSUMES INPUT HAS NO EMPTY LINES or LINES WITH ONLY \n
func Day3(puzzle string) int {
	trimedPuzzle := strings.Trim(puzzle, "\n ")
	spew.Dump(trimedPuzzle)
	lines := strings.Split(trimedPuzzle, "\n")
	var box [][]rune
	for _, l := range lines {
		box = append(box, slices.Clone([]rune(l)))
	}

	// do search
	startx, endx := -1, -1
	searchingEndOfDigits := false
	sum := 0

	// 1 2 3 4 5 6  in this case we want to trigger branch three after
	//				the row is over. if the row is over that means x == 0
	//				we can alwasy assume the first one will NEVER end on x ==0
	//				. Also in the case we have length 1 number, endx wont be found
	//              SOlution is assign endx in the first part too.
	// 1 2 . .
	for y, row := range box {
		for x, rune := range row {
			if x != 0 && unicode.IsDigit(rune) && searchingEndOfDigits {
				// extends the digit end
				endx = x
			} else if unicode.IsDigit(rune) && !searchingEndOfDigits {
				// looks for the very first digit and saves location
				startx = x
				endx = x
				searchingEndOfDigits = true
			} else if searchingEndOfDigits {
				// it means we
				if symbolAroundDigit(box, startx, endx, y) {
					sum += getDigitFromRunes(row, startx, endx)
				}
				searchingEndOfDigits = false
				endx = -1
				startx = -1
			}
		}
		// handle case that we ended on an edge without searching
		if searchingEndOfDigits {
			if symbolAroundDigit(box, startx, endx, y) {
				sum += getDigitFromRunes(row, startx, endx)
			}
			searchingEndOfDigits = false
			endx = -1
			startx = -1
		}
	}
	return sum
}

// Performs the search of elements around digits
// box will always have atleast one row.
func symbolAroundDigit(box [][]rune, startx, endx, y int) bool {
	starty := y - 1
	endy := y + 1
	// MODIFYING PARAMS HERE
	startx = startx - 1
	endx = endx + 1
	if starty < 0 {
		starty = 0
	}
	if endy >= len(box) {
		endy = len(box) - 1
	}
	if startx < 0 {
		startx = 0
	}
	if endx >= len(box[0]) {
		endx = len(box) - 1
	}

	for y := starty; y <= endy; y++ {
		for x := startx; x <= endx; x++ {
			r := box[y][x]
			fmt.Printf("%c", r)
			if r != '.' && (unicode.IsSymbol(r) || unicode.IsPunct(r)) {
				return true
			}
		}
		fmt.Println()
	}
	fmt.Println()
	return false
}

func getDigitFromRunes(row []rune, startx, endx int) int {
	s := string(row[startx : endx+1])
	d, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return d
}
