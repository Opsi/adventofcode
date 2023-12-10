package main

import (
	"fmt"
	"os"

	"github.com/Opsi/adventofcode/four"
	"github.com/Opsi/adventofcode/nine"
	"github.com/Opsi/adventofcode/one"
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

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("please provide the identifier of the puzzle like 1.1 or 5.2")
	}
	puzzle := os.Args[1]
	switch puzzle {
	case "1.1":
		return calcAnswer("input/one.txt", one.One)
	case "1.2":
		return calcAnswer("input/one.txt", one.Two)
	case "2.1":
		return calcAnswer("input/two.txt", two.One)
	case "2.2":
		return calcAnswer("input/two.txt", two.Two)
	case "3.1":
		return calcAnswer("input/three.txt", three.One)
	case "3.2":
		return calcAnswer("input/three.txt", three.Two)
	case "4.1":
		return calcAnswer("input/four.txt", four.One)
	case "4.2":
		return calcAnswer("input/four.txt", four.Two)
	case "9.1":
		return calcAnswer("input/nine.txt", nine.One)
	case "9.2":
		return calcAnswer("input/nine.txt", nine.Two)
	default:
		return fmt.Errorf("unknown puzzle %s", puzzle)
	}
}

func calcAnswer(path string, answerFunc func([]string) (int, error)) error {
	lines, err := util.ReadLines(path)
	if err != nil {
		return fmt.Errorf("read lines: %v", err)
	}
	value, err := answerFunc(lines)
	if err != nil {
		return fmt.Errorf("calculate answer: %v", err)
	}
	fmt.Println(value)
	return nil
}
