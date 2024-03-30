package day3

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestDay1(t *testing.T) {
	challenge := `
467..114..
...*......
..35..633.
.....1#...
617*......
.....+.58.
..592.....
......755.
...$.*...1
.664.598./
`
	result := Day3(challenge)
	spew.Dump(result)
	if result != 4363 {
		t.Error("Result should be 4361. Got: " + fmt.Sprint(result))
	}
	file, err := os.Open("./dump.txt")
	if err != nil {
		t.Error(err)
	}

	scanner := bufio.NewScanner(file)
	bodyString := ""
	for scanner.Scan() {
		bodyString += fmt.Sprintf("%s\n", scanner.Text())
	}
	fmt.Println(bodyString)
	challenge2 := bodyString
	result = Day3(challenge2)
	spew.Dump(result)
	if result != 514969 {
		t.Error("Result should be 4361. Got: " + fmt.Sprint(result))
	}
	// GOT IT WOO
}
