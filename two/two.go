package two

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	ID   int
	Sets []CubeSet
}

type CubeSet struct {
	Reds   int
	Greens int
	Blues  int
}

var oneSet = CubeSet{
	Reds:   12,
	Greens: 13,
	Blues:  14,
}

func (g Game) MinCubeSet() CubeSet {
	var result CubeSet
	for _, set := range g.Sets {
		result.Reds = max(result.Reds, set.Reds)
		result.Greens = max(result.Greens, set.Greens)
		result.Blues = max(result.Blues, set.Blues)
	}
	return result
}

func (c1 CubeSet) FitsInto(c2 CubeSet) bool {
	return c1.Reds <= c2.Reds && c1.Greens <= c2.Greens && c1.Blues <= c2.Blues
}

func (c CubeSet) Power() int {
	return c.Reds * c.Greens * c.Blues
}

func parseGame(s string) (Game, error) {
	var game Game
	splitted := strings.Split(s, ":")
	if len(splitted) != 2 {
		return Game{}, fmt.Errorf("invalid game: %s", s)
	}
	id, err := parseID(splitted[0])
	if err != nil {
		return Game{}, fmt.Errorf("parse id: %v", err)
	}
	game.ID = id
	sets := strings.Split(splitted[1], ";")
	for _, set := range sets {
		cs, err := parseSet(set)
		if err != nil {
			return Game{}, fmt.Errorf("parse set: %v", err)
		}
		game.Sets = append(game.Sets, cs)
	}
	return game, nil
}

func parseID(s string) (int, error) {
	splitted := strings.Split(s, " ")
	if len(splitted) != 2 {
		return 0, fmt.Errorf("invalid id: %s", s)
	}
	id, err := strconv.Atoi(splitted[1])
	if err != nil {
		return 0, fmt.Errorf("invalid id: %s", splitted[1])
	}
	return id, nil
}

func parseSet(s string) (CubeSet, error) {
	var cs CubeSet
	entries := strings.Split(s, ",")
	for _, entry := range entries {
		trimmed := strings.TrimSpace(entry)
		splitted := strings.Split(trimmed, " ")
		if len(splitted) != 2 {
			return CubeSet{}, fmt.Errorf("invalid entry: '%s'", entry)
		}
		amount, err := strconv.Atoi(splitted[0])
		if err != nil {
			return CubeSet{}, fmt.Errorf("invalid amount: '%s'", splitted[0])
		}
		color := splitted[1]
		switch color {
		case "red":
			cs.Reds = amount
		case "green":
			cs.Greens = amount
		case "blue":
			cs.Blues = amount
		default:
			return CubeSet{}, fmt.Errorf("invalid color: '%s'", color)
		}
	}
	return cs, nil
}

func One(lines []string) (int, error) {
	sum := 0
	for i, line := range lines {
		game, err := parseGame(line)
		if err != nil {
			return 0, fmt.Errorf("parse line %d: %v", i, err)
		}
		if game.MinCubeSet().FitsInto(oneSet) {
			sum += game.ID
		}
	}
	return sum, nil
}

func Two(lines []string) (int, error) {
	sum := 0
	for i, line := range lines {
		game, err := parseGame(line)
		if err != nil {
			return 0, fmt.Errorf("parse line %d: %v", i, err)
		}
		sum += game.MinCubeSet().Power()
	}
	return sum, nil
}
