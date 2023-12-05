package main

import (
	"fmt"
	"os"

	"github.com/Opsi/adventofcode/adventofcode"
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
		document := readDocument("input/one.txt")
		value, err := adventofcode.OneOne(document)
		if err != nil {
			return fmt.Errorf("one one: %v", err)
		}
		fmt.Println(value)
		return nil
	default:
		return fmt.Errorf("unknown puzzle %s", puzzle)
	}
}

func readDocument(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("read file: %v", err))
	}
	return string(data)
}
