package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/Opsi/adventofcode/five"
	"github.com/Opsi/adventofcode/four"
	"github.com/Opsi/adventofcode/nine"
	"github.com/Opsi/adventofcode/one"
	"github.com/Opsi/adventofcode/six"
	"github.com/Opsi/adventofcode/three"
	"github.com/Opsi/adventofcode/two"
	"github.com/Opsi/adventofcode/util"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type AnswerFunc = func([]string) (int, error)

func run() error {
	solveFunctions := map[string]AnswerFunc{
		"1.1": one.One,
		"1.2": one.Two,
		"2.1": two.One,
		"2.2": two.Two,
		"3.1": three.One,
		"3.2": three.Two,
		"4.1": four.One,
		"4.2": four.Two,
		"5.1": five.One,
		"5.2": five.Two,
		"6.1": six.One,
		"6.2": six.Two,
		"9.1": nine.One,
		"9.2": nine.Two,
	}

	if len(os.Args) != 2 {
		return fmt.Errorf("please provide the identifier of the puzzle like 1.1 or 5.2")
	}
	puzzle := os.Args[1]
	identifierRegex := regexp.MustCompile(`^(\d+)\.\d+$`)
	matches := identifierRegex.FindStringSubmatch(puzzle)
	if len(matches) != 2 {
		return fmt.Errorf("invalid puzzle identifier %s", puzzle)
	}
	inputPath := fmt.Sprintf("input/%s.txt", matches[1])
	lines, err := util.ReadLines(inputPath)
	if err != nil {
		return fmt.Errorf("read lines for file %s: %v", inputPath, err)
	}

	answerFunc, ok := solveFunctions[puzzle]
	if !ok {
		return fmt.Errorf("unknown puzzle %s", puzzle)
	}

	value, err := answerFunc(lines)
	if err != nil {
		return fmt.Errorf("calculate answer: %v", err)
	}
	fmt.Println(value)
	return nil
}
