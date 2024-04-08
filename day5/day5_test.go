// https://adventofcode.com/2023/day/4
package day4

import (
	"io"
	"os"
	"testing"
)

func TestDay4(t *testing.T) {
	input1, err := os.Open("./simple.txt")
	if err != nil {
		t.Error(err)
	}
	// A test with extra spaces thrownn in
	input2, err := os.Open("./simple2.txt")
	if err != nil {
		t.Error(err)
	}

	input3, err := os.Open("./challenge.txt")
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		Puzzle io.Reader
		Lowest int
	}{
		{
			input1,
			35,
		},
		{
			input2,
			35,
		},
		{
			input3,
			26273516, // correct awnser

		},
	}

	for _, p := range tests {
		r := Day4(p.Puzzle)
		if r != p.Lowest {
			t.Errorf("Got %d, wanted %d", r, p.Lowest)
		}
	}

}
