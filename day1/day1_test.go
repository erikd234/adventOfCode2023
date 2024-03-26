package day1

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestDay1(t *testing.T) {
	challenge := `
1abc2

pqr3stu8vwx
a1b2c3d4e5f
treb7uchet文
`
	result := Day1(challenge)
	spew.Dump(result)
	if result != 142 {
		t.Error("Result should be 142. Got: " + fmt.Sprint(result))
	}
}
