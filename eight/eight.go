package eight

import (
	"fmt"
	"regexp"
)

type Direction bool

const (
	Left  Direction = false
	Right Direction = true
)

func (d Direction) String() string {
	if d == Left {
		return "Left"
	}
	return "Right"
}

type Input struct {
	Directions []Direction
	Nodes      map[string]Node
}

const (
	Root = "AAA"
	Goal = "ZZZ"
)

type Node struct {
	Left  string
	Right string
}

func parseInput(lines []string) (Input, error) {
	if len(lines) < 3 {
		return Input{}, fmt.Errorf("invalid input with only %d lines", len(lines))
	}
	directions := make([]Direction, len(lines[0]))
	for i, c := range lines[0] {
		switch c {
		case 'L':
			directions[i] = Left
		case 'R':
			directions[i] = Right
		default:
			return Input{}, fmt.Errorf("invalid direction '%c'", c)
		}
	}
	nodeLineRegex := regexp.MustCompile(`^(\w+) = \((\w+), (\w+)\)$`)
	input := Input{
		Directions: directions,
		Nodes:      make(map[string]Node),
	}
	for _, line := range lines[1:] {
		matches := nodeLineRegex.FindStringSubmatch(line)
		if len(matches) != 4 {
			return Input{}, fmt.Errorf("invalid node line '%s'", line)
		}
		input.Nodes[matches[1]] = Node{
			Left:  matches[2],
			Right: matches[3],
		}
	}
	if _, ok := input.Nodes[Root]; !ok {
		return Input{}, fmt.Errorf("no root node found")
	}
	if _, ok := input.Nodes[Goal]; !ok {
		return Input{}, fmt.Errorf("no goal node found")
	}
	return input, nil
}

func One(lines []string) (int, error) {
	input, err := parseInput(lines)
	if err != nil {
		return 0, fmt.Errorf("parse input: %v", err)
	}
	stepCount := 0
	currNodeID := Root
	for {
		if currNodeID == Goal {
			return stepCount, nil
		}
		dirctionIndex := stepCount % len(input.Directions)
		direction := input.Directions[dirctionIndex]
		stepCount++
		node, ok := input.Nodes[currNodeID]
		if !ok {
			return 0, fmt.Errorf("no node with id %s", currNodeID)
		}
		if direction == Left {
			currNodeID = node.Left
		} else {
			currNodeID = node.Right
		}
	}
}

func Two(lines []string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
