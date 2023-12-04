package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

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
	num, err := sumPart1("./input-example.txt")
	// num, err := sumPart1("./input-part-1.txt")
	check(err)
	return num
}

type Line struct {
	data string
}

func (l Line) GetNumber() (int, error) {
	var first rune = 0
	var last rune = 0

	for _, r := range l.data {
		if unicode.IsDigit(r) {
			if first == 0 {
				first = r
			}
			last = r
		}
	}

	lineNum := string(first) + string(last)

	fmt.Printf("%s - %s\n", l.data, lineNum)

	num, err := strconv.Atoi(lineNum)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func sumPart1(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		line := &Line{
			data: text,
		}

		num, err := line.GetNumber()
		if err != nil {
			continue
		}

		sum += num
	}

	return sum, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
