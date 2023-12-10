package three

import (
	"fmt"
	"strconv"
	"strings"
)

type Field struct {
	// Cells contains the values of the cells in the field.
	// If the value is 0, the cell is empty.
	// If the value is negative, the cell contains a symbol.
	// If the value is positive, the cell contains an ID.
	// Each ID is unique and represents a positive integer in the field.
	Cells    [][]int
	IDValues map[int]int
}

func (f Field) HasSymbolNeighbour(row, col int) bool {
	for rd := -1; rd <= 1; rd++ {
		for cd := -1; cd <= 1; cd++ {
			if rd == 0 && cd == 0 {
				continue
			}
			row := row + rd
			col := col + cd
			if row < 0 || row >= len(f.Cells) {
				continue
			}
			if col < 0 || col >= len(f.Cells[row]) {
				continue
			}
			if f.Cells[row][col] < 0 {
				return true
			}
		}
	}
	return false
}

func (f Field) GearRatio(row, col int) int {
	if f.Cells[row][col] != -2 {
		return -1
	}
	// find the neighbouring ids
	ids := make(map[int]struct{})
	for rd := -1; rd <= 1; rd++ {
		for cd := -1; cd <= 1; cd++ {
			if rd == 0 && cd == 0 {
				continue
			}
			row := row + rd
			col := col + cd
			if row < 0 || row >= len(f.Cells) {
				continue
			}
			if col < 0 || col >= len(f.Cells[row]) {
				continue
			}
			value := f.Cells[row][col]
			if value <= 0 {
				continue
			}
			ids[value] = struct{}{}
			if len(ids) > 2 {
				// too many parts
				return -1
			}
		}
	}
	if len(ids) != 2 {
		// not enough parts
		return -1
	}
	product := 1
	for id := range ids {
		product *= f.IDValues[id]
	}
	return product
}

func One(lines []string) (int, error) {
	field, err := parseField(lines)
	if err != nil {
		return 0, fmt.Errorf("parse field: %v", err)
	}
	partNumbers := make(map[int]struct{})
	for row := 0; row < len(field.Cells); row++ {
		for col := 0; col < len(field.Cells[row]); col++ {
			cell := field.Cells[row][col]
			if cell <= 0 {
				continue
			}
			if field.HasSymbolNeighbour(row, col) {
				partNumbers[cell] = struct{}{}
			}
		}
	}
	sum := 0
	for partNumber := range partNumbers {
		sum += field.IDValues[partNumber]
	}
	return sum, nil
}

func Two(lines []string) (int, error) {
	field, err := parseField(lines)
	if err != nil {
		return 0, fmt.Errorf("parse field: %v", err)
	}
	sum := 0
	for row := 0; row < len(field.Cells); row++ {
		for col := 0; col < len(field.Cells[row]); col++ {
			gearRatio := field.GearRatio(row, col)
			if gearRatio > 0 {
				sum += gearRatio
			}
		}
	}
	return sum, nil
}

func parseField(lines []string) (Field, error) {
	field := Field{
		Cells:    make([][]int, len(lines)),
		IDValues: make(map[int]int),
	}
	currID := 0
	for row, line := range lines {
		field.Cells[row] = make([]int, len(line))
		currNrString := ""
		for col, c := range line {
			switch {
			case strings.ContainsRune("0123456789", c):
				if currNrString == "" {
					currID++
				}
				field.Cells[row][col] = currID
				currNrString += string(c)
			default:
				var value int
				switch c {
				case '.': // empty
					value = 0
				case '*': // gear
					value = -2
				default: // symbol
					value = -1
				}
				field.Cells[row][col] = value
				if currNrString != "" {
					nr, err := strconv.Atoi(currNrString)
					if err != nil {
						return Field{}, fmt.Errorf("parse number string '%s': %v", currNrString, err)
					}
					field.IDValues[currID] = nr
					currNrString = ""
				}
			}
		}
		if currNrString != "" {
			nr, err := strconv.Atoi(currNrString)
			if err != nil {
				return Field{}, fmt.Errorf("parse number string '%s': %v", currNrString, err)
			}
			field.IDValues[currID] = nr
			currNrString = ""
		}
	}
	return field, nil
}
