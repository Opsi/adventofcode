package eight

import (
	"fmt"
	"regexp"
	"strings"
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
	LeftNodeID  string
	RightNodeId string
	IsGoal      bool
}

func parseInputOne(lines []string) (Input, error) {
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
			LeftNodeID:  matches[2],
			RightNodeId: matches[3],
			IsGoal:      matches[1] == Goal,
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
	input, err := parseInputOne(lines)
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
			currNodeID = node.LeftNodeID
		} else {
			currNodeID = node.RightNodeId
		}
	}
}

func parseInputTwo(lines []string) (Input, error) {
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
			LeftNodeID:  matches[2],
			RightNodeId: matches[3],
			IsGoal:      strings.HasSuffix(matches[1], "Z"),
		}
	}
	return input, nil
}

func calcCyclesToGoal(input Input, rootNodeID string) (int, error) {
	cycleCount := 0
	currNodeID := rootNodeID
	for {
		// we go through all directions in a cycle
		for _, direction := range input.Directions {
			node, ok := input.Nodes[currNodeID]
			if !ok {
				return 0, fmt.Errorf("no node with id %s", currNodeID)
			}
			if direction == Left {
				currNodeID = node.LeftNodeID
			} else {
				currNodeID = node.RightNodeId
			}
		}
		cycleCount++

		// check if we reached the goal node
		node, ok := input.Nodes[currNodeID]
		if !ok {
			return 0, fmt.Errorf("no node with id %s", currNodeID)
		}
		if node.IsGoal {
			return cycleCount, nil
		}
	}
}

// Calculate the greatest common divisor (GCD) using the Euclidean algorithm.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Calculate the least common multiple (LCM) of two integers.
func lcm(a, b int) int {
	gcdAB := gcd(a, b)
	return (a * b) / gcdAB
}

func leastCommonMultiple(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	currLCM := numbers[0]
	for _, n := range numbers[1:] {
		currLCM = lcm(currLCM, n)
	}
	return int(currLCM)
}

func Two(lines []string) (int, error) {
	// after some data exploration i found out that each
	// root node runs into exactly one goal node after
	// a certain number of cycles. This makes this whole
	// problem a lot easier.
	input, err := parseInputTwo(lines)
	if err != nil {
		return 0, fmt.Errorf("parse input: %w", err)
	}

	rootNodeIDs := make([]string, 0)
	for nodeID := range input.Nodes {
		if strings.HasSuffix(nodeID, "A") {
			rootNodeIDs = append(rootNodeIDs, nodeID)
		}
	}

	// for each root node calculate the number of cycles to go
	// through to get to its goal node
	cyclesToGoal := make([]int, len(rootNodeIDs))
	for i, rootNodeID := range rootNodeIDs {
		var err error
		cyclesToGoal[i], err = calcCyclesToGoal(input, rootNodeID)
		if err != nil {
			return 0, fmt.Errorf("calc cycles to goal of %s: %w", rootNodeID, err)
		}
	}

	// return the least common multiple of all cycles
	lcm := leastCommonMultiple(cyclesToGoal)
	return lcm * len(input.Directions), nil
}
