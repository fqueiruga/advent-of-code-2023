package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
		num, err := sumPart1("./input-part-2.txt")
		check(err)
		return num
	}

	// solve part 1 here
	num, err := sumPart1("./input-example.txt")
	// num, err := sumPart1("./input-part-1.txt")
	check(err)
	return num
}

var numsAsLetters = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

type Line struct {
	data string
}

func (l Line) GetNumber() (int, error) {
	var first rune = 0
	var last rune = 0

	storeDigit := func(r rune) {
		if first == 0 {
			first = r
		}
		last = r
	}

	for i, r := range l.data {
		if unicode.IsDigit(r) {
			storeDigit(r)
			continue
		}

		for k, literal := range numsAsLetters {
			if strings.HasPrefix(l.data[i:], literal) {
				storeDigit([]rune(strconv.Itoa(k))[0])
				continue
			}
		}

	}

	lineNum := string(first) + string(last)
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
