package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	sum := 0

	filename := "./input-part-1.txt"

	file, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		game := NewGame(text)
		if game == nil {
			return 0
		}

		if game.IsValid(bag) {
			sum += game.Id
		}
	}

	return sum
}

// Map of color -> number of balls
type Set map[string]int

var bag Set = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type Game struct {
	Id   int
	sets []Set
}

func NewGame(line string) *Game {
	split := strings.Split(line, ":")
	id, err := strconv.Atoi(strings.Split(split[0], " ")[1])
	if err != nil {
		return nil
	}

	setsStr := strings.TrimSpace(split[1])
	setsStr = strings.ReplaceAll(setsStr, "; ", ";")
	setsStr = strings.ReplaceAll(setsStr, ", ", ",")

	sets := []Set{}
	for _, setStr := range strings.Split(setsStr, ";") {
		set := make(map[string]int)

		for _, setColorStr := range strings.Split(setStr, ",") {
			value, err := strconv.Atoi(strings.Split(setColorStr, " ")[0])
			if err != nil {
				fmt.Printf("Error parsing color set string %s\n", setColorStr)
				return nil
			}
			color := strings.Split(setColorStr, " ")[1]

			set[color] = value
		}

		sets = append(sets, set)
	}

	return &Game{
		Id:   id,
		sets: sets,
	}
}

func (g Game) IsValid(bag Set) bool {
	for _, set := range g.sets {
		for color, amount := range set {
			if amount > bag[color] {
				return false
			}
		}
	}
	return true
}
