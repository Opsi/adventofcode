package main

import (
	"fmt"
	"os"

	"github.com/Opsi/adventofcode/nine"
	"github.com/Opsi/adventofcode/one"
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
		value, err := one.One(document)
		if err != nil {
			return fmt.Errorf("one one: %v", err)
		}
		fmt.Println(value)
		return nil
	case "1.2":
		document := readDocument("input/one.txt")
		value, err := one.Two(document)
		if err != nil {
			return fmt.Errorf("one two: %v", err)
		}
		fmt.Println(value)
		return nil
	case "9.1":
		document := readDocument("input/nine.txt")
		value, err := nine.One(document)
		if err != nil {
			return fmt.Errorf("nine one: %v", err)
		}
		fmt.Println(value)
		return nil
	case "9.2":
		document := readDocument("input/nine.txt")
		value, err := nine.Two(document)
		if err != nil {
			return fmt.Errorf("nine two: %v", err)
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
