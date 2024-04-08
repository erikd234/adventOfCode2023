package day4

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// For this puzzle we need to parse text into a some form of key value pairs
// Then we need to interpret those key value pairs
// These key value pairs are then used to map seeds into
// Its actually a nestes system of maps where
//  1. Seed -> Soil
//  2. soild -> fertilizer
//  3. fertilizer -> water
//  4. water -> light
//  5. light - > temperature
//
// Then finally the tempurature value gets you the location value.
// We then return the LOWEST locatin value of the lot
//
// Strategy:
// I think theres a few ways to solve it. The first that comes to mind is the following
//
// FIRST RUNTHROUGH
//  1. Parse the seeds to keys
//
// 2nd Runtthrough -> end
//
//  1. Parse map
//
//  2. Use map to convert to next keys
//
//     The structure of the input is the following
//     seeds: 12 3 23 23 23 23... number seperated by whitespace
//     arbitrarty new lines
//     maplabel map:
//     deststart sourcestart range
//     N numer of them
//
//     Then we have N number
//     nmaplabel map:
//     same format as above
//
//     At the end of the file we have
//     humidity-to-location map:
//     same format as above
//
//     Thus we will end when we fine the location string in the map label:
//     OR if we reached the end of file early we will panic since location was not seen
//
//To parse the map itself we need to do the following
//For each range line
// 1. Create a new empty map
// 2. Create keys as source++ to range with destition++ to range
// 3. Now interate from 0 to highest seed number making key = i = value IF it not existing already

type MapLineParsed struct {
	Dest  int
	Src   int
	Range int
}

func Day4(puzzle io.Reader) int {
	scanner := bufio.NewScanner(puzzle)
	var keys []int

	// Parse seed
	for scanner.Scan() {
		fmt.Println("Started the seed scan")
		line := scanner.Text()
		_, seedsSeperatedBySpace, found := strings.Cut(line, ":")
		if !found {
			continue
		}
		seeds := strings.Fields(seedsSeperatedBySpace)
		for _, seed := range seeds {
			seedInt, err := strconv.Atoi(seed)
			if err != nil {
				panic(err)
			}
			keys = append(keys, seedInt)
		}
		break
	}
	fmt.Println("End seedScan")
	spew.Dump(keys)

	// parse maps until its finished
	// Every 2 maps until EOF recalcuate keys
	var mapLines []MapLineParsed
	firstMap := true

	fmt.Println("Start Scanning Maps")
nextLine:
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case firstMap == true && strings.Contains(line, "map"):
			fmt.Println("First Map found")
			firstMap = false
		case strings.Contains(line, "map"):
			// modifying the key array
			spew.Dump(mapLines)
			fmt.Println("Keys Before")
			spew.Dump(keys)
			for i := 0; i < len(keys); i++ {
				keys[i] = convertKey(keys[i], mapLines)
			}
			fmt.Println("Keys After")
			spew.Dump(keys)
			mapLines = []MapLineParsed{}
			fmt.Println("SHOULD BE EMPTY")
			spew.Dump(mapLines)
			fmt.Println("Map Section Ended")
		default:
			mapLine, empty := parseMapLine(line)
			if empty {
				continue nextLine
			}
			fmt.Println("Reading line")
			fmt.Println(line)
			mapLines = append(mapLines, mapLine)
		}

	}
	// modifying the key array
	// Final exit will not have the ending
	fmt.Println("Final Keys")
	spew.Dump(keys)
	for i := 0; i < len(keys); i++ {
		keys[i] = convertKey(keys[i], mapLines)
	}
	fmt.Println("Final After")
	spew.Dump(keys)
	return slices.Min(keys)
}

func parseMapLine(line string) (mlp MapLineParsed, empty bool) {
	entriesString := strings.Fields(line)
	if len(entriesString) < 3 {
		return MapLineParsed{}, true
	}
	var entriesInt []int
	for _, s := range entriesString {
		sInt, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		entriesInt = append(entriesInt, sInt)
	}
	if len(entriesInt) > 3 {
		panic("Map line too long")
	}
	return MapLineParsed{
		Dest:  entriesInt[0],
		Src:   entriesInt[1],
		Range: entriesInt[2],
	}, false
}

// Converts the key using parsed maplines
func convertKey(key int, mapLines []MapLineParsed) int {
	newKey := key
	for _, ml := range mapLines {
		start := ml.Src
		end := ml.Src + ml.Range
		if start <= newKey && newKey <= end {
			// for example 23 68 5 - 25 would map to 70 so 25 + (68-23)
			newKey = key + (ml.Dest - ml.Src)
			return newKey
		}
		if newKey < 0 {
			panic("NEWKEY IS LESS THAN ZERO OH ON")
		}
	}
	return newKey
}
