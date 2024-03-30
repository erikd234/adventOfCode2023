// https://adventofcode.com/2023/day/4
package day4

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestDay4(t *testing.T) {
	input1 :=
		`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21  53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 23 56 72 | 74 77 10 23 35 67 36 11`

	input2, err := os.Open("./challenge.txt")
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		Puzzle    io.Reader
		SumPoints int
	}{
		{
			strings.NewReader(input1),
			14,
		},
		{
			input2,
			25004,
		},
	}

	for _, p := range tests {
		r := Day4(p.Puzzle)
		if r != p.SumPoints {
			t.Error(fmt.Sprintf("Got %d for sum wanted %d", r, p.SumPoints))
		}
	}

}
