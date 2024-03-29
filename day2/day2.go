package day2

import (
	"strconv"
	"strings"
)

type GameParamaters struct {
	RedCount   int
	GreenCount int
	BlueCount  int
}

// Must pasrse the game string to find the amounts.
// Each must not surpase the game paramaters limit
func Day2(game string, params GameParamaters) bool {

	split := strings.Split(game, " ")
	for i, word := range split {
		if strings.Contains(word, "red") {
			amount, err := strconv.Atoi(split[i-1])
			if err != nil {
				panic(err)
			}
			if amount > params.RedCount {
				return false
			}
		}
		if strings.Contains(word, "green") {
			amount, err := strconv.Atoi(split[i-1])
			if err != nil {
				panic(err)
			}
			if amount > params.GreenCount {
				return false
			}
		}
		if strings.Contains(word, "blue") {
			amount, err := strconv.Atoi(split[i-1])
			if err != nil {
				panic(err)
			}
			if amount > params.BlueCount {
				return false
			}
		}
	}
	return true
}
