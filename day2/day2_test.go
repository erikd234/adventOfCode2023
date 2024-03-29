// https://adventofcode.com/2023/day/2
package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestDay2(t *testing.T) {
	var sumIds int
	params := GameParamaters{RedCount: 12, GreenCount: 13, BlueCount: 14}
	fmt.Println("testing")
	spew.Dump("testing")
	games := []struct {
		Game               string
		ExpectedGameResult bool
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
	}

	re := regexp.MustCompile("(Game) ([0-9]*)")
	for _, g := range games {
		r := Day2(g.Game, params)
		if g.ExpectedGameResult != r {
			errMsg := fmt.Sprintf("%v \n Day2 game identified game as %v and not %v",
				g.Game,
				r,
				g.ExpectedGameResult)

			t.Errorf(errMsg)
		} else if r == true {
			gamePrefix := re.FindSubmatch([]byte(g.Game))
			fmt.Printf("Group1: %s Group2: %s \n", string(gamePrefix[1]), string(gamePrefix[2]))
			id, err := strconv.Atoi(string(gamePrefix[2]))
			if err != nil {
				panic(err)
			}
			sumIds += id
		}
	}
	fmt.Printf("The result was: %d \n", sumIds)
	if sumIds != 8 {
		errMsg := fmt.Sprintf("The result was: %d, wanted 8 \n", sumIds)
		t.Error(errMsg)
	}
}
